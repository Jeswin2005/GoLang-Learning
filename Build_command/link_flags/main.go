package main

import "fmt"

// go build -ldflags="-s -w" -o myapp main.go
// -s -w → remove symbol table & DWARF(Debug format used by debuggers)

func main(){
	fmt.Println("Hello")
}