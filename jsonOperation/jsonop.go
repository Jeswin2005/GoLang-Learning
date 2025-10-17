package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 1. Marshalling Struct to JSON
	p := Person{
		Name: "Jeswin",
		Age:  20,
	}

	// Convert struct to JSON
	jsonBytes, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println("Struct to JSON:")
	fmt.Println(string(jsonBytes))

	// 1a. Write JSON to file
	file, err := os.Create("person.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	file.Write(jsonBytes)
	fmt.Println("JSON written to person.json")

	// 2. Read JSON from file
	fileRead, err := os.Open("person.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileRead.Close()

	var pFromFile Person
	decoder := json.NewDecoder(fileRead)
	err = decoder.Decode(&pFromFile)
	if err != nil {
		fmt.Println("Error decoding JSON from file:", err)
		return
	}
	fmt.Println("\nJSON read from file to Struct:")
	fmt.Printf("%+v\n", pFromFile)

	// 3. Dynamic JSON using map
	dynamicJSON := `{"title":"Go Developer","skills":["Go","Python","JS"],"experience":3}`
	var data map[string]interface{}
	err = json.Unmarshal([]byte(dynamicJSON), &data)
	if err != nil {
		fmt.Println("Error unmarshalling dynamic JSON:", err)
		return
	}

	fmt.Println("\nDynamic JSON to Map:")
	fmt.Println(data)

	fmt.Println("Title:", data["title"])
	fmt.Println("Experience:", data["experience"])

	skills := data["skills"].([]interface{})
	fmt.Print("Skills: ")
	for _, skill := range skills {
		fmt.Print(skill.(string), " ")
	}
	fmt.Println()
}
