package main

import "fmt"

func applyOperation(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("Sum of 5 and 3:", applyOperation(5, 3, add))
}
