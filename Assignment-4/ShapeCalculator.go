package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Rectangle struct
type Rectangle struct {
	Length, Width float64
}

// Implement Area and Perimeter methods for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Implement Area and Perimeter methods for Rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Function to print shape details
func printShapeDetails(s Shape) {
	fmt.Printf("Shape details:\n")
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	// Create Circle and Rectangle instances
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Length: 10, Width: 4}

	// Print details for Circle
	fmt.Println("Circle:")
	printShapeDetails(circle)

	// Print details for Rectangle
	fmt.Println("\nRectangle:")
	printShapeDetails(rectangle)
}
