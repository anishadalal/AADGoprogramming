package main

import "fmt"

// Student struct with Name, Age, and Grade fields
type Student struct {
	Name  string
	Age   int
	Grade string
}

// Show method for Student struct (pointer receiver)
func (s *Student) Show() {
	fmt.Printf("Student Name: %s\nAge: %d\nGrade: %s\n", s.Name, s.Age, s.Grade)
}

func main() {
	// Create an instance of Student using a pointer
	student := &Student{Name: "Anisha", Age: 20, Grade: "A"}

	// Call the Show method on the student instance
	student.Show()
}
