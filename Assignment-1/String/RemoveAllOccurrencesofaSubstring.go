package main

import (
	"fmt"
	"strings"
)

func removeSubstring(s, substr string) string {
	return strings.ReplaceAll(s, substr, "")
}

func main() {
	str := "Hello, world! World is great."
	fmt.Println(removeSubstring(str, "world")) // Hello, ! World is great.
}
