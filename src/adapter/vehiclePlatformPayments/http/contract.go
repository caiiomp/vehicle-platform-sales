package http

type createPaymentRequest struct {
	WebhookUrl string  `json:"webhook_url"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
}

type createPaymentResponse struct {
	PaymentID  string  `json:"payment_id"`
	WebhookUrl string  `json:"webhook_url"`
	Status     string  `json:"status"`
	Amount     float64 `json:"amount"`
}
