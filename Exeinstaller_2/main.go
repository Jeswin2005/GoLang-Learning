package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func executeCommand(exe string) error {
	path := "/home/jeswin-pt8024/first-proj/hostzip/bin/" + exe

	cmd := exec.Command(path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("run failed for %s: %w", exe, err)
	}
	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("main: starting supervisor")

	logworkerTicker := time.NewTicker(5 * time.Minute)
	updaterTicker := time.NewTicker(2 * time.Minute)
	defer logworkerTicker.Stop()
	defer updaterTicker.Stop()

	go func() {
		if err := executeCommand("logworker"); err != nil {
			log.Println("Error running logworker:", err)
		}
	}()

	go func() {
		if err := executeCommand("updater"); err != nil {
			log.Println("Error running updater:", err)
		}
	}()

	for {
		select {
		case <-logworkerTicker.C:
			{
				go func() {
					if err := executeCommand("logworker"); err != nil {
						log.Println("Error running logworker:", err)
					}
				}()
			}

		case <-updaterTicker.C:
			{
				go func() {
					if err := executeCommand("updater"); err != nil {
						log.Println("Error running updater:", err)
					}
				}()
			}

		}

	}
}
