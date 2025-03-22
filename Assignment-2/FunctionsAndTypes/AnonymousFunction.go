package main

import "fmt"

func main() {
	cube := func(num int) int {
		return num * num * num
	}
	fmt.Println("Cube of 3:", cube(3))
}
