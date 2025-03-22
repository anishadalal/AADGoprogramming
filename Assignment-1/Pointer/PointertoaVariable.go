package main

import "fmt"

func main() {
	var num int = 58
	var ptr *int = &num

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", ptr)
	fmt.Println("Value at pointer ptr:", *ptr)

	*ptr = 100 // Modify value using pointer
	fmt.Println("New value of num:", num)
}
