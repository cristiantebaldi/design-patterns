package payment

type Processor interface {
	ProcessPayment(amount float64) error
	PaymentInfo() string
}

type ProcessorFactory interface {
	CreateProcessor() (Processor, error)
}
