package pagarme

import "testing"

func TestNewProcessorValidatesConfiguration(t *testing.T) {
	tests := []struct {
		name                 string
		apiKey               string
		installments         int
		captureAutomatically bool
		wantError            bool
	}{
		{name: "missing API key", installments: 1, wantError: true},
		{name: "installments below range", apiKey: "demo-key", installments: 0, wantError: true},
		{name: "installments above range", apiKey: "demo-key", installments: 13, wantError: true},
		{name: "valid configuration", apiKey: "demo-key", installments: 3, captureAutomatically: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewProcessor(test.apiKey, test.installments, test.captureAutomatically)
			if (err != nil) != test.wantError {
				t.Fatalf("NewProcessor() error = %v, wantError %t", err, test.wantError)
			}
		})
	}
}
