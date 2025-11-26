package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	currentVersion := "v2"

	updateZip := "update.zip"

	http.HandleFunc("/current-version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, currentVersion)
	})

	http.HandleFunc("/update.zip", func(w http.ResponseWriter, r *http.Request) {

		if _, err := os.Stat(updateZip); os.IsNotExist(err) {
			http.Error(w, "update.zip not found on server", 404)
			return
		}

		http.ServeFile(w, r, updateZip)
	})

	fmt.Println("Update Server running at http://localhost:8080")
	fmt.Println("Serving ZIP:", updateZip)
	fmt.Println("Current version exposed:", currentVersion)
	http.ListenAndServe(":8080", nil)
}
