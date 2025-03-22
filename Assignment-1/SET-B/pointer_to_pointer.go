package main

import "fmt"

func main() {
	var a int = 58     // A normal integer variable
	var ptr *int       // Pointer to an integer
	var ptrToPtr **int // Pointer to pointer

	// Assign the address of 'a' to 'ptr'
	ptr = &a

	// Assign the address of 'ptr' to 'ptrToPtr'
	ptrToPtr = &ptr

	// Print the value of 'a', the address of 'a', the value of ptr, and the value of ptrToPtr
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of ptr (Address of a):", ptr)
	fmt.Println("Value pointed by ptr (Value of a):", *ptr)
	fmt.Println("Value of ptrToPtr (Address of ptr):", ptrToPtr)
	fmt.Println("Value pointed by ptrToPtr (Address of a):", *ptrToPtr)
	fmt.Println("Value pointed by pointer pointed by ptrToPtr (Value of a):", **ptrToPtr)
}
