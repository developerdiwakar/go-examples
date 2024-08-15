package payment

type CreditCardPayment struct {
	CreditCard CreditCard
}

type CreditCard struct {
	NameOnCard string
	Number     string
	Cvc        string
	ExpiryDate string
}

func (c *CreditCardPayment) ProcessPayment(amount float64) (*PaymentResponse, error) {
	fee, err := c.GetFee()
	if err != nil {
		return nil, err
	}
	paymentRes := &PaymentResponse{
		Status:  "success",
		Message: "Payment Successful by Credit Card",
		Balance: Balance{
			Amount: amount, // Total Amount
			Fee:    fee,
		},
	}
	return paymentRes, nil
}

func (c *CreditCardPayment) CalculateFee(amount float64) (float64, error) {
	fee, err := c.GetFee()
	if err != nil {
		return 0, err
	}
	totalFee := amount * fee
	return totalFee, nil

}

func (c *CreditCardPayment) GetFee() (float64, error) {
	// Get Fee logic from DB or other sources
	return 0.03, nil // 3% fee
}
