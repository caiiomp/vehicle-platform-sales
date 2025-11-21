package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type VehiclePlatformPaymentsHttpClient interface {
	GeneratePayment(ctx context.Context, amount float64, status string) (string, error)
}

type vehiclePlatformPaymentsHttpClient struct {
	client                      *http.Client
	vehiclePlatformPaymentsHost string
	vehiclePlatformSalesHost    string
}

func NewVehiclePlatformSalesHttpClient(client *http.Client, vehiclePlatformPaymentsHost, vehiclePlatformSalesHost string) VehiclePlatformPaymentsHttpClient {
	return &vehiclePlatformPaymentsHttpClient{
		client:                      client,
		vehiclePlatformPaymentsHost: vehiclePlatformPaymentsHost,
		vehiclePlatformSalesHost:    vehiclePlatformSalesHost,
	}
}

func (ref *vehiclePlatformPaymentsHttpClient) GeneratePayment(ctx context.Context, amount float64, status string) (string, error) {
	url := ref.vehiclePlatformPaymentsHost + "/payments"
	webhookUrl := ref.vehiclePlatformSalesHost + "/sales/webhook"

	payment := createPaymentRequest{
		WebhookUrl: webhookUrl,
		Amount:     amount,
		Status:     status,
	}

	data, err := json.Marshal(payment)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := ref.client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	rawResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to create payment at vehicle platform payments: %v", string(rawResponse))
	}

	var createPaymentResponse createPaymentResponse
	if err = json.Unmarshal(rawResponse, &createPaymentResponse); err != nil {
		return "", err
	}

	return createPaymentResponse.PaymentID, nil
}
