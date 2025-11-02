package main

import "fmt"

// GOOS=linux GOARCH=amd64 go build -o myapp-linux (build for linux from any os)
// GOOS=windows GOARCH=amd64 go build -o myapp.exe (build for Windows from any os)
// GOOS=darwin GOARCH=arm64 go build -o myapp-macos-arm (build for macOS from any os)
func main(){
	fmt.Println("Hello")
}