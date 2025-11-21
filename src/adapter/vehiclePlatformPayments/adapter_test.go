package vehicleplatformpayments

import (
	"context"
	"testing"

	mocks "github.com/caiiomp/vehicle-platform-sales/src/core/_mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePayment(t *testing.T) {
	ctx := context.TODO()
	amount := float64(50000)
	status := "APPROVED"
	paymentID := uuid.NewString()

	httpClientMocked := mocks.NewVehiclePlatformPaymentsHttpClient(t)

	httpClientMocked.On("GeneratePayment", ctx, amount, status).
		Return(paymentID, nil)

	adapter := NewVehiclePlatformPaymentsAdapter(httpClientMocked)

	expected := paymentID

	actual, err := adapter.GeneratePayment(ctx, amount, status)

	assert.Equal(t, expected, actual)
	assert.Nil(t, err)
}
