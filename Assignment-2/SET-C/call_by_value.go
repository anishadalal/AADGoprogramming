package main

import "fmt"

// Function to demonstrate call by value
func incrementValue(num int) {
	num += 1
	fmt.Printf("Inside function, value: %d\n", num)
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	fmt.Printf("Before function call, value: %d\n", number)
	incrementValue(number)
	fmt.Printf("After function call, value: %d\n", number) // Unchanged
}
