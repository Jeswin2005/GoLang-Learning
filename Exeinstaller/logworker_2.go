package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	logFile := "/home/jeswin-pt8024/logs/commonlogs.log"

	time.Sleep(30 * time.Second)

	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log: %v\n", err)
		return
	}
	defer f.Close()

	_, _ = f.WriteString(fmt.Sprintf("logworker - version 2 ran at %s\n", time.Now().Format(time.RFC3339)))
}
