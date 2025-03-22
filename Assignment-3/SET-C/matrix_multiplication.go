package main

import "fmt"

func main() {
	var r1, c1, r2, c2 int

	fmt.Print("Enter rows and columns for the first matrix: ")
	fmt.Scan(&r1, &c1)
	fmt.Print("Enter rows and columns for the second matrix: ")
	fmt.Scan(&r2, &c2)

	if c1 != r2 {
		fmt.Println("Matrix multiplication is not possible with the given dimensions.")
		return
	}

	// Accept first matrix
	fmt.Println("Enter elements of the first matrix:")
	mat1 := make([][]int, r1)
	for i := 0; i < r1; i++ {
		mat1[i] = make([]int, c1)
		for j := 0; j < c1; j++ {
			fmt.Scan(&mat1[i][j])
		}
	}

	// Accept second matrix
	fmt.Println("Enter elements of the second matrix:")
	mat2 := make([][]int, r2)
	for i := 0; i < r2; i++ {
		mat2[i] = make([]int, c2)
		for j := 0; j < c2; j++ {
			fmt.Scan(&mat2[i][j])
		}
	}

	// Multiply matrices
	result := make([][]int, r1)
	for i := 0; i < r1; i++ {
		result[i] = make([]int, c2)
		for j := 0; j < c2; j++ {
			for k := 0; k < c1; k++ {
				result[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}

	// Display result
	fmt.Println("Resultant Matrix after Multiplication:")
	for _, row := range result {
		fmt.Println(row)
	}
}
