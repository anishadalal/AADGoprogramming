package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

type Robot struct{}

func (r Robot) Speak() {
	fmt.Println("Beep Boop!")
}

func communicate(s Speaker) {
	s.Speak()
}

func main() {
	communicate(Dog{})
	communicate(Robot{})
}
