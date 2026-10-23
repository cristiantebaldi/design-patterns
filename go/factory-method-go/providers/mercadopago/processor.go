package mercadopago

import "fmt"

type Processor struct {
	accessToken  string
	installments int
}

func NewProcessor(accessToken string, installments int) (*Processor, error) {
	if accessToken == "" {
		return nil, fmt.Errorf("MercadoPago: Access Token é obrigatório")
	}
	if installments < 1 || installments > 12 {
		return nil, fmt.Errorf("MercadoPago: número de parcelas deve estar entre 1 e 12")
	}

	return &Processor{accessToken: accessToken, installments: installments}, nil
}

func (p *Processor) ProcessPayment(amount float64) error {
	fmt.Printf("\nProcessando pagamento de R$ %.2f via Cartão de Crédito (MercadoPago)\n", amount)
	fmt.Printf("Parcelas: %d\n", p.installments)
	fmt.Println("Status: Pagamento aprovado")
	return nil
}

func (p *Processor) PaymentInfo() string {
	return "Cartão de Crédito (MercadoPago) - Taxa: 2.3%"
}
