package pix

import "testing"

func TestNewProcessorValidatesConfiguration(t *testing.T) {
	tests := []struct {
		name         string
		pixKey       string
		merchantName string
		wantError    bool
	}{
		{name: "missing PIX key", merchantName: "Loja", wantError: true},
		{name: "missing merchant name", pixKey: "email@exemplo.com", wantError: true},
		{name: "valid configuration", pixKey: "email@exemplo.com", merchantName: "Loja"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewProcessor(test.pixKey, test.merchantName)
			if (err != nil) != test.wantError {
				t.Fatalf("NewProcessor() error = %v, wantError %t", err, test.wantError)
			}
		})
	}
}
