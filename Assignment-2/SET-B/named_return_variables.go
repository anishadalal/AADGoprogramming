package main

import "fmt"

// Function with named return variables
func calculateRectangleAreaAndPerimeter(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // Implicit return of named variables
}

func main() {
	var length, width float64
	fmt.Print("Enter the length of the rectangle: ")
	fmt.Scanln(&length)
	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scanln(&width)

	area, perimeter := calculateRectangleAreaAndPerimeter(length, width)
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", area, perimeter)
}
