package main

import "fmt"

func main(){
	firstChannel := make(chan string)
	secondChannel := make(chan string)

	go func(){
		firstChannel <- "data from 1st"
	}()

	go func(){
		secondChannel <- "data from 2nd"
	}()

	select {
	case msgFirst := <-firstChannel:
		fmt.Println(msgFirst)
	case msgSecond := <-secondChannel:
		fmt.Println(msgSecond)
	}
}