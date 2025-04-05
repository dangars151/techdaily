package main

import (
	"context"
	"strconv"
	"sync"
	"time"
)

func main() {
	arr := make(chan int, 3)
	arr <- 1
	arr <- 2
	arr <- 3
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "a", 1)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				select {
				case <-ctx.Done():
					{
						println("stop " + strconv.Itoa(i))
						wg.Done()
						return
					}
				case j := <-arr:
					{
						println(j)
					}
				}
			}
		}(i)
	}

	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
}
