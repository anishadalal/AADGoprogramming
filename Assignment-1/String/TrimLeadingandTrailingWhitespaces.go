package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "   hello   "
	result := strings.TrimSpace(str)
	fmt.Println(result) // hello
}
