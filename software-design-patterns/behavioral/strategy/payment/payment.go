package payment

type PaymentMethod interface {
	ProcessPayment(amount float64) (*PaymentResponse, error)
	CalculateFee(amount float64) (float64, error)
	GetFee() (float64, error)
}

type PaymentResponse struct {
	Status  string
	Message string
	Balance Balance
}

type Balance struct {
	Amount float64
	Fee    float64
}

type PaymentContext struct {
	PayMethod PaymentMethod
}

func NewPaymentContext(payMethod PaymentMethod) *PaymentContext {
	return &PaymentContext{PayMethod: payMethod}
}

func (pc *PaymentContext) Pay(amount float64) (*PaymentResponse, error) {
	return pc.PayMethod.ProcessPayment(amount)
}
