package main

import "fmt"

// Function with named return variables
func sumAndDifference(a, b int) (sum, difference int) {
	sum = a + b
	difference = a - b
	return // No need to explicitly return variables as they're named
}

func main() {
	resultSum, resultDifference := sumAndDifference(10, 5)
	fmt.Printf("Sum: %d, Difference: %d\n", resultSum, resultDifference)
}
