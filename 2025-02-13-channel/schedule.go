package main

import (
	"fmt"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(2)

	go func() {
		for i := 1; i <= 10000; i++ {
			fmt.Println("111111111111111")
			//runtime.Gosched()
		}
	}()

	go func() {
		for i := 1; i <= 10000; i++ {
			fmt.Println("2222222222222222222222")
			//runtime.Gosched()
		}
	}()

	time.Sleep(time.Second)
}
