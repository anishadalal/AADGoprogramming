package main

import "fmt"

func main() {
	var password string
	const correctPassword = "golang123"

	for {
		fmt.Print("Enter the password: ")
		fmt.Scanln(&password)

		if password == correctPassword {
			fmt.Println("Access granted!")
			break
		} else {
			fmt.Println("Incorrect password. Try again.")
		}
	}
}
