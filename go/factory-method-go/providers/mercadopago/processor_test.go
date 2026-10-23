package mercadopago

import "testing"

func TestNewProcessorValidatesConfiguration(t *testing.T) {
	tests := []struct {
		name         string
		accessToken  string
		installments int
		wantError    bool
	}{
		{name: "missing access token", installments: 1, wantError: true},
		{name: "installments below range", accessToken: "demo-token", installments: 0, wantError: true},
		{name: "installments above range", accessToken: "demo-token", installments: 13, wantError: true},
		{name: "valid configuration", accessToken: "demo-token", installments: 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewProcessor(test.accessToken, test.installments)
			if (err != nil) != test.wantError {
				t.Fatalf("NewProcessor() error = %v, wantError %t", err, test.wantError)
			}
		})
	}
}
