package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

func main() {
	for i := 0; i < 3; i++ {
		go func(i int) {
			client := resty.New().R()

			res, err := client.SetBody(&ClientComment{
				ID:        i + 1,
				Content:   "content",
				CreatedAt: time.Now().Format(time.RFC3339),
			}).Post("http://localhost:8080/comments")

			fmt.Println(string(res.Body()), err)
		}(i)
	}

	time.Sleep(10 * time.Second)
}

type ClientComment struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}
