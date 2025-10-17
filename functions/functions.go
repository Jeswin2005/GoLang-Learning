package main

import (
	"fmt"
	"math"
)

// 1. Call by value
func swap(a int, b int) {
	temp := a
	a = b
	b = temp
}

// 2. Call by referenece
func swap1(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// 3. function closure
func Counter() func() int {
	count := 1
	return func() int {
		count++
		return count
	}
}

// 4. Method
// circle struct
type Circle struct {
	x, y, radius float64
}

// method for a circle
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func main() {
	a := 10
	b := 20

	// 1. Call by value
	fmt.Println("1. Call by value")
	fmt.Printf("before swap: %d %d", a, b)
	swap(a, b)
	fmt.Printf("\nafter swap: %d %d", a, b)

	// 2. Call by reference
	fmt.Println("\n2. Call by reference")
	fmt.Printf("before swap: %d %d", a, b)
	swap1(&a, &b)
	fmt.Printf("\nafter swap: %d %d", a, b)

	// 3, function as value
	fmt.Println("\n3. function as value")
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Printf("Sum: %d", sum(a, b))

	// 4. function closure
	incrementor := Counter()
	incrementor1 := Counter()
	fmt.Println("\n4. function closure")
	fmt.Printf("Count 1:%d", incrementor())
	fmt.Printf("\nCount 2:%d", incrementor())
	fmt.Printf("\nCount 3:%d", incrementor())
	fmt.Printf("\nCount 1:%d", incrementor1())

	// 5. Method (specific to a struct)
	fmt.Println("\n5. Method")
	circle := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("Circle area: %f", circle.area())
}
