package director

import (
	"github.com/cristiantebaldi/builder-go/builder"
	"github.com/cristiantebaldi/builder-go/order"
)

type CheckoutDirector struct{}

func NewCheckoutDirector() *CheckoutDirector {
	return &CheckoutDirector{}
}

func (d *CheckoutDirector) CreateBasicCreditCardCheckout(
	b builder.CreditCardBuilder,
	orderID, customerEmail, productName string,
	productPrice float64,
	card order.CardDetails,
) (*order.Checkout, error) {
	b.Reset()
	b.SetOrderID(orderID)
	b.SetCustomerEmail(customerEmail)
	b.AddItem(productName, productPrice, 1)
	b.SetCardDetails(card)
	return b.Build()
}

func (d *CheckoutDirector) CreateBasicPixCheckout(
	b builder.PixBuilder,
	orderID, customerEmail, productName string,
	productPrice float64,
	pix order.PixDetails,
) (*order.Checkout, error) {
	b.Reset()
	b.SetOrderID(orderID)
	b.SetCustomerEmail(customerEmail)
	b.AddItem(productName, productPrice, 1)
	b.SetPixDetails(pix)
	return b.Build()
}

func (d *CheckoutDirector) CreateSmartphoneBundle(
	b builder.CreditCardBuilder,
	orderID, customerEmail string,
	address order.Address,
	card order.CardDetails,
) (*order.Checkout, error) {
	b.Reset()
	b.SetOrderID(orderID)
	b.SetCustomerEmail(customerEmail)
	b.SetShippingAddress(address)
	b.SetCardDetails(card)
	b.AddItem("Smartphone", 1000, 1)
	b.AddItem("Case", 100, 1)
	b.AddItem("Screen Protector", 10, 1)
	b.AddItem("Charging Cable", 20, 1)
	b.AddItem("Headphones", 50, 1)
	b.AddItem("Earbuds", 30, 1)
	b.AddItem("Power Bank", 100, 1)
	b.AddItem("Wireless Charger", 100, 1)
	b.AddItem("Bluetooth Speaker", 100, 1)
	b.ApplyDiscount(15, order.Percentage)
	return b.Build()
}

func (d *CheckoutDirector) CreateGamerPackage(
	b builder.CreditCardBuilder,
	orderID, customerEmail string,
	address order.Address,
	card order.CardDetails,
) (*order.Checkout, error) {
	b.Reset()
	b.SetOrderID(orderID)
	b.SetCustomerEmail(customerEmail)
	b.SetShippingAddress(address)
	b.SetCardDetails(card)
	b.AddItem("Gaming Mouse", 100, 1)
	b.AddItem("Gaming Keyboard", 100, 1)
	b.AddItem("Gaming Headset", 100, 1)
	b.AddItem("Gaming Mousepad", 100, 1)
	b.AddItem("Gaming Chair", 100, 1)
	b.AddItem("Gaming Table", 100, 1)
	b.ApplyDiscount(15, order.Percentage)
	return b.Build()
}
