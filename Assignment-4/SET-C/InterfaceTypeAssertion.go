package main

import "fmt"

// Define an interface
type MyInterface interface {
	Display()
}

// Define a struct
type Person struct {
	Name string
	Age  int
}

// Implement the interface method for Person
func (p Person) Display() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func main() {
	// Create a variable of type MyInterface
	var i MyInterface

	// Assign a Person struct to the interface
	i = Person{Name: "Anisha", Age: 20}

	// Use type assertion to extract the value of the underlying type
	if p, ok := i.(Person); ok {
		fmt.Println("Using Type Assertion:")
		fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
	} else {
		fmt.Println("Type assertion failed")
	}

	// Display using the interface method
	fmt.Println("Using Interface Method:")
	i.Display()
}
