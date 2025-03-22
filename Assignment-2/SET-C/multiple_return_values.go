package main

import "fmt"

// Function to perform division and return quotient and remainder
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	var num1, num2 int
	fmt.Print("Enter two numbers (dividend and divisor): ")
	fmt.Scanln(&num1, &num2)

	if num2 == 0 {
		fmt.Println("Division by zero is not allowed.")
		return
	}

	quotient, remainder := divide(num1, num2)
	fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
}
