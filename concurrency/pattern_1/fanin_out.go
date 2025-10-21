package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	time.Sleep(time.Millisecond * 500)
	ch <- id * 2
}

func main() {
	ch := make(chan int)
	for i := 1; i <= 3; i++ {
		go worker(i, ch) // Fan-Out launches multiple go routine
	}

	for i := 1; i <= 3; i++ {
		result := <-ch // Fan-In Collect results from multiple goroutines into a single channel
		fmt.Println("Result:", result)
	}
}
