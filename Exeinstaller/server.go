package main

import (
	"fmt"
	"net/http"
)

func main() {
	currentBin := "Version1.bin"

	http.HandleFunc("/current-bin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, currentBin)
	})

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, currentBin)
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
