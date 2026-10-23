package mercadopago

import "github.com/cristiantebaldi/factory-method-go/payment"

type Factory struct{}

func NewFactory() Factory {
	return Factory{}
}

func (Factory) CreateProcessor() (payment.Processor, error) {
	return NewProcessor("demo-mercadopago-token", 2)
}
