package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
}

// Implement interface with a struct
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	var s Shape = Rectangle{width: 5, height: 3}
	fmt.Println("Area:", s.Area())
}
