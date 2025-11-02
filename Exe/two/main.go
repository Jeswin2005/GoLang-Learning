package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main(){
	fmt.Println("Current Code")

	cmd := exec.Command(
		`C:\Users\jeswin-pt8024\Desktop\Go\Exe\one\main.exe`,"hello", "world", "123",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err!= nil{
		fmt.Println("Error running child")
	} else {
		fmt.Println("Child runned successfully")
	}
}
