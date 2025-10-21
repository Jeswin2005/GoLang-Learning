package main

import (
	"fmt"
	"time"
)

func someFunc(num string){
	fmt.Println(num)
}

func main(){
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")

	fmt.Println("in main function")

	time.Sleep(time.Second*2)

}