package pix

import "fmt"

type Processor struct {
	pixKey       string
	merchantName string
}

func NewProcessor(pixKey string, merchantName string) (*Processor, error) {
	if pixKey == "" {
		return nil, fmt.Errorf("PIX: chave é obrigatória")
	}
	if merchantName == "" {
		return nil, fmt.Errorf("PIX: nome do comerciante é obrigatório")
	}

	return &Processor{pixKey: pixKey, merchantName: merchantName}, nil
}

func (p *Processor) ProcessPayment(amount float64) error {
	fmt.Printf("\nProcessando pagamento de R$ %.2f via PIX\n", amount)
	fmt.Printf("Recebedor: %s\n", p.merchantName)
	fmt.Println("Gerando QR Code...")
	fmt.Println("Status: QR Code gerado com sucesso")
	return nil
}

func (p *Processor) PaymentInfo() string {
	return "PIX - Taxa: 0% - Processamento instantâneo"
}
