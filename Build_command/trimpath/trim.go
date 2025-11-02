package main

import "fmt"

// go build -trimpath trim.go
// exe not include file path
func main(){
	fmt.Println("Path of the file not exist in the binary")
}