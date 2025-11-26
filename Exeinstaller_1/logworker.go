package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	logFile := "/home/jeswin-pt8024/logs/commonlogs.log"

	if err := os.MkdirAll("/home/jeswin-pt8024/logs", 0755); err != nil {
		fmt.Println("mkdir logs:", err)
		return
	}

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open log:", err)
		return
	}
	defer f.Close()
	_, _ = f.WriteString(fmt.Sprintf("logworker version 2 before sleep ran at %s\n", time.Now().Format(time.RFC3339)))

	time.Sleep(4 * time.Minute)

	_, _ = f.WriteString(fmt.Sprintf("logworker version 2 after sleep ran at %s\n", time.Now().Format(time.RFC3339)))
}
