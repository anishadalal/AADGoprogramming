package main

import "fmt"

// Function that returns sum, difference, and product
func calculate(a, b int) (int, int, int) {
	sum := a + b
	difference := a - b
	product := a * b
	return sum, difference, product
}

func main() {
	sum, diff, prod := calculate(10, 5)
	fmt.Printf("Sum: %d, Difference: %d, Product: %d\n", sum, diff, prod)
}
