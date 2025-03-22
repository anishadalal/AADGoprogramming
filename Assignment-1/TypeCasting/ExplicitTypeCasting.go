package main

import "fmt"

func main() {
	var numFloat float64 = 3.14
	var numInt int = int(numFloat) // Explicit type conversion
	fmt.Printf("Float: %.2f, Integer: %d\n", numFloat, numInt)
}
