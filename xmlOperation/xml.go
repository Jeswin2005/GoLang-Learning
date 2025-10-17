package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age,attr"`
}

func main() {
	// 1. Marshalling Struct to XML
	p := Person{
		Name: "Jeswin",
		Age:  20,
	}

	xmlBytes, err := xml.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling XML:", err)
		return
	}

	// Add XML header
	finalXML := xml.Header + string(xmlBytes)
	fmt.Println("Struct to XML:")
	fmt.Println(finalXML)

	// 1a. Write XML to file
	file, err := os.Create("person.xml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	file.WriteString(finalXML)
	fmt.Println("XML written to person.xml")

	// 2. Read XML from file
	fileRead, err := os.Open("person.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileRead.Close()

	var pFromFile Person
	decoder := xml.NewDecoder(fileRead)
	err = decoder.Decode(&pFromFile)
	if err != nil {
		fmt.Println("Error decoding XML from file:", err)
		return
	}

	fmt.Println("\nXML read from file to Struct:")
	fmt.Printf("%+v\n", pFromFile)

	// 3. Example of unmarshalling a hardcoded XML string
	xmlStr := `<Person age="25"><name>Alice</name></Person>`
	var p2 Person
	err = xml.Unmarshal([]byte(xmlStr), &p2)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}
	fmt.Println("\nXML to Struct (from string):")
	fmt.Printf("%+v\n", p2)
}
