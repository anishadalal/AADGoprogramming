package main

import (
	"fmt"
	"sort"
)

// Define a struct for Student
type Student struct {
	RollNo     int
	Name       string
	Percentage float64
}

// Method to display student information
func displayStudents(students []Student) {
	fmt.Println("Student Information (Descending Order of Percentage):")
	for _, student := range students {
		fmt.Printf("Roll No: %d, Name: %s, Percentage: %.2f%%\n", student.RollNo, student.Name, student.Percentage)
	}
}

func main() {
	var n int

	// Input the number of students
	fmt.Print("Enter the number of students: ")
	fmt.Scan(&n)

	// Create a slice to store students
	students := make([]Student, n)

	// Input student information
	for i := 0; i < n; i++ {
		fmt.Printf("Enter details for student %d:\n", i+1)
		fmt.Print("Roll No: ")
		fmt.Scan(&students[i].RollNo)
		fmt.Print("Name: ")
		fmt.Scan(&students[i].Name)
		fmt.Print("Percentage: ")
		fmt.Scan(&students[i].Percentage)
	}

	// Sort students in descending order of percentage
	sort.Slice(students, func(i, j int) bool {
		return students[i].Percentage > students[j].Percentage
	})

	// Display sorted student information
	displayStudents(students)
}
