package main

import "fmt"

func main() {
	var a, b int

	// Ask the user to enter two numbers
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&b)

	// Swap the numbers using arithmetic operations
	a = a + b
	b = a - b
	a = a - b

	// Display the swapped numbers
	fmt.Println("After swapping:")
	fmt.Println("First number:", a)
	fmt.Println("Second number:", b)
}
