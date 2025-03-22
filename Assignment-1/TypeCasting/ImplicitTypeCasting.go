package main

import "fmt"

func main() {
	var numInt int = 10
	var numFloat float64 = float64(numInt) // Implicit type conversion
	fmt.Printf("Integer: %d, Float: %.2f\n", numInt, numFloat)
}
