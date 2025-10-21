package main

import "fmt"
	

func main() {
	fmt.Println("Unbuffered Channel")
	unbufChan := make(chan string) // Unbuffered channel

	// Goroutine sending data
	go func() {
		fmt.Println("Sending 'hello' to unbuffered channel")
		unbufChan <- "hello" // sender blocks until receiver receives
		fmt.Println("Sent 'hello' successfully")
	}()


	// Receiving data
	msg := <-unbufChan
	fmt.Println("Received from unbuffered channel:", msg)

	fmt.Println("\nBuffered Channel")
	bufChan := make(chan string, 2) // Buffered channel with size 2

	// Sending data (does NOT block until buffer is full)
	bufChan <- "first"
	fmt.Println("Sent 'first' to buffered channel")
	bufChan <- "second"
	fmt.Println("Sent 'second' to buffered channel")

	// Receiving data
	fmt.Println("Received from buffered channel:", <-bufChan)
	fmt.Println("Received from buffered channel:", <-bufChan)
}
