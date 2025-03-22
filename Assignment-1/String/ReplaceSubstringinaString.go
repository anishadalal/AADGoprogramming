package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	result := strings.Replace(str, "world", "Go", -1)
	fmt.Println(result) // hello Go
}
