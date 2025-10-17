package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Strings (Immutability)
	fmt.Println("\n1. String Basics")
	s1 := "Go is fast!"
	s2 := "Concurrency"
	fmt.Printf("Initial String: \"%s\"\n", s1)

	// 2. String Length
	fmt.Println("\n2. String Length")
	// len() counts bytes. For basic letters, bytes = characters.
	fmt.Printf("Byte Length (len()): %d\n", len(s1))
	// To count actual characters (runes, good for emojis/Unicode)
	fmt.Printf("Character Length (runes): %d\n", len([]rune(s1)))

	// 3. String Concatenation
	fmt.Println("\n3. Concatenation")
	// Use the '+' operator to join strings.
	s3 := s1 + " " + s2
	fmt.Printf("Joined with '+': \"%s\"\n", s3)

	// 4. Compare Strings
	fmt.Println("\n4. Comparison")
	// Use '==' to check if two strings are exactly the same.
	isEqual := s1 == "Go is fast!"
	fmt.Printf("s1 == \"Go is fast!\": %t\n", isEqual)

	// 5. Split String
	fmt.Println("\n5. Splitting")
	// strings.Split breaks a string into a slice of strings based on a separator.
	parts := strings.Split(s1, " ")
	fmt.Printf("Split by space: %v\n", parts)

	// 6. Substring Extraction (Slicing)
	fmt.Println("\n6. Substring Slicing")
	// Use [start:end]
	word := s1[0:2]
	fmt.Printf("s1[0:2]: \"%s\"\n", word)
	rest := s1[3:]
	fmt.Printf("s1[3:]: \"%s\"\n", rest)

	// 7. String Replacement
	fmt.Println("\n7. Replacement")
	s5 := "cat cat cat"
	s5 = strings.ReplaceAll(s5, "cat", "dog")
	fmt.Printf("ReplaceAll: \"%s\"\n", s5)

	// 8. String Interpolation
	fmt.Println("\n8. Interpolation (Formatting)")
	// Go uses fmt.Sprintf to insert variables into a formatted string template.
	count := 3
	formatted := fmt.Sprintf("We have %d dogs now.", count)
	fmt.Printf("fmt.Sprintf: \"%s\"\n", formatted)
}
