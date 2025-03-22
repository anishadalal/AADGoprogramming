package main

import "fmt"

type Animal interface {
	Speak()
}

type LivingBeing interface {
	Animal // Embedded interface
	Breathe()
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func (c Cat) Breathe() {
	fmt.Println("Cat is breathing...")
}

func main() {
	var l LivingBeing = Cat{}
	l.Speak()
	l.Breathe()
}
