package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

const (
	binDir   = "/home/jeswin-pt8024/mini-project/hostzip/bin"
	logExe   = "logworker"
	updateExe = "updater"
)

func runOnce(ctx context.Context, exe string) error {
	path := filepath.Join(binDir, exe)

	if fi, err := os.Stat(path); err != nil {
		return fmt.Errorf("executable not found: %s: %w", path, err)
	} else if fi.IsDir() {
		return fmt.Errorf("expected file but found dir: %s", path)
	}

	cmd := exec.CommandContext(ctx, path)
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigc
		log.Printf("main: signal received: %v - shutting down", sig)
		cancel()
	}()

	logTicker := time.NewTicker(5 * time.Minute)
	updateTicker := time.NewTicker(2 * time.Minute)
	defer logTicker.Stop()
	defer updateTicker.Stop()

	go func() {
		log.Println("main: initial logworker run")
		if err := runOnce(ctx, logExe); err != nil {
			log.Println("logworker error:", err)
		}
	}()

	go func() {
		log.Println("main: initial updater run")
		if err := runOnce(ctx, updateExe); err != nil {
			log.Println("updater error:", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("main: exiting")
			time.Sleep(500 * time.Millisecond)
			return
		case <-logTicker.C:
			go func() {
				if err := runOnce(ctx, logExe); err != nil {
					log.Println("logworker run error:", err)
				}
			}()
		case <-updateTicker.C:
			go func() {
				if err := runOnce(ctx, updateExe); err != nil {
					log.Println("updater run error:", err)
				}
			}()
		}
	}
}


