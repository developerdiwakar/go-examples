package payment

type PaymentMethod interface {
	ProcessPayment(amount float32) (*PaymentResponse, error)
	CalculateFee(amount float32) (*float32, error)
	GetFee() (*float32, error)
}

type PaymentResponse struct {
	Status  string
	Message string
	Balance Balance
}

type Balance struct {
	Amount float32
	Fee    float32
}

type PaymentContext struct {
	PaymentMethod PaymentMethod
}

func NewPaymentContext(method PaymentMethod) *PaymentContext {
	return &PaymentContext{PaymentMethod: method}
}

func (pc *PaymentContext) Pay(amount float32) (*PaymentResponse, error) {
	return pc.PaymentMethod.ProcessPayment(amount)
}
