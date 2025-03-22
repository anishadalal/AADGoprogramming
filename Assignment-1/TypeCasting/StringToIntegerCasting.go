package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "100"
	num, err := strconv.Atoi(str) // Casting string to integer
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("String to Integer:", num)
	}
}
