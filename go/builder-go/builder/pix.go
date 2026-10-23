package builder

import (
	"fmt"

	"github.com/cristiantebaldi/builder-go/order"
)

type PixCheckoutBuilder struct {
	common
	pix order.PixDetails
}

func NewPixCheckoutBuilder() *PixCheckoutBuilder {
	return &PixCheckoutBuilder{}
}

func (b *PixCheckoutBuilder) Reset() {
	b.common.reset()
	b.pix = order.PixDetails{}
}

func (b *PixCheckoutBuilder) SetOrderID(orderID string) {
	b.common.setOrderID(orderID)
}

func (b *PixCheckoutBuilder) AddItem(name string, price float64, quantity int) {
	b.common.addItem(name, price, quantity)
}

func (b *PixCheckoutBuilder) SetCustomerEmail(email string) {
	b.common.setCustomerEmail(email)
}

func (b *PixCheckoutBuilder) SetShippingAddress(address order.Address) {
	b.common.setShippingAddress(address)
}

func (b *PixCheckoutBuilder) SetPixDetails(details order.PixDetails) {
	b.pix = details
	b.result = nil
}

func (b *PixCheckoutBuilder) ApplyDiscount(amount float64, discountType order.DiscountType) {
	b.common.applyDiscount(amount, discountType)
}

func (b *PixCheckoutBuilder) Build() (*order.Checkout, error) {
	if b.pix.Key == "" || b.pix.PayerName == "" {
		return nil, fmt.Errorf("PIX details are incomplete")
	}
	return b.common.build(order.Pix, nil, &b.pix)
}

func (b *PixCheckoutBuilder) GetResult() (*order.Checkout, error) {
	return b.common.getResult()
}
