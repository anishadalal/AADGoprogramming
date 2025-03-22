package main

import "fmt"

func printValue(v interface{}) {
	fmt.Println("Value:", v)
}

func main() {
	printValue(42)
	printValue("Hello")
	printValue(3.14)
}
