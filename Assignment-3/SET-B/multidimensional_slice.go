package main

import "fmt"

func main() {
	// Creating a multidimensional slice
	multiSlice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Printing the multidimensional slice
	fmt.Println("Multidimensional Slice:")
	for _, row := range multiSlice {
		fmt.Println(row)
	}
}
