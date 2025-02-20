package main

import (
	"fmt"
	"time"
)

func main() {
	//ch := make(chan int, 100) // Giới hạn 100 phần tử trong channel
	//
	//go sendData(ch)
	//processData(ch)

	//bufferedChan := make(chan int, 5)
	//bufferedChan <- 1
	//bufferedChan <- 2
	//bufferedChan <- 3
	//bufferedChan <- 4
	//bufferedChan <- 5
	//
	//fmt.Printf("BufferChan has len = %d, cap = %d\n", len(bufferedChan), cap(bufferedChan))
	//bufferedChan <- 6 // block here
	//fmt.Println(111)

	//bufferedChan := make(chan int, 5)
	//fmt.Printf("BufferChan has len = %d, cap = %d\n", len(bufferedChan), cap(bufferedChan))
	//<-bufferedChan // block here
	//fmt.Println(111)

	bufferedChan := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		bufferedChan <- i
	}

	for v := range bufferedChan {
		fmt.Println(v)
	}

	//bufferedChan := make(chan int, 1)
	//unbufferedChan := make(chan int)

	//bufferedChan <- 1 // OK
	//unbufferedChan <- 1 // deadlock
}

// Tại sao dùng channel truyền data trong khi có thể implement truyền thẳng trên go func?
// Tránh race condition
// Giao tiếp an toàn mà không cần lock
func f1() {
	ch := make(chan int, 1000)
	count := 0

	for i := 0; i < 1000; i++ {
		go func() { ch <- 1 }()
	}

	for i := 0; i < 1000; i++ {
		count += <-ch
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Count:", count)
}

// Flow Control: điều tiết tốc độ xử lý để tránh quá tải hệ thống
func sendData(ch chan int) {
	for i := 0; i < 1000; i++ {
		ch <- i // Đẩy dữ liệu vào channel, tránh quá tải
	}
	close(ch)
}

func processData(ch chan int) {
	for n := range ch {
		fmt.Println("Processing:", n)
		time.Sleep(10 * time.Millisecond) // Giả lập xử lý
	}
}
