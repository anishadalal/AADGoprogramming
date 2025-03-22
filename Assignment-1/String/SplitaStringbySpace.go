package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	result := strings.Split(str, " ")
	fmt.Println(result) // [hello world]
}
