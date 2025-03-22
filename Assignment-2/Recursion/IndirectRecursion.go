package main

import "fmt"

// Function to check even number using indirect recursion
func isEven(n int) bool {
	if n == 0 {
		return true
	}
	return isOdd(n - 1)
}

// Function to check odd number using indirect recursion
func isOdd(n int) bool {
	if n == 0 {
		return false
	}
	return isEven(n - 1)
}

func main() {
	number := 5
	if isEven(number) {
		fmt.Printf("%d is even.\n", number)
	} else {
		fmt.Printf("%d is odd.\n", number)
	}
}
