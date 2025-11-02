package main

import "fmt"

func demo() {
	defer fmt.Println("function is deferred")
	fmt.Println("function started")
	fmt.Println("function ended")
}

func main() {
	defer fmt.Println("Main is deferred")
	fmt.Println("Main started")
	demo()
	fmt.Println("Main ended")

	fmt.Println()
	fmt.Println("Defer in Loop")
	for i:=0; i<3; i++ {
		defer fmt.Println(i)
	}
}