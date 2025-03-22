package main

import "fmt"

// Function to calculate factorial using recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	result := factorial(5)
	fmt.Printf("Factorial of 5 is: %d\n", result)
}
