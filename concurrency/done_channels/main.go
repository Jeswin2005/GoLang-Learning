package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Worker stopped")
			done <- true // signal main that we are done
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	done := make(chan bool)

	go worker(done)

	time.Sleep(3 * time.Second) // let it work
	done <- true                 // tell worker to stop
	<-done                       // wait for confirmation

	fmt.Println("Main finished")
}
