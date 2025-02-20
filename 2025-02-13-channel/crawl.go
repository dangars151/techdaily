package main

import (
	"fmt"
	"time"
)

const (
	numberOfURLs    = 10000
	numberOfWorkers = 5
)

func crawlURL(queue <-chan int, name string) {
	for v := range queue {
		fmt.Printf("Worker %s is crawling URL %d\n", name, v)
		time.Sleep(time.Second)
	}

	fmt.Printf("Worker %s done and exit\n", name)
}

func startQueue() <-chan int {
	queue := make(chan int, 100)

	go func() {
		for i := 1; i <= numberOfURLs; i++ {
			queue <- i
			fmt.Printf("URL %d has been enqueued\n", i)
		}

		close(queue)
	}()

	return queue
}

func main() {
	queue := startQueue()

	for i := 1; i <= numberOfWorkers; i++ {
		go crawlURL(queue, fmt.Sprintf("%d", i))
	}

	time.Sleep(time.Minute * 5)
}
