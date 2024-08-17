package main

import (
	"encoding/json"
	"go-exercises/design-patterns/behavioral/strategy2/payment"
	"log"
)

func main() {
	// Get Amount
	amount := 599
	// Create new payment processor
	processor := &payment.PaymentProcessor{}
	// Example pay using Credit Card
	processor.SetStrategy(&payment.CreditCardStrategy{
		NameOnCard: "Diwakar",
		Number:     "4111111111111111",
		ExpiryDate: "12/28",
	})

	paymentRes := processor.Checkout(float64(amount))

	paymentResBytes, _ := json.MarshalIndent(paymentRes, "", "   ")
	log.Printf("Credit PaymentRes: %v\n", string(paymentResBytes))

	// Example payment using Paypal
	processor.SetStrategy(&payment.PaypalStrategy{
		Email:    "test@email.com",
		Password: "secretpassword",
	})

	paypalPaymentRes := processor.Checkout(float64(amount))
	paypalPayResBytes, _ := json.MarshalIndent(paypalPaymentRes, "", "   ")
	log.Printf("Paypal PaymentRes: %v\n", string(paypalPayResBytes))
}
