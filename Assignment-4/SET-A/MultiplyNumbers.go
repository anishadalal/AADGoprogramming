package main

import "fmt"

// Method to multiply two numbers
func multiply(a, b int) int {
	return a * b
}

func main() {
	// Define two numbers
	num1 := 6
	num2 := 7

	// Call the multiply method and print the result
	result := multiply(num1, num2)
	fmt.Printf("Multiplication of %d and %d is: %d\n", num1, num2, result)
}
