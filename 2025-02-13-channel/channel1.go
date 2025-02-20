package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//r := rand.New(rand.NewSource(time.Now().Unix()))
	//
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//
	//go func() {
	//	time.Sleep(time.Second * time.Duration(r.Intn(5)))
	//	ch1 <- 1
	//}()
	//
	//go func() {
	//	time.Sleep(time.Second * time.Duration(r.Intn(5)))
	//	ch2 <- 2
	//}()
	//
	//fmt.Println(<-ch1)
	//fmt.Println(<-ch2)

	//c := make(chan int)
	//quit := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	//fibonacci(c, quit)

	//myChan := make(chan int)
	//
	//go sender(myChan, "S1")
	//go sender(myChan, "S2")
	//go sender(myChan, "S3")
	//
	//start := 0
	//
	//for {
	//	start += <-myChan
	//	fmt.Println(start)
	//
	//	if start >= 300 {
	//		break
	//	}
	//}

	//myChan := publisher()
	//maxConsumer := 5
	//
	//for i := 1; i <= maxConsumer; i++ {
	//	go consumer(myChan, fmt.Sprintf("%d", i))
	//}
	//
	//time.Sleep(time.Second * 10)

	s := sumAllStreams(
		streamNumbers(1, 2, 3, 4, 5),
		streamNumbers(8, 8, 3, 3, 10, 12, 14),
		streamNumbers(1, 1, 2, 2, 4, 4, 6),
	)

	fmt.Println(<-s)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func sender(c chan<- int, name string) {
	for i := 1; i <= 100; i++ {
		c <- 1
		fmt.Printf("%s has sent 1 to channel\n", name)
		runtime.Gosched()
	}
}

func publisher() <-chan int {
	c := make(chan int)

	go func() {
		for i := 1; i <= 1000; i++ {
			c <- i
		}

		close(c)
	}()

	return c
}

func consumer(c <-chan int, name string) {
	counter := 0

	for value := range c {
		fmt.Printf("Consumer %s is doing task %d\n", name, value)
		counter++
		time.Sleep(time.Millisecond * 20)
	}

	fmt.Printf("Consumer %s has finished %d task(s)\n", name, counter)
}

func streamNumbers(numbers ...int) <-chan int {
	c := make(chan int)

	go func() {
		for n := range numbers {
			c <- n
		}

		close(c)
	}()

	return c
}

func sumAllStreams(streams ...<-chan int) <-chan int {
	sumChan := make(chan int)
	counter := 0
	wc := new(sync.WaitGroup)

	wc.Add(len(streams))
	fmt.Println(len(streams))

	for i := 0; i < len(streams); i++ {
		go func(s <-chan int) {
			for n := range s {
				fmt.Println("n", n, "counter", counter)
				counter += n
			}
			wc.Done()
		}(streams[i])
	}

	go func() {
		wc.Wait()
		sumChan <- counter
	}()

	return sumChan
}
