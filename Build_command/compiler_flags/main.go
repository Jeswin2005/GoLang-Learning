package main

import "fmt"

// go build -gcflags="-N" -o opt.exe main.g // no optimzation
// go build -gcflags="-l" -o inline.exe main.go // disable inlining
func main(){
	fmt.Println("Hello")
}