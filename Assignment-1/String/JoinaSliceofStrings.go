package main

import (
	"fmt"
	"strings"
)

func main() {
	strSlice := []string{"hello", "world"}
	result := strings.Join(strSlice, " ")
	fmt.Println(result) // hello world
}
