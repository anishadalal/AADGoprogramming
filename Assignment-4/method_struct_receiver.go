package main

import "fmt"

type Person struct {
	name string
}

func (p Person) greet() {
	fmt.Println("Hello,", p.name)
}

func main() {
	p := Person{name: "Anisha"}
	p.greet()
}
