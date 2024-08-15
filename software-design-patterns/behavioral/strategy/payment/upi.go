package payment

type UpiPayment struct {
	UpiCard UpiCard
}

type UpiCard struct {
	Vpa string
}

func (u *UpiPayment) ProcessPayment(amt float32) (*PaymentResponse, error) {
	fee, err := u.CalculateFee(amt)
	if err != nil {
		return nil, err
	}
	// Implement your UPI Payment logic
	// ...

	payableAmt := amt + *fee
	// Return PaymentResponse
	paymentRes := &PaymentResponse{
		Status:  "success",
		Message: "UPI payment ",
		Balance: Balance{
			Amount: payableAmt,
			Fee:    *fee,
		},
	}

	return paymentRes, nil
}

func (u *UpiPayment) CalculateFee(amt float32) (*float32, error) {
	feePercent, err := u.GetFee()
	if err != nil {
		return nil, err
	}
	totalfee := amt * (*feePercent)

	return &totalfee, nil
}

func (u *UpiPayment) GetFee() (*float32, error) {
	fee := float32(0) // 3%
	return &fee, nil
}
