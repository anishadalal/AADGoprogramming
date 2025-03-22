package main

import (
	"fmt"
	"strconv"
)

func main() {
	numFloat := 3.14159
	str := strconv.FormatFloat(numFloat, 'f', 2, 64) // Casting float to string
	fmt.Println("Float to String:", str)
}
