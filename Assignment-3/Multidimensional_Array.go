package main

import "fmt"

func main() {
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Multidimensional Array:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
