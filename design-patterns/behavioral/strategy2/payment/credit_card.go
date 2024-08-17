package payment

import (
	"fmt"
	"go-exercises/design-patterns/behavioral/strategy2/helpers"
)

// CreditCardStrategy implements the PaymentStrategy for credit cards
type CreditCardStrategy struct {
	NameOnCard string
	Number     string
	ExpiryDate string
}

// Pay implements the final Payment functionality
func (c *CreditCardStrategy) Pay(amount float64) PaymentResponse {
	orderRes := doCreditPayment(amount, true)

	paymentRes := PaymentResponse{
		Status:  orderRes.Status,
		Message: orderRes.Message,
		TxnId:   orderRes.TxnId,
	}

	return paymentRes
}

// CalculateFee implements Fee calculation logic
func (c *CreditCardStrategy) CalculateFee(amount float64) float64 {
	fee := 0.03
	return amount * fee
}

// doPayment implement the actual payment logic and return transaction ID
func doCreditPayment(amt float64, success bool) PaymentOrderResponse {
	var orderRes PaymentOrderResponse
	// integrate your payment gateway here to get transaction ID
	// ...
	orderRes.Status = "failed"
	orderRes.Message = "Payment Failed"
	if success {
		txnId := helpers.GenerateUniqueString()
		orderRes.Status = "success"
		orderRes.TxnId = txnId
		orderRes.Message = fmt.Sprintf("Payment of %2f processed with Credit Card.", amt)
	}
	return orderRes
}
