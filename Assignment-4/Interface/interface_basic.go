package main

import "fmt"

// Define an interface
type Speaker interface {
	Speak()
}

// Implement interface with a struct
type Person struct {
	name string
}

func (p Person) Speak() {
	fmt.Println("Hello, my name is", p.name)
}

func main() {
	var s Speaker = Person{name: "Anisha"}
	s.Speak()
}
