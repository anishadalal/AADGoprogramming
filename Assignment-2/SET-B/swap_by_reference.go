package main

import "fmt"

// Function to swap two numbers using pointers
func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	var x, y int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&x)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&y)

	fmt.Printf("Before swap: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swap: x = %d, y = %d\n", x, y)
}
