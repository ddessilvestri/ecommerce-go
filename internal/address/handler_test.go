package address

import (
	"encoding/json"
	"testing"

	"github.com/ddessilvestri/ecommerce-go/models"
	"github.com/stretchr/testify/assert"
)

// Simple test for address validation
func TestAddressValidation(t *testing.T) {
	tests := []struct {
		name    string
		address models.Address
		isValid bool
	}{
		{
			name: "Valid Address",
			address: models.Address{
				Title:      "Home",
				Name:       "John Doe",
				Address:    "123 Main St",
				City:       "New York",
				State:      "NY",
				PostalCode: "10001",
				Phone:      "+1-555-123-4567",
			},
			isValid: true,
		},
		{
			name: "Missing Title",
			address: models.Address{
				Name:       "John Doe",
				Address:    "123 Main St",
				City:       "New York",
				State:      "NY",
				PostalCode: "10001",
				Phone:      "+1-555-123-4567",
			},
			isValid: false,
		},
		{
			name: "Missing Name",
			address: models.Address{
				Title:      "Home",
				Address:    "123 Main St",
				City:       "New York",
				State:      "NY",
				PostalCode: "10001",
				Phone:      "+1-555-123-4567",
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test validation logic
			isValid := tt.address.Title != "" &&
				tt.address.Name != "" &&
				tt.address.Address != "" &&
				tt.address.City != "" &&
				tt.address.Phone != "" &&
				tt.address.PostalCode != ""

			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

// Test address model JSON marshaling
func TestAddressJSON(t *testing.T) {
	address := models.Address{
		Id:         1,
		Title:      "Home",
		Name:       "John Doe",
		Address:    "123 Main St",
		City:       "New York",
		State:      "NY",
		PostalCode: "10001",
		Phone:      "+1-555-123-4567",
	}

	// Test JSON marshaling
	jsonData, err := json.Marshal(address)
	assert.NoError(t, err)
	assert.Contains(t, string(jsonData), "Home")
	assert.Contains(t, string(jsonData), "John Doe")
}
