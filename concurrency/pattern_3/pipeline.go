package main

import "fmt"

// Stage 1: multiply by 2
func stage1(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * 2
		}
		close(out)
	}()
	return out
}

// stage 1 output is input for stage 2 so only stage 1 returns a read/recieve only channel
// Stage 2: add 1
func stage2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v + 1
		}
		close(out)
	}()
	return out
}

func main() {
	nums := make(chan int, 3)
	nums <- 1
	nums <- 2
	nums <- 3
	close(nums)

	out1 := stage1(nums)
	out2 := stage2(out1)

	for v := range out2 {
		fmt.Println(v)
	}
}
