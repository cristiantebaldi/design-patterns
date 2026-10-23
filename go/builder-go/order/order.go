package order

import "fmt"

type PaymentMethod string

const (
	CreditCard PaymentMethod = "credit_card"
	Pix        PaymentMethod = "pix"
)

type DiscountType string

const (
	Percentage DiscountType = "percentage"
	Fixed      DiscountType = "fixed"
)

type Address struct {
	Address string
	City    string
}

type Item struct {
	Name     string
	Price    float64
	Quantity int
}

type CardDetails struct {
	Number     string
	HolderName string
	Expiration string
	CVV        string
}

type PixDetails struct {
	Key       string
	PayerName string
}

type Discount struct {
	Amount float64
	Type   DiscountType
}

type Checkout struct {
	OrderID        string
	Items          []Item
	CustomerEmail  string
	PaymentMethod  PaymentMethod
	CardDetails    *CardDetails
	PixDetails     *PixDetails
	Shipping       *Address
	Subtotal       float64
	DiscountAmount float64
	TotalAmount    float64
}

func (c Checkout) Summary() string {
	return fmt.Sprintf("Pedido #%s - %s - $%.2f", c.OrderID, c.PaymentMethod, c.TotalAmount)
}
