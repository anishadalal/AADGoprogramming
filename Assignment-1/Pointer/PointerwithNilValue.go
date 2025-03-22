package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("Pointer value before assignment:", ptr)

	if ptr == nil {
		fmt.Println("Pointer is nil.")
	}

	num := 50
	ptr = &num
	fmt.Println("Pointer value after assignment:", ptr)
	fmt.Println("Value at pointer:", *ptr)
}
