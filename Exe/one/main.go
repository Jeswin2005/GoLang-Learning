package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanf("%s",&name)
	fmt.Println("Welcome ",name)
	if len(os.Args) > 1 {
		fmt.Println("Arguments received:", os.Args[0:])
	} else {
		fmt.Println("hi im from another code")
	}
}
