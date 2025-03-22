package main

import (
	"fmt"
)

func firstNonRepeatedChar(s string) byte {
	count := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if count[s[i]] == 1 {
			return s[i]
		}
	}
	return 0 // return 0 if no non-repeated character is found
}

func main() {
	str := "swiss"
	fmt.Println(string(firstNonRepeatedChar(str))) // w
}
