package main

import "fmt"

// Function that returns quotient and remainder
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	quotient, remainder := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
}
