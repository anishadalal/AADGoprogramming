package main

import "fmt"

type Animal interface {
	Speak()
}

type Walker interface {
	Walk()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

func (d Dog) Walk() {
	fmt.Println("Dog is walking...")
}

func main() {
	var a Animal = Dog{}
	var w Walker = Dog{}

	a.Speak()
	w.Walk()
}
