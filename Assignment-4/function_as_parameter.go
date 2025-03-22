package main

import "fmt"

// Function that takes another function as a parameter
func operation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Simple addition function
func add(x, y int) int {
	return x + y
}

func main() {
	result := operation(5, 3, add)
	fmt.Println("Result:", result)
}
