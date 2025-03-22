package main

import "fmt"

// Interface with different types
type Shape interface{}

func describeShape(s Shape) {
	// Type switch to handle different types of shapes
	switch v := s.(type) {
	case int:
		fmt.Println("It's an integer:", v)
	case string:
		fmt.Println("It's a string:", v)
	case float64:
		fmt.Println("It's a float64:", v)
	default:
		fmt.Println("Unknown type:", v)
	}
}

func main() {
	// Test with different types
	describeShape(10)       // int
	describeShape("Circle") // string
	describeShape(3.14)     // float64
	describeShape(true)     // unknown type
}
