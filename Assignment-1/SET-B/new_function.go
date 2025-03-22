package main

import "fmt"

func main() {
	// Use the 'new' function to create a pointer to an integer
	ptr := new(int)

	// Assign a value to the memory location
	*ptr = 42

	// Print the value stored at the memory location
	fmt.Println("Value of the integer:", *ptr)

	// Print the memory address of the variable
	fmt.Println("Memory address of the integer:", ptr)

	// Use 'new' function to create a pointer to a float
	floatPtr := new(float64)
	*floatPtr = 3.14

	// Print the value of the float
	fmt.Println("Value of the float:", *floatPtr)

	// Print the memory address of the float
	fmt.Println("Memory address of the float:", floatPtr)
}
