package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
}

func (p *Person) Describe() {
	fmt.Println("Person's name is", p.name)
}

func main() {
	var d Describer = &Person{name: "Anisha"}
	d.Describe()
}
