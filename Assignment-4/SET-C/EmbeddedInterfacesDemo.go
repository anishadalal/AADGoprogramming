package main

import "fmt"

// Define an interface for basic information
type BasicInfo interface {
	GetName() string
	GetAge() int
}

// Define another interface for additional details
type AdditionalInfo interface {
	GetID() int
}

// Define an embedded interface combining the above two
type FullInfo interface {
	BasicInfo
	AdditionalInfo
}

// Define a struct
type Person struct {
	Name string
	Age  int
	ID   int
}

// Implement BasicInfo methods for Person
func (p Person) GetName() string {
	return p.Name
}

func (p Person) GetAge() int {
	return p.Age
}

// Implement AdditionalInfo method for Person
func (p Person) GetID() int {
	return p.ID
}

// Function to display full information
func displayFullInfo(info FullInfo) {
	fmt.Println("Full Information:")
	fmt.Printf("Name: %s, Age: %d, ID: %d\n", info.GetName(), info.GetAge(), info.GetID())
}

func main() {
	// Create a Person instance
	person := Person{Name: "Anisha", Age: 20, ID: 4222}

	// Display full information using the embedded interface
	displayFullInfo(person)
}
