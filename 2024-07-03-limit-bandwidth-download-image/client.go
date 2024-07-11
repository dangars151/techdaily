package main

import (
	"github.com/go-resty/resty/v2"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 11; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			resty.New().R().Get("http://localhost:8080/download")
		}()
	}

	wg.Wait()
}
