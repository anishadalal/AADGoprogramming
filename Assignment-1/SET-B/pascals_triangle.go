package main

import "fmt"

// Function to generate and print Pascal's Triangle
func printPascalsTriangle(n int) {
	// 2D slice to store the values of Pascal's Triangle
	triangle := make([][]int, n)

	// Generate the triangle
	for i := 0; i < n; i++ {
		// Initialize the row
		triangle[i] = make([]int, i+1)

		// Set the first and last element of each row as 1
		triangle[i][0] = 1
		triangle[i][i] = 1

		// Calculate the inner elements
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	// Print the triangle
	for i := 0; i < n; i++ {
		// Print leading spaces for formatting
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}

		// Print the elements of the row
		for j := 0; j <= i; j++ {
			fmt.Print(triangle[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	var n int

	// Ask the user to enter the number of rows
	fmt.Print("Enter the number of rows for Pascal's Triangle: ")
	fmt.Scanln(&n)

	// Print Pascal's Triangle
	printPascalsTriangle(n)
}
