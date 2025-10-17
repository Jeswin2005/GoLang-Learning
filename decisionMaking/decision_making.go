package main

import "fmt"

func main() {
	// 1. If Statement
	num := 10
	if num > 5 {
		fmt.Println("Number is greater than 5")
	}

	// 2. If-Else Statement
	if num%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// 3. Nested If Statements
	if num > 0 {
		if num < 20 {
			fmt.Println("Number is positive and less than 20")
		}
	}

	// 4. Switch Statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}

}
