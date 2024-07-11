package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"time"
)

const maxClients int64 = 10
const maxFrame = 10
const maxSize = 64

func main() {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	key := "key"

	s := NewServer(rdb, ctx, key)

	r := gin.Default()
	r.GET("/download", func(c *gin.Context) {
		err := s.LimitClients()

		if err != nil {
			fmt.Println("failed")
			fmt.Printf("%v", err)
		} else {
			fmt.Println("success")
		}
	})
	r.Run()
}

type server struct {
	rdb *redis.Client
	ctx context.Context
	key string
}

func NewServer(rdb *redis.Client, ctx context.Context, key string) *server {
	s := server{
		rdb: rdb,
		ctx: ctx,
		key: key,
	}

	err := rdb.Set(ctx, key, 0, 0).Err()
	if err != nil {
		panic(err)
	}

	return &s
}

func (s *server) LimitClients() error {
	a := s.rdb.Incr(s.ctx, s.key)

	if a.Val() > maxClients {
		for i := 0; i < 10; i++ {
			err := s.rdb.Decr(s.ctx, s.key).Err()

			if err == nil {
				break
			}
		}

		return errors.New("error")
	}

	return nil
}

func (s *server) DownloadImage() {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
}
