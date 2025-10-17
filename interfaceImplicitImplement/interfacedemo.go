package main

import (
	"fmt"
	"math"
)

// Interface Definition (The Contract)
type Shaper interface {
	Area() 
	Perimeter() 
}

type Circle struct {
	Radius float64
}

func (c Circle) Area_Circle() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle (Implements Shaper)
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

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	fmt.Println("Demonstrating Polymorphism with Interfaces and Type Assertion")

	fmt.Println("circle area")
	c.Area_Circle()
	fmt.Println("This is a Circle with Radius:", c.Radius)

	fmt.Println("rectangle area")
	r.Area()

	fmt.Println("rectangle perimeter")
	r.Perimeter()
	fmt.Println("This is a Rectangle with Width:", r.Width, "and Height:", r.Height)

}
