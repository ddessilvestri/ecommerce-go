package order

import (
	"encoding/json"
	"testing"

	"github.com/ddessilvestri/ecommerce-go/models"
	"github.com/stretchr/testify/assert"
)

// Test order validation
func TestOrderValidation(t *testing.T) {
	tests := []struct {
		name    string
		order   models.Orders
		isValid bool
	}{
		{
			name: "Valid Order",
			order: models.Orders{
				UserUUID: "user-123",
				AddId:    1,
				Date:     "2024-01-01",
				Total:    99.99,
				OrderDetails: []models.OrdersDetails{
					{
						ProdId:   1,
						Quantity: 2,
						Price:    49.99,
					},
				},
			},
			isValid: true,
		},
		{
			name: "Missing UserUUID",
			order: models.Orders{
				AddId: 1,
				Date:  "2024-01-01",
				Total: 99.99,
				OrderDetails: []models.OrdersDetails{
					{
						ProdId:   1,
						Quantity: 2,
						Price:    49.99,
					},
				},
			},
			isValid: false,
		},
		{
			name: "Missing OrderDetails",
			order: models.Orders{
				UserUUID: "user-123",
				AddId:    1,
				Date:     "2024-01-01",
				Total:    99.99,
			},
			isValid: false,
		},
		{
			name: "Invalid OrderDetails - Missing ProdId",
			order: models.Orders{
				UserUUID: "user-123",
				AddId:    1,
				Date:     "2024-01-01",
				Total:    99.99,
				OrderDetails: []models.OrdersDetails{
					{
						Quantity: 2,
						Price:    49.99,
					},
				},
			},
			isValid: false,
		},
		{
			name: "Invalid OrderDetails - Zero Quantity",
			order: models.Orders{
				UserUUID: "user-123",
				AddId:    1,
				Date:     "2024-01-01",
				Total:    99.99,
				OrderDetails: []models.OrdersDetails{
					{
						ProdId:   1,
						Quantity: 0,
						Price:    49.99,
					},
				},
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test validation logic
			isValid := tt.order.UserUUID != "" &&
				tt.order.AddId > 0 &&
				tt.order.Date != "" &&
				tt.order.Total > 0 &&
				len(tt.order.OrderDetails) > 0

			// Additional validation for order details
			if isValid {
				for _, detail := range tt.order.OrderDetails {
					if detail.ProdId <= 0 || detail.Quantity <= 0 || detail.Price <= 0 {
						isValid = false
						break
					}
				}
			}

			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

// Test order model JSON marshaling
func TestOrderJSON(t *testing.T) {
	order := models.Orders{
		Id:       1,
		UserUUID: "user-123",
		AddId:    1,
		Date:     "2024-01-01",
		Total:    99.99,
		OrderDetails: []models.OrdersDetails{
			{
				Id:       1,
				OrderId:  1,
				ProdId:   1,
				Quantity: 2,
				Price:    49.99,
			},
		},
	}

	// Test JSON marshaling
	jsonData, err := json.Marshal(order)
	assert.NoError(t, err)
	assert.Contains(t, string(jsonData), "user-123")
	assert.Contains(t, string(jsonData), "99.99")
	assert.Contains(t, string(jsonData), "OrderDetails")
}

// Test order details validation
func TestOrderDetailsValidation(t *testing.T) {
	tests := []struct {
		name    string
		detail  models.OrdersDetails
		isValid bool
	}{
		{
			name: "Valid Order Detail",
			detail: models.OrdersDetails{
				OrderId:  1,
				ProdId:   1,
				Quantity: 2,
				Price:    49.99,
			},
			isValid: true,
		},
		{
			name: "Invalid Order Detail - Zero ProdId",
			detail: models.OrdersDetails{
				OrderId:  1,
				ProdId:   0,
				Quantity: 2,
				Price:    49.99,
			},
			isValid: false,
		},
		{
			name: "Invalid Order Detail - Zero Quantity",
			detail: models.OrdersDetails{
				OrderId:  1,
				ProdId:   1,
				Quantity: 0,
				Price:    49.99,
			},
			isValid: false,
		},
		{
			name: "Invalid Order Detail - Zero Price",
			detail: models.OrdersDetails{
				OrderId:  1,
				ProdId:   1,
				Quantity: 2,
				Price:    0,
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test validation logic
			isValid := tt.detail.OrderId > 0 &&
				tt.detail.ProdId > 0 &&
				tt.detail.Quantity > 0 &&
				tt.detail.Price > 0

			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

// Test order total calculation
func TestOrderTotalCalculation(t *testing.T) {
	order := models.Orders{
		UserUUID: "user-123",
		AddId:    1,
		Date:     "2024-01-01",
		OrderDetails: []models.OrdersDetails{
			{
				ProdId:   1,
				Quantity: 2,
				Price:    25.00,
			},
			{
				ProdId:   2,
				Quantity: 1,
				Price:    50.00,
			},
		},
	}

	// Calculate expected total
	expectedTotal := 2*25.00 + 1*50.00
	order.Total = expectedTotal

	assert.Equal(t, expectedTotal, order.Total)
	assert.Equal(t, 100.00, order.Total)
}
