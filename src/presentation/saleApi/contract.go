package saleApi

type saleWebhookRequest struct {
	PaymentID string `json:"payment_id"`
	Status    string `json:"status"`
}
