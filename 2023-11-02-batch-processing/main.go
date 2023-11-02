package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

const (
	DefaultSize        = 3
	DefaultExpiredTime = 3 * time.Second
)

// items length = size || timeout => fire batch

type Batch struct {
	Items        []interface{}
	Size         int
	ExpiredTime  time.Duration
	producerChan chan interface{}
	quit         chan int
	fire         func(items []interface{})
}

func NewBatch() *Batch {
	return &Batch{
		Items:        make([]interface{}, 0),
		Size:         DefaultSize,
		ExpiredTime:  DefaultExpiredTime,
		producerChan: make(chan interface{}),
		quit:         make(chan int),
	}
}

func (b *Batch) Start() {
	for {
		select {
		case item := <-b.producerChan:
			b.Items = append(b.Items, item)
			if len(b.Items) == b.Size {
				b.Fire(b.Items)
				b.Items = b.Items[:0]
			}
		case <-time.After(b.ExpiredTime):
			if len(b.Items) > 0 {
				b.Fire(b.Items)
				b.Items = b.Items[:0]
			}
		case <-b.quit:
			if len(b.Items) > 0 {
				b.Fire(b.Items)
				b.Items = b.Items[:0]
			}
			break
		}

	}
}

func (b *Batch) Stop() {
	b.quit <- 0
}

func (b *Batch) Add(item interface{}) {
	b.producerChan <- item
}

func (b *Batch) SetFire(f func(items []interface{})) {
	b.fire = f
}

func (b *Batch) Fire(items []interface{}) {
	b.fire(items)
}

type Comment struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

type Msg struct {
	payload Comment
	result  chan int
}

func main() {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "test",
	})

	b := NewBatch()
	b.SetFire(func(items []interface{}) {
		commentWithChanList := make([]Msg, 0, len(items))
		commentList := make([]Comment, 0, len(items))

		fmt.Println(len(items))

		for _, item := range items {
			msg, _ := item.(Msg)
			commentWithChanList = append(commentWithChanList, msg)
			commentList = append(commentList, msg.payload)
		}

		_, err := db.Model(&commentList).Insert()

		fmt.Println(err)

		for i, c := range commentList {
			result := commentWithChanList[i].result
			result <- c.ID
		}
	})
	go b.Start()

	r := gin.Default()
	r.POST("/comments", func(c *gin.Context) {
		var comment Comment
		c.BindJSON(&comment)

		result := make(chan int)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		b.Add(Msg{payload: comment, result: result})
		select {
		case <-ctx.Done():
			close(result)
			c.JSON(503, "")
		case id := <-result:
			c.JSON(200, id)
		}
	})
	r.Run()
}
