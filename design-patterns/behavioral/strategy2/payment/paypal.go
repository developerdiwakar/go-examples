package payment

import (
	"fmt"
	"go-exercises/software-design-patterns/behavioral/strategy2/helpers"
)

type PaypalStrategy struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *PaypalStrategy) CalculateFee(amt float64) float64 {
	fee := 0.05
	return amt * fee
}

func (p *PaypalStrategy) Pay(amt float64) PaymentResponse {
	orderRes := doPaypalPayment(amt, true)

	paymentRes := PaymentResponse{
		TxnId:   orderRes.TxnId,
		Message: orderRes.Message,
		Status:  orderRes.Status,
	}
	return paymentRes
}

func doPaypalPayment(amt float64, success bool) PaymentOrderResponse {
	var orderRes PaymentOrderResponse
	// integrate your payment gateway here to get transaction ID
	// ...
	orderRes.Status = "failed"
	orderRes.Message = "Payment Failed"
	if success {
		txnId := helpers.GenerateUniqueString()
		orderRes.Status = "success"
		orderRes.TxnId = txnId
		orderRes.Message = fmt.Sprintf("Payment of %2f processed with Paypal.", amt)
	}
	return orderRes
}
