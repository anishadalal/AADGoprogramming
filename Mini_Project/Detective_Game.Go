package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Suspect struct holds information about each suspect
type Suspect struct {
	name      string
	alibi     string
	suspicion int
}

func main() {
	fmt.Println("🔍 Welcome to a Murder Mystery!")
	fmt.Println("A crime has been committed, and it's totally up to you to solve it.")
	fmt.Println("You have three main suspects: Tepi, Khedi, and Buchi.")

	// Creating suspects
	suspects := []Suspect{
		{"Tepi", "I was at the casino all night.", 2},
		{"Khedi", "I was at home cooking dinner.", 1},
		{"Buchi", "I was with my brother at the bar.", 3},
	}

	// Investigation phase
	fmt.Println("\nWhere do you want to investigate first?")
	fmt.Println("1. Casino\n2. Khedi's House\n3. The Bar")

	choice := getInput()
	switch choice {
	case "1":
		fmt.Println("\nYou visit the casino. The staff says Tepi was there but left early.")
		suspects[0].suspicion += 2
	case "2":
		fmt.Println("\nKhedi's house looks clean. A neighbor confirms seeing her at home.")
		suspects[1].suspicion -= 1
	case "3":
		fmt.Println("\nAt the bar, the bartender says Buchi arrived late and seemed nervous.")
		suspects[2].suspicion += 2
	default:
		fmt.Println("\nYou waste time wandering. The case gets colder.")
	}

	// Interrogation phase
	fmt.Println("\nNow, you interrogate the suspects. Who do you want to press harder?")
	fmt.Println("1. Tepi\n2. Khedi\n3. Buchi")

	choice = getInput()
	switch choice {
	case "1":
		fmt.Println("\nTepi gets defensive. His story has inconsistencies.")
		suspects[0].suspicion += 2
	case "2":
		fmt.Println("\nKhedi calmly repeats her alibi. Nothing seems off.")
	case "3":
		fmt.Println("\nBuchi sweats under pressure. His story changes slightly.")
		suspects[2].suspicion += 2
	}

	// Determining the most suspicious suspect
	mostSuspicious := suspects[0]
	for _, suspect := range suspects {
		if suspect.suspicion > mostSuspicious.suspicion {
			mostSuspicious = suspect
		}
	}

	// Final decision
	fmt.Println("\nBased on your investigation, who do you accuse?")
	fmt.Println("1. Tepi\n2. Khedi\n3. Buchi")

	choice = getInput()
	if (choice == "1" && mostSuspicious.name == "Tepi") ||
		(choice == "2" && mostSuspicious.name == "Mkhedi") ||
		(choice == "3" && mostSuspicious.name == "Buchi") {
		fmt.Printf("\nCongratulations! You correctly identified %s as the culprit. Case solved!\n", mostSuspicious.name)
	} else {
		fmt.Printf("\nYou accused the wrong person! The real culprit was %s. The mafia is still out there...\n", mostSuspicious.name)
	}
}

// getInput reads user input from the console
func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
