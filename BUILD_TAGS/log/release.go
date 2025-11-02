//go:build !debug && !test

package log

// go build -tags=dep debug.go

func DebugLog(msg string) {
}