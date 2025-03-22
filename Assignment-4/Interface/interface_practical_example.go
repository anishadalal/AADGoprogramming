package main

import "fmt"

type Payment interface {
	Pay(amount float64)
}

type CreditCard struct {
	name string
}

func (c CreditCard) Pay(amount float64) {
	fmt.Println("Paid", amount, "using Credit Card of", c.name)
}

type PayPal struct {
	email string
}

func (p PayPal) Pay(amount float64) {
	fmt.Println("Paid", amount, "using PayPal account:", p.email)
}

func main() {
	var p Payment = CreditCard{name: "Anisha"}
	p.Pay(100.50)

	p = PayPal{email: "anisha@example.com"}
	p.Pay(50.75)
}
