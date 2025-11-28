package vehicleplatformpayments

import (
	"context"

	"github.com/caiiomp/vehicle-platform-sales/src/adapter/vehiclePlatformPayments/http"
	interfaces "github.com/caiiomp/vehicle-platform-sales/src/core/_interfaces"
)

type vehiclePlatformPaymentsAdapter struct {
	httpClient http.VehiclePlatformPaymentsHttpClient
}

func NewVehiclePlatformPaymentsAdapter(httpClient http.VehiclePlatformPaymentsHttpClient) interfaces.VehiclePlatformPaymentsAdapter {
	return &vehiclePlatformPaymentsAdapter{
		httpClient: httpClient,
	}
}

func (ref *vehiclePlatformPaymentsAdapter) GeneratePayment(ctx context.Context, amount float64, status string) (string, error) {
	return ref.httpClient.GeneratePayment(ctx, amount, status)
}
