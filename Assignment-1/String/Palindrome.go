package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	var result strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		result.WriteByte(s[i])
	}
	return result.String()
}

func isPalindrome(s string) bool {
	reversed := reverse(s)
	return s == reversed
}

func main() {
	str := "madam"
	fmt.Println(isPalindrome(str)) // true
}
