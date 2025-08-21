package domain

import "testing"

func TestDelivery_RequiredField(t *testing.T) {
	run(t, deliveryRequiredValidateTestCases)
}

func TestDelivery_PhoneField(t *testing.T) {
	run(t, deliveryPhoneValidateTestCases)
}

func TestDelivery_EmailField(t *testing.T) {
	run(t, deliveryEmailValidateTestCases)
}

func run(t *testing.T, tests []DeliveryValidateTestCase) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.delivery.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Delivery.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if actual != tt.expected {
				t.Errorf("Delivery.Validate() = %v, want %v", actual, tt.expected)
			}

			if tt.wantErr && tt.errMsg != "" {
				if err == nil {
					t.Error("Expected error but got nil")
					return
				}
				if !contains(err.Error(), tt.errMsg) {
					t.Errorf("Error message '%s' should contain '%s'", err.Error(), tt.errMsg)
				}
			}
		})
	}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
