/*
🎯 Problem Statement (Interview Style)
You’re building a modular payment system for an e-commerce platform.

Your internal system is designed to interact with payment providers using this interface:

type PaymentProcessor interface {
	Pay(amount float64) string // amount is in INR (rupees)
}
You are now asked to integrate a third-party payment provider — Razorpay — into this system.

However, the Razorpay SDK you must use has the following struct and method:

type Razorpay struct{}

// Accepts amount in paise (int), not rupees (float64)
func (r Razorpay) MakePaymentINR(paise int) string
❗ Constraints
You are not allowed to change the Razorpay struct or its method.

The Razorpay struct and method are assumed to be part of a closed-source vendor library.

Your existing business logic depends on the PaymentProcessor interface. You must continue using it as-is.

Your system expects payments in rupees as float64.
*/

package main

import "fmt"

type PaymentProcessor interface {
	Pay(amount float64) string // amount in INR
}

// Razorpay: third-party SDK
type Razorpay struct{}

func (r Razorpay) MakePaymentINR(paise int) string {
	return fmt.Sprintf("Payment done of ₹%.2f", float64(paise)/100)
}

// Adapter
type RazorpayAdapter struct {
	razorpay Razorpay
}

func (rp RazorpayAdapter) Pay(amount float64) string {
	paise := int(amount * 100)
	return rp.razorpay.MakePaymentINR(paise)
}

func NewRazorpayAdapter(r Razorpay) PaymentProcessor {
	return RazorpayAdapter{razorpay: r}
}

func main() {
	pp := NewRazorpayAdapter(Razorpay{})
	fmt.Println(pp.Pay(10))
}
