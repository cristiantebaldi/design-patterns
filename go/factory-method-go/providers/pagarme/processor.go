package pagarme

import (
	"fmt"
)

type Processor struct {
	apiKey               string
	installments         int
	captureAutomatically bool
}

func NewProcessor(apiKey string, installments int, captureAutomatically bool) (*Processor, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("PagarMe: API Key é obrigatória")
	}
	if installments < 1 || installments > 12 {
		return nil, fmt.Errorf("PagarMe: número de parcelas deve estar entre 1 e 12")
	}

	return &Processor{
		apiKey:               apiKey,
		installments:         installments,
		captureAutomatically: captureAutomatically,
	}, nil
}

func (p *Processor) ProcessPayment(amount float64) error {
	fmt.Printf("\nProcessando pagamento de R$ %.2f via Cartão de Crédito (PagarMe)\n", amount)
	fmt.Printf("Parcelas: %d\n", p.installments)
	fmt.Printf("Captura automática: %t\n", p.captureAutomatically)
	fmt.Println("Status: Pagamento aprovado")
	return nil
}

func (p *Processor) PaymentInfo() string {
	return "Cartão de Crédito (PagarMe) - Taxa: 2.8%"
}
