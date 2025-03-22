package main

import (
	"fmt"
)

func main() {
	var tasks []string
	var input string

	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. Show Tasks")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scan(&input)

		switch input {
		case "1":
			fmt.Print("Enter task: ")
			var task string
			fmt.Scan(&task)
			tasks = append(tasks, task)
		case "2":
			fmt.Println("To-Do List:")
			if len(tasks) == 0 {
				fmt.Println("No tasks yet.")
			} else {
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}
		case "3":
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
