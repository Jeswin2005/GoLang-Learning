package main

import "fmt"

func main() {
	// 1. The 'for' Loop for initialize; cond; increment
	fmt.Println("1. (for i=0; i<3; i++)")
	for i := 0; i < 3; i++ {
		fmt.Printf("Iteration %d\n", i)
	}
	fmt.Println()

	// 2. The 'for' Loop as a While-style Loop (omitting init/post)
	fmt.Println("2. While-style Loop (for condition)")
	num := 10
	for num > 0 {
		fmt.Printf("Num is %d\n", num)
		num -= 4 // Decrement the number
	}
	fmt.Println()

	// 3. The for range
	fmt.Println("3. Range Loop (for index, value := range slice)")
	cities := []string{"Tokyo", "London", "New York"}
	for i, city := range cities {
		fmt.Printf("City %d: %s\n", i, city)
	}
	fmt.Println()

	// 4. The for range without index
	fmt.Println("4. Range Loop (for index, value := range slice) without index")
	fruits := []string{"apple", "banana", "mango"}
	for _, fruit := range fruits {
		fmt.Printf("fruit: %s\n", fruit)
	}
	fmt.Println()

	// 5. The Infinite Loop (for {}) with a break
	fmt.Println("5. Infinite Loop (for {}) with break")
	count := 0
	for {
		if count >= 2 {
			fmt.Println("Break condition met.")
			break
		}
		fmt.Printf("Count is currently %d\n", count)
		count++
	}
}
