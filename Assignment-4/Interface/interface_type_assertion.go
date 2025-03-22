package main

import "fmt"

func describe(i interface{}) {
	value, ok := i.(string)
	if ok {
		fmt.Println("Value is a string:", value)
	} else {
		fmt.Println("Value is not a string")
	}
}

func main() {
	describe("Hello")
	describe(100)
}
