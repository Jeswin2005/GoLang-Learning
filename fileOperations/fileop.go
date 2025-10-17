package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Step 1: Create a file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File created successfully!")

	// Step 2: Write data to file
	content := "Hello im jeswin..\n How are you all\n"
	bytesWritten, err := file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("Wrote %d bytes to file.\n", bytesWritten)

	// Step 3: Reopen file for reading
	readFile, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer readFile.Close()

	fmt.Println("\nFile content (using bufio.Reader):")
	reader := bufio.NewReader(readFile)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Print(line)
	}

	// step 4 : modify a string in the file with given string
	fmt.Print("Enter string to find: ")
	var oldStr string
	fmt.Scanln(&oldStr)

	fmt.Print("Enter string to replace with: ")
	var newStr string
	fmt.Scanln(&newStr)

	// Read the entire file content
	readFile.Seek(0, 0)
	fileContent, err := io.ReadAll(readFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Replace old string with new string
	modifiedContent := string(fileContent)
	modifiedContent = strings.ReplaceAll(modifiedContent, oldStr, newStr)

	// Write the modified content back
	err = os.WriteFile("example.txt", []byte(modifiedContent), 0644)
	if err != nil {
		fmt.Println("Error writing modified content:", err)
		return
	}

	fmt.Println("File modified successfully!")

	// Optional: display new file content
	fmt.Println("\nModified File Content:")
	fmt.Println(modifiedContent)

}
