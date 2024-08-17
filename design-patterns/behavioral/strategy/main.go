package main

import (
	"encoding/json"
	"go-exercises/design-patterns/behavioral/strategy/payment"
	"log"
)

func main() {
	amount := float32(399)

	// Add Credit Card Details
	creditCard := payment.CreditCard{
		NameOnCard: "Diwakar",
		Number:     "4111111111111111",
		Cvc:        "456",
		ExpiryDate: "12/26",
	}
	// Create Credit Payment Method
	creditCardPayment := &payment.CreditCardPayment{CreditCard: creditCard}
	// Create Payment Context
	credit := payment.NewPaymentContext(creditCardPayment)
	// Process Payment by Credit Card
	creditPaymentRes, err := credit.Pay(399)
	if err != nil {
		log.Printf("CreditCard Payment Failed: %v\n", err)
	}
	cpResBytes, _ := json.MarshalIndent(creditPaymentRes, "", "   ")
	log.Printf("CreditCard Payment Response: %v\n", string(cpResBytes))

	// Add UPI Detail
	upiCard := payment.UpiCard{
		Vpa: "test@upi",
	}
	// Create UPI Payment Method
	upiPayment := &payment.UpiPayment{UpiCard: upiCard}

	// Create payment conetext
	upi := payment.NewPaymentContext(upiPayment)
	// Process Payment by UPI
	upiPaymentRes, err := upi.Pay(amount)

	if err != nil {
		log.Printf("UPI Payment Failed: %v\n", err)
	}
	upiResBytes, _ := json.MarshalIndent(upiPaymentRes, "", "   ")
	log.Printf("UPI Payment Response: %v\n", string(upiResBytes))
}
