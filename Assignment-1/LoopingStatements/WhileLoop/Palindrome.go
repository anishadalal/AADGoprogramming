package main

import "fmt"

func main() {
	var num, originalNum, reverse int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	originalNum = num
	for num > 0 {
		reverse = reverse*10 + num%10
		num /= 10
	}

	if originalNum == reverse {
		fmt.Printf("%d is a Palindrome.\n", originalNum)
	} else {
		fmt.Printf("%d is not a Palindrome.\n", originalNum)
	}
}
