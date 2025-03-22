package main

import "fmt"

type Person struct {
	name string
	age  int
}

func printPerson(p Person) {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}

func main() {
	p := Person{name: "Anisha", age: 20}
	printPerson(p)
}
