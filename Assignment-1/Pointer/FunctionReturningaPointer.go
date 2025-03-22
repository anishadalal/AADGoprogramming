package main

import "fmt"

func createPointer(val int) *int {
	return &val
}

func main() {
	ptr := createPointer(25)
	fmt.Println("Pointer address:", ptr)
	fmt.Println("Value pointed to:", *ptr)
}
