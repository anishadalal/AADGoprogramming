package main

import "fmt"

// Struct to store student information
type Student struct {
	Name     string
	RollNo   int
	Division string
	College  string
}

func main() {
	// Creating an instance of Student
	student := Student{
		Name:     "Anisha",
		RollNo:   4222,
		Division: "A",
		College:  "BJS College",
	}

	// Printing student information
	fmt.Println("Student Name:", student.Name)
	fmt.Println("Roll Number:", student.RollNo)
	fmt.Println("Division:", student.Division)
	fmt.Println("College Name:", student.College)
}
