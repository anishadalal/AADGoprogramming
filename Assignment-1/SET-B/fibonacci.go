package main

import "fmt"

func main() {
	var n int

	// Ask the user to enter the number of terms in the Fibonacci series
	fmt.Print("Enter the number of terms: ")
	fmt.Scanln(&n)

	// First two terms of the Fibonacci series
	a, b := 0, 1

	// Print the Fibonacci series up to n terms
	fmt.Println("Fibonacci series:")
	for i := 1; i <= n; i++ {
		fmt.Print(a, " ")

		// Calculate the next Fibonacci number
		a, b = b, a+b
	}
	fmt.Println()
}
