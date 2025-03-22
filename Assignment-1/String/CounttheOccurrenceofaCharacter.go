package main

import "fmt"

func countChar(s string, char byte) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == char {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countChar("hello", 'l')) // 2   
}
