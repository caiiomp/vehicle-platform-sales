package interfaces

import "context"

type VehiclePlatformPaymentsAdapter interface {
	GeneratePayment(ctx context.Context, amount float64, status string) (string, error)
}
