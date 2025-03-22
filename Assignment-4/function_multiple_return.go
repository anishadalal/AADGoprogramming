package main

import "fmt"

// Function that returns multiple values
func calculate(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	sum, product := calculate(5, 3)
	fmt.Println("Sum:", sum, "Product:", product)
}
