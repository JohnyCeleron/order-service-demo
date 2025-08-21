package domain

import "testing"

func TestPayment_Validate(t *testing.T) {
	tests := []struct {
		name     string
		payment  Payment
		expected bool
		wantErr  bool
		errMsg   string
	}{
		{
			name: "valid payment",
			payment: Payment{
				Transaction:  "txn_123",
				RequestID:    "req_123",
				Currency:     "USD",
				Provider:     "stripe",
				Amount:       1500,
				PaymentDt:    1633024800,
				Bank:         "chase",
				DeliveryCost: 500,
				GoodsTotal:   900,
				CustomFee:    100,
			},
			expected: true,
			wantErr:  false,
		},
		{
			name: "missing required field - transaction",
			payment: Payment{
				Transaction:  "",
				Currency:     "USD",
				Provider:     "stripe",
				Amount:       1500,
				PaymentDt:    1633024800,
				Bank:         "chase",
				DeliveryCost: 500,
				GoodsTotal:   900,
				CustomFee:    100,
			},
			expected: false,
			wantErr:  true,
			errMsg:   "required",
		},
		{
			name: "amount zero",
			payment: Payment{
				Transaction:  "txn_123",
				Currency:     "USD",
				Provider:     "stripe",
				Amount:       0,
				PaymentDt:    1633024800,
				Bank:         "chase",
				DeliveryCost: 500,
				GoodsTotal:   900,
				CustomFee:    100,
			},
			expected: false,
			wantErr:  true,
			errMsg:   "gt",
		},
		{
			name: "negative delivery cost",
			payment: Payment{
				Transaction:  "txn_123",
				Currency:     "USD",
				Provider:     "stripe",
				Amount:       1500,
				PaymentDt:    1633024800,
				Bank:         "chase",
				DeliveryCost: -100,
				GoodsTotal:   900,
				CustomFee:    100,
			},
			expected: false,
			wantErr:  true,
			errMsg:   "gte",
		},
		{
			name: "goods total zero",
			payment: Payment{
				Transaction:  "txn_123",
				Currency:     "USD",
				Provider:     "stripe",
				Amount:       1500,
				PaymentDt:    1633024800,
				Bank:         "chase",
				DeliveryCost: 500,
				GoodsTotal:   0,
				CustomFee:    100,
			},
			expected: false,
			wantErr:  true,
			errMsg:   "gt",
		},
		{
			name: "negative custom fee",
			payment: Payment{
				Transaction:  "txn_123",
				Currency:     "USD",
				Provider:     "stripe",
				Amount:       1500,
				PaymentDt:    1633024800,
				Bank:         "chase",
				DeliveryCost: 500,
				GoodsTotal:   900,
				CustomFee:    -50,
			},
			expected: false,
			wantErr:  true,
			errMsg:   "gte",
		},
		{
			name: "amount consistency error - sum mismatch",
			payment: Payment{
				Transaction:  "txn_123",
				Currency:     "USD",
				Provider:     "stripe",
				Amount:       2000,
				PaymentDt:    1633024800,
				Bank:         "chase",
				DeliveryCost: 500,
				GoodsTotal:   900,
				CustomFee:    100,
			},
			expected: false,
			wantErr:  true,
			errMsg:   "amount consistency error",
		},
		{
			name: "valid with zero custom fee",
			payment: Payment{
				Transaction:  "txn_123",
				Currency:     "USD",
				Provider:     "stripe",
				Amount:       1400,
				PaymentDt:    1633024800,
				Bank:         "chase",
				DeliveryCost: 500,
				GoodsTotal:   900,
				CustomFee:    0,
			},
			expected: true,
			wantErr:  false,
		},
		{
			name: "valid with zero delivery cost",
			payment: Payment{
				Transaction:  "txn_123",
				Currency:     "USD",
				Provider:     "stripe",
				Amount:       1000,
				PaymentDt:    1633024800,
				Bank:         "chase",
				DeliveryCost: 0,
				GoodsTotal:   900,
				CustomFee:    100,
			},
			expected: true,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tt.payment.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Payment.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if actual != tt.expected {
				t.Errorf("Payment.Validate() = %v, want %v", actual, tt.expected)
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
