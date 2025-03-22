package main

import "fmt"

// Structure to hold student details
type Student struct {
	RollNo  int
	Name    string
	Mark1   float64
	Mark2   float64
	Mark3   float64
	Total   float64
	Average float64
}

func main() {
	var n int
	fmt.Print("Enter the number of students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for student %d:\n", i+1)
		fmt.Print("Roll Number: ")
		fmt.Scan(&students[i].RollNo)
		fmt.Print("Name: ")
		fmt.Scan(&students[i].Name)
		fmt.Print("Mark 1: ")
		fmt.Scan(&students[i].Mark1)
		fmt.Print("Mark 2: ")
		fmt.Scan(&students[i].Mark2)
		fmt.Print("Mark 3: ")
		fmt.Scan(&students[i].Mark3)

		// Calculating total and average
		students[i].Total = students[i].Mark1 + students[i].Mark2 + students[i].Mark3
		students[i].Average = students[i].Total / 3
	}

	// Displaying student details
	fmt.Println("\nStudent Details:")
	for _, student := range students {
		fmt.Printf("Roll No: %d, Name: %s, Total: %.2f, Average: %.2f\n",
			student.RollNo, student.Name, student.Total, student.Average)
	}
}
