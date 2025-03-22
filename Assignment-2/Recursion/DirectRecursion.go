package main

import "fmt"

// Function to calculate sum of natural numbers using direct recursion
func sum(n int) int {
	if n == 1 {
		return 1
	}
	return n + sum(n-1)
}

func main() {
	result := sum(5)
	fmt.Printf("Sum of first 5 natural numbers: %d\n", result)
}
