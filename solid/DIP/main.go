/*
🧩 Dependency Inversion Principle Challenge
🧠 Problem:
You are building a payment system for an e-commerce platform.

🎯 Requirements:
The system must be able to process different types of payments (e.g., credit card, UPI, PayPal).

A receipt should be generated after successful payment.

The payment processor must also log the transaction.

Later, the team might add new payment methods or switch logging backends (e.g., from console to remote server) — but without modifying the core logic.

⚠️ Constraints:
Design your solution so that high-level components do not depend directly on low-level details.

Make your system open to extension, closed to modification.

Do not tightly couple concrete services like payment methods or loggers.

🏁 Deliverable:
Write the Go code for this system. Make sure it demonstrates:

Flexible payment method support

Pluggable logging

Receipt generation

Compliance with the Dependency Inversion Principle
*/

package main

import "fmt"

type Logger interface {
	Log(string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(message string) {
	fmt.Printf("Message from Console Logger '%s'", message)
}

type PaymentMethod interface {
	Pay(amount string) string
}

type UPI struct {
	logger  Logger
	receipt ReceiptGenerator
}

type CreditCard struct {
	logger  Logger
	receipt ReceiptGenerator
}

type Paypal struct {
	logger  Logger
	receipt ReceiptGenerator
}

func (upi UPI) Pay(amount float64) string {
	return fmt.Sprintf("Paid amount '%f' via UPI.", amount)
}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid amount '%f' via Credit Card.", amount)
}

func (pp Paypal) Pay(amount float64) string {
	return fmt.Sprintf("Paid amount '%f' via Paypal.", amount)
}

type ReceiptGenerator interface {
	Generate(string)
}

type OfflineReceiptGenerator struct{}
type OnlineReceiptGenerator struct{}

func (rg OfflineReceiptGenerator) Generate(message string) {
	fmt.Printf("Generated Offline Receipt - %s", message)
}

func (rg OnlineReceiptGenerator) Generate(message string) {
	fmt.Printf("Generated Online Receipt - %s", message)
}

func main() {
	logger := ConsoleLogger{}
	receipt := OnlineReceiptGenerator{}

	// Injecting abstractions
	payment := UPI{
		logger:  logger,
		receipt: receipt,
	}

	message := payment.Pay(500)
	payment.logger.Log(message)
	payment.receipt.Generate(message)
}
