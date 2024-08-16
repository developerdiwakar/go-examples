package payment

// PaymentStrategy defines the interface for payment strategies with fees and responses
type PaymentStrategy interface {
	Pay(amount float64) PaymentResponse
	CalculateFee(amount float64) float64
}

// PaymentResponse represents a generic response after a payment is processed
type PaymentResponse struct {
	TxnId    string  `json:"txnId,omitempty"`
	Status   string  `json:"status"`
	Message  string  `json:"message"`
	TotalAmt float64 `json:"totalAmt,omitempty"`
	Fee      float64 `json:"fee,omitempty"`
}

// PaymenOrderResponse to capture actual response coming from payment gateway
type PaymentOrderResponse struct {
	TxnId   string `json:"txnId,omitempty"`
	Status  string `json:"status"`
	Message string `json:"message"`
	// ...
}

// PaymentProcessor uses a strategy to process payments
type PaymentProcessor struct {
	PaymentStrategy PaymentStrategy
}

// SetStrategy allows setting the payment strategy dynamically
func (pp *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	pp.PaymentStrategy = strategy
}

// Checkout handles the payment process with fee calculation and response handling
func (pp *PaymentProcessor) Checkout(amount float64) PaymentResponse {
	fee := pp.PaymentStrategy.CalculateFee(amount)
	total := amount + fee
	response := pp.PaymentStrategy.Pay(total)
	// Add Additional response field
	if response.Status == "success" {
		response.Fee = fee
		response.TotalAmt = total
	}
	return response
}
