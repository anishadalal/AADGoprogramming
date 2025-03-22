package main

import (
	"fmt"
	"strings"
)

func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

func main() {
	fmt.Println(toUpperCase("hello"))
}
