package main

import "fmt"

// Function to add two numbers
func addNumbers(a int, b int) int {
	return a + b
}

func main() {
	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	result := addNumbers(num1, num2)
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, result)
}
