package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	logDir := "/home/jeswin-pt8024/logs"
	err := os.MkdirAll(logDir,0755)
	if err != nil {
		fmt.Println("Error in creating log directory",err)
		return
	}	

	filePtr, err := os.OpenFile(logDir + "/logworker.log", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error in creating/writing log file",err)
	}
	defer filePtr.Close()

	filePtr.WriteString(fmt.Sprintf("Logworker before sleep version 2 ran at %s\n",time.Now()))

	time.Sleep(4 * time.Minute)

	filePtr.WriteString(fmt.Sprintf("Logworker after sleep version 2 ran at %s\n",time.Now()))
}