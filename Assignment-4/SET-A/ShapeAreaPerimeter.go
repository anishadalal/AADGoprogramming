package main

import (
	"fmt"
	"math"
)

// Shape interface with methods Area and Perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle type
type Circle struct {
	Radius float64
}

// Rectangle type
type Rectangle struct {
	Width, Height float64
}

// Implementing Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Implementing Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Implementing Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implementing Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Displaying results for Circle
	fmt.Printf("Circle: Area = %.2f, Perimeter = %.2f\n", circle.Area(), circle.Perimeter())

	// Displaying results for Rectangle
	fmt.Printf("Rectangle: Area = %.2f, Perimeter = %.2f\n", rectangle.Area(), rectangle.Perimeter())
}
