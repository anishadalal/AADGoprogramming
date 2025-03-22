package main

import "fmt"

func main() {
	var num1, num2 float64
	var choice int

	// Accept two numbers from the user
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	// Display the operation options
	fmt.Println("Choose an operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	// Accept user's choice of operation
	fmt.Print("Enter your choice (1-4): ")
	fmt.Scanln(&choice)

	// Perform the selected operation
	switch choice {
	case 1:
		fmt.Printf("Result: %.2f\n", num1+num2)
	case 2:
		fmt.Printf("Result: %.2f\n", num1-num2)
	case 3:
		fmt.Printf("Result: %.2f\n", num1*num2)
	case 4:
		if num2 != 0 {
			fmt.Printf("Result: %.2f\n", num1/num2)
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	default:
		fmt.Println("Invalid choice. Please select a valid operation (1-4).")
	}
}
