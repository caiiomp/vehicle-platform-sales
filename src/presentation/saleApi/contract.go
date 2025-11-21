package saleApi

type paymentURI struct {
	PaymentID string `uri:"payment_id"`
}

type updateSaleRequest struct {
	Status string `json:"status"`
}

type saleWebhookRequest struct {
	PaymentID string `json:"payment_id"`
	Status    string `json:"status"`
}
