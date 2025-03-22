package main

import "fmt"

// Function to check if a number is a palindrome
func isPalindrome(num int) bool {
	original := num
	reversed := 0

	for num > 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}

	return original == reversed
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if number < 0 {
		fmt.Println("Negative numbers cannot be palindromes.")
	} else if isPalindrome(number) {
		fmt.Printf("%d is a palindrome.\n", number)
	} else {
		fmt.Printf("%d is not a palindrome.\n", number)
	}
}
