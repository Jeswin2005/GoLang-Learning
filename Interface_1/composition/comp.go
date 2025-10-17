package main

import (
	"fmt"
	"math"
)

// Interfaces
type Area interface {
	area() float64
}

type Perimeter interface {
	perimeter() float64
}

type AreaPerimeter interface {
	Area
	Perimeter
}

// Circle implements Area
type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle implements AreaPerimeter
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	c := Circle{Radius: 5.0}
	fmt.Println("Circle radius:", c.Radius)
	fmt.Println("Circle area:", c.area())

	r := Rectangle{Width: 10, Height: 20}
	fmt.Println("Rectangle width:", r.Width, "height:", r.Height)
	fmt.Println("Rectangle area:", r.area())
	fmt.Println("Rectangle perimeter:", r.perimeter())
}
