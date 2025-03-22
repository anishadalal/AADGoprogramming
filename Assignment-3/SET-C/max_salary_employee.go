package main

import "fmt"

type Employee struct {
	Eno    int
	Ename  string
	Salary float64
}

func main() {
	var n int
	fmt.Print("Enter the number of employees: ")
	fmt.Scan(&n)

	employees := make([]Employee, n)

	// Accept employee records
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for employee %d:\n", i+1)
		fmt.Print("Employee Number: ")
		fmt.Scan(&employees[i].Eno)
		fmt.Print("Employee Name: ")
		fmt.Scan(&employees[i].Ename)
		fmt.Print("Salary: ")
		fmt.Scan(&employees[i].Salary)
	}

	// Find maximum salary
	maxSalary := employees[0]
	for _, emp := range employees {
		if emp.Salary > maxSalary.Salary {
			maxSalary = emp
		}
	}

	// Display employee with maximum salary
	fmt.Println("\nEmployee with Maximum Salary:")
	fmt.Printf("Eno: %d, Name: %s, Salary: %.2f\n", maxSalary.Eno, maxSalary.Ename, maxSalary.Salary)
}
