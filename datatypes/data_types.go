package main

import "fmt"

const PI = 3.14

func main() {
	var a int = 5
	c := 7.5
	var name string = "Jeswin"
	var isLearning bool = true
	var radius float64

	fmt.Println("Name: ", name)
	fmt.Println("Learning Go: ", isLearning)
	fmt.Println("Enter radius: ")
	fmt.Scan(&radius)
	area := PI * radius * radius
	fmt.Println("Area of Circle: ", area)

	fmt.Println("\nData Types")
	fmt.Printf("a is of type %T with value %d\n", a, a)
	fmt.Printf("c is of type %T with value %.2f\n", c, c)

	var ch rune = 'a'
	fmt.Printf("ch is of type %T with value %c\n", ch, ch)

}
