package main

import (
	"fmt"
	"net/http"
)

func main() {
	currentVersion := "v1"
	updateZip := "update.zip"

	http.HandleFunc("/current-version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, currentVersion)
	})

	http.HandleFunc("/update.zip", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, updateZip)
	})

	fmt.Println("Server started on : http://localhost:8080")
	fmt.Println("current version: ", currentVersion)
	http.ListenAndServe(":8080", nil)
}
