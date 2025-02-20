package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Bộ nhớ trước khi tạo slice:")
	PrintMemUsage()

	_ = createSlice()

	fmt.Println("Bộ nhớ sau khi tạo slice:")
	PrintMemUsage()

	runtime.GC()
	fmt.Println("Bộ nhớ sau khi chạy GC:")
	PrintMemUsage()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Bộ nhớ đang sử dụng: %v MB\n", m.Alloc/1024/1024)
}

func createSlice() []int {
	slice := make([]int, 1_000_000)
	return slice
}
