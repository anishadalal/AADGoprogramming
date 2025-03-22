package main

import "fmt"

func main() {
	var input string

	for {
		fmt.Print("Enter a word (or type 'stop' to exit): ")
		fmt.Scanln(&input)

		if input == "stop" {
			fmt.Println("Exiting the program.")
			break
		}
		fmt.Printf("You entered: %s\n", input)
	}
}
