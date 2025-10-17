package main

import (
	"fmt"
	"math"
)

// Interface Definition (The Contract)
type Shaper interface {
	Area() float64
	Perimeter() float64
}

// Concrete Type 1: Circle (Implements Shaper)
type Circle struct {
	Radius float64
}

// Method 1 for Circle: Calculates Area
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Method 2 for Circle: Calculates Perimeter (Circumference)
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Concrete Type 2: Rectangle (Implements Shaper)
type Rectangle struct {
	Width, Height float64
}

// Method 1 for Rectangle: Calculates Area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method 2 for Rectangle: Calculates Perimeter
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Polymorphic Function (Accepts the Interface)
func printShapeDetails(s Shaper) {
	fmt.Printf("Shape Details: (Type: %T)\n", s)
	fmt.Printf("  Area: %.2f\n", s.Area())
	fmt.Printf("  Perimeter: %.2f\n", s.Perimeter())

	// Type Assertion Section
	// Check if s is a Circle
	if c, ok := s.(Circle); ok {
		fmt.Println("  (Type Assertion) This is a Circle with Radius:", c.Radius)
	}

	// Check if s is a Rectangle
	if r, ok := s.(Rectangle); ok {
		fmt.Println("  (Type Assertion) This is a Rectangle with Width:", r.Width, "and Height:", r.Height)
	}
}

func main() {
	// Create instances of the two distinct types
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	fmt.Println("Demonstrating Polymorphism with Interfaces and Type Assertion")

	printShapeDetails(c)
	fmt.Println()
	printShapeDetails(r)
}
