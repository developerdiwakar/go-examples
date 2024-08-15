package payment

// CreditCardPayment is a Payment Method
type CreditCardPayment struct {
	CreditCard CreditCard
}

// CreditCard contains credit card details
type CreditCard struct {
	NameOnCard string
	Number     string
	Cvc        string
	ExpiryDate string
}

// ProcessPayment method will process the Credit Card Payment and return the PaymentResponse.
func (c *CreditCardPayment) ProcessPayment(amount float32) (*PaymentResponse, error) {
	fee, err := c.CalculateFee(amount)
	if err != nil {
		return nil, err
	}
	// Calculate total payable amount
	payableAmt := amount + *fee
	// Prepare payment response
	paymentRes := &PaymentResponse{
		Status:  "success",
		Message: "Payment Successful by Credit Card",
		Balance: Balance{
			Amount: payableAmt, // Total Amount
			Fee:    *fee,
		},
	}
	return paymentRes, nil
}

func (c *CreditCardPayment) CalculateFee(amount float32) (*float32, error) {
	feePercent, err := c.GetFee()
	if err != nil {
		return nil, err
	}
	totalFee := amount * (*feePercent)
	return &totalFee, nil

}

func (c *CreditCardPayment) GetFee() (*float32, error) {
	// Implement Fee logic from DB or other sources
	feePercent := float32(0.03)
	return &feePercent, nil // 3% fee
}
