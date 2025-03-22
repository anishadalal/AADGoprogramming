package main

import "fmt"

func main() {
	// Create a map to store student names as keys and their grades as values
	studentGrades := make(map[string]string)

	// Add student grades to the map
	studentGrades["Anisha"] = "A"
	studentGrades["Abantika"] = "B"
	studentGrades["Srijita"] = "A"

	// Display all student grades
	fmt.Println("Student Grades:")
	for student, grade := range studentGrades {
		fmt.Printf("%s: %s\n", student, grade)
	}

	// Retrieve and display the grade of a specific student
	var name string
	fmt.Print("\nEnter student name to search grade: ")
	fmt.Scan(&name)

	if grade, exists := studentGrades[name]; exists {
		fmt.Printf("%s's grade: %s\n", name, grade)
	} else {
		fmt.Println("Student not found!")
	}
}
