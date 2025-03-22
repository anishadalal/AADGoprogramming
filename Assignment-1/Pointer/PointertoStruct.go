package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func changeName(p *Person) {
	p.Name = "Anisha"
}

func main() {
	person := Person{Name: "Abantika", Age: 20}
	fmt.Println("Before change:", person)

	changeName(&person) // Pass pointer to struct
	fmt.Println("After change:", person)
}
