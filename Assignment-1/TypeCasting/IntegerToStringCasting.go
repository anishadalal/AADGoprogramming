package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 123
	str := strconv.Itoa(num) // Casting integer to string
	fmt.Println("Integer to String:", str)
}
