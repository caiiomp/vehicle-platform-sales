package saleApi

type saleQuery struct {
	Status string `form:"status"`
}

type saleWebhookRequest struct {
	PaymentID string `json:"payment_id"`
	Status    string `json:"status"`
}
