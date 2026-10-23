package builder

import "github.com/cristiantebaldi/builder-go/order"

type CheckoutBuilder interface {
	Reset()
	SetOrderID(orderID string)
	AddItem(name string, price float64, quantity int)
	SetCustomerEmail(email string)
	SetShippingAddress(address order.Address)
	ApplyDiscount(amount float64, discountType order.DiscountType)
	Build() (*order.Checkout, error)
	GetResult() (*order.Checkout, error)
}

type CreditCardBuilder interface {
	CheckoutBuilder
	SetCardDetails(details order.CardDetails)
}

type PixBuilder interface {
	CheckoutBuilder
	SetPixDetails(details order.PixDetails)
}
