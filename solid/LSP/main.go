/*
🧪 Task: LSP Challenge
You’re designing a payment processing system.

Your requirements:

There are multiple payment types: Credit Card, PayPal, and Gift Card.

All support a method Pay(amount float64).

Only Credit Card and PayPal support refunds.

Gift Cards cannot be refunded.

🎯 Your Goal:
Design the interfaces and types so that:

Substitution works safely

You do not violate LSP

Don’t force all payment types to implement Refund()

Write a ProcessPayment(p PaymentMethod) function that works for all types
*/

package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64)
}

type Refund interface {
	Refund()
}

type CreditCard struct{}
type Paypal struct{}
type GiftCard struct{}

func (cd CreditCard) Pay(amount float64) {
	fmt.Println("Processing payment via Credit Card")
}

func (cd CreditCard) Refund() {
	fmt.Println("Processing refund via Credit Card")
}

func (p Paypal) Pay(amount float64) {
	fmt.Println("Processing payment via p Paypal")
}

func (p Paypal) Refund() {
	fmt.Println("Processing refund via p Paypal")
}

func (gc GiftCard) Pay(amount float64) {
	fmt.Println("Processing payment via Gift Card")
}

func ProcessPayment(p PaymentMethod, amount float64) {
	p.Pay(amount)
}

func main() {
	fmt.Println("Welcome to LSP")
}
