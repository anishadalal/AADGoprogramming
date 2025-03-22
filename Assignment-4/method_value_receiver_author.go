package main

import "fmt"

type Author struct {
	name string
}

func (a Author) printName() {
	fmt.Println("Author:", a.name)
}

func main() {
	a := Author{name: "Satyajit Ray"}
	a.printName()
}
