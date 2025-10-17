package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type student struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Dept string  `json:"dept"`
	CGPA float64 `json:"cgpa"`
}

var students = []student{
	{ID: "1", Name: "Jason", Dept: "CSE", CGPA: 8.9},
	{ID: "2", Name: "Jeswin", Dept: "ECE", CGPA: 9.5},
	{ID: "3", Name: "Kishore", Dept: "MECH", CGPA: 7.5},
}

func main() {
	http.HandleFunc("/students", handleStudents)

	fmt.Println("Server starting on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	// GET METHOD
	case http.MethodGet:
		err := json.NewEncoder(w).Encode(students)
		if err != nil {
			log.Printf("Error encoding JSON: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		// POST METHOD
	case http.MethodPost:
		var newStudent student
		err := json.NewDecoder(r.Body).Decode(&newStudent)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		for _, s := range students {
			if s.ID == newStudent.ID {
				http.Error(w, "Student with this ID already exists", http.StatusConflict)
				return
			}
		}

		students = append(students, newStudent)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newStudent)

		// PUT METHOD
	case http.MethodPut:
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Student ID is required", http.StatusBadRequest)
			return
		}

		var updateStudent student
		err := json.NewDecoder(r.Body).Decode(&updateStudent)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		isUpdated := false

		for i, s := range students {
			if s.ID == id {
				students[i] = updateStudent
				isUpdated = true
				break
			}
		}

		if !isUpdated {
			http.Error(w, "Student not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(updateStudent)

		// DELETE METHOD
	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Student ID is required", http.StatusBadRequest)
			return
		}

		isDeleted := false

		for i, s := range students {
			if s.ID == id {
				students = append(students[:i], students[i+1:]...)
				isDeleted = true
				break
			}
		}

		if !isDeleted {
			http.Error(w, "Student not found", http.StatusNotFound)
			return
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
