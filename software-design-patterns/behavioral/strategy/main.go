package main

import (
	"encoding/json"
	"go-exercises/software-design-patterns/behavioral/strategy/payment"
	"log"
)

func main() {
	// Create Payment Methods
	creditCard := payment.CreditCard{
		NameOnCard: "Diwakar",
		Number:     "4111111111111111",
		Cvc:        "456",
		ExpiryDate: "12/26",
	}

	// Create Payment Contexts
	creditCardPayment := payment.NewPaymentContext(&payment.CreditCardPayment{CreditCard: creditCard})
	creditPaymentRes, err := creditCardPayment.Pay(399)
	if err != nil {
		log.Printf("CreditCard Payment Failed: %v\n", err)
	}
	cpBytes, _ := json.MarshalIndent(creditPaymentRes, "", "   ")
	log.Printf("CreditCard Payment: %v\n", string(cpBytes))
}
