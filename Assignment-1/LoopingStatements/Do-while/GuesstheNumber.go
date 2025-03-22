package main

import "fmt"

func main() {
	var guess int

	for {
		fmt.Print("Guess the number: ")
		fmt.Scanln(&guess)

		if guess == 42 {
			fmt.Println("Correct! You guessed the number.")
			break
		} else {
			fmt.Println("Wrong guess. Try again!")
		}
	}
}
