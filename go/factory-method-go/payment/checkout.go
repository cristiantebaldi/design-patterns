package payment

import (
	"fmt"
	"io"
	"os"
)

type Checkout struct {
	output io.Writer
}

func NewCheckout() *Checkout {
	return &Checkout{output: os.Stdout}
}

func (c *Checkout) ProcessPayment(amount float64, processor Processor) error {
	if processor == nil {
		return fmt.Errorf("payment processor is nil")
	}
	return processPayment(c.output, amount, processor)
}

type FactoryCheckout struct {
	factory ProcessorFactory
	output  io.Writer
}

func NewFactoryCheckout(factory ProcessorFactory) *FactoryCheckout {
	return &FactoryCheckout{factory: factory, output: os.Stdout}
}

func (c *FactoryCheckout) ProcessPayment(amount float64) error {
	if c.factory == nil {
		return fmt.Errorf("payment processor factory is nil")
	}
	processor, err := c.factory.CreateProcessor()
	if err != nil {
		return fmt.Errorf("create payment processor: %w", err)
	}
	if processor == nil {
		return fmt.Errorf("payment processor is nil")
	}
	return processPayment(c.output, amount, processor)
}

func processPayment(output io.Writer, amount float64, processor Processor) error {
	fmt.Fprintln(output, "\nInformações do pagamento:")
	fmt.Fprintln(output, processor.PaymentInfo())
	return processor.ProcessPayment(amount)
}
