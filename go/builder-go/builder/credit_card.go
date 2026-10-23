package builder

import (
	"fmt"

	"github.com/cristiantebaldi/builder-go/order"
)

type CreditCardCheckoutBuilder struct {
	common
	card order.CardDetails
}

func NewCreditCardCheckoutBuilder() *CreditCardCheckoutBuilder {
	return &CreditCardCheckoutBuilder{}
}

func (b *CreditCardCheckoutBuilder) Reset() {
	b.common.reset()
	b.card = order.CardDetails{}
}

func (b *CreditCardCheckoutBuilder) SetOrderID(orderID string) {
	b.common.setOrderID(orderID)
}

func (b *CreditCardCheckoutBuilder) AddItem(name string, price float64, quantity int) {
	b.common.addItem(name, price, quantity)
}

func (b *CreditCardCheckoutBuilder) SetCustomerEmail(email string) {
	b.common.setCustomerEmail(email)
}

func (b *CreditCardCheckoutBuilder) SetShippingAddress(address order.Address) {
	b.common.setShippingAddress(address)
}

func (b *CreditCardCheckoutBuilder) SetCardDetails(details order.CardDetails) {
	b.card = details
	b.result = nil
}

func (b *CreditCardCheckoutBuilder) ApplyDiscount(amount float64, discountType order.DiscountType) {
	b.common.applyDiscount(amount, discountType)
}

func (b *CreditCardCheckoutBuilder) Build() (*order.Checkout, error) {
	if b.card.Number == "" || b.card.HolderName == "" || b.card.Expiration == "" || b.card.CVV == "" {
		return nil, fmt.Errorf("card details are incomplete")
	}
	return b.common.build(order.CreditCard, &b.card, nil)
}

func (b *CreditCardCheckoutBuilder) GetResult() (*order.Checkout, error) {
	return b.common.getResult()
}
