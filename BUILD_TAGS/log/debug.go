//go:build debug || test

package log

import "fmt"

// go build -tags=debug debug.go
func DebugLog(msg string) {
	fmt.Printf("[DEBUG] %s\n", msg)
}
