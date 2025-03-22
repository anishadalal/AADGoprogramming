package main

import (
	"fmt"
	"strings"
)

func countWords(s string) int {
	words := strings.Fields(s)
	return len(words)
}

func main() {
	str := "Hello Go world"
	fmt.Println(countWords(str)) // 3
}
