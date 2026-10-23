package builder

import (
	"fmt"

	"github.com/cristiantebaldi/builder-go/order"
)

type common struct {
	orderID       string
	items         []order.Item
	customerEmail string
	shipping      *order.Address
	discount      *order.Discount
	result        *order.Checkout
}

func (b *common) reset() {
	*b = common{}
}

func (b *common) setOrderID(orderID string) {
	b.orderID = orderID
	b.result = nil
}

func (b *common) addItem(name string, price float64, quantity int) {
	b.items = append(b.items, order.Item{Name: name, Price: price, Quantity: quantity})
	b.result = nil
}

func (b *common) setCustomerEmail(email string) {
	b.customerEmail = email
	b.result = nil
}

func (b *common) setShippingAddress(address order.Address) {
	b.shipping = &address
	b.result = nil
}

func (b *common) applyDiscount(amount float64, discountType order.DiscountType) {
	b.discount = &order.Discount{Amount: amount, Type: discountType}
	b.result = nil
}

func (b *common) build(method order.PaymentMethod, card *order.CardDetails, pix *order.PixDetails) (*order.Checkout, error) {
	if b.orderID == "" {
		return nil, fmt.Errorf("order ID is required")
	}
	if len(b.items) == 0 {
		return nil, fmt.Errorf("at least one item is required")
	}
	if b.customerEmail == "" {
		return nil, fmt.Errorf("customer email is required")
	}

	subtotal := 0.0
	for _, item := range b.items {
		if item.Name == "" {
			return nil, fmt.Errorf("item name is required")
		}
		if item.Price < 0 {
			return nil, fmt.Errorf("item price cannot be negative")
		}
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("item quantity must be greater than zero")
		}
		subtotal += item.Price * float64(item.Quantity)
	}

	discountAmount, err := calculateDiscount(subtotal, b.discount)
	if err != nil {
		return nil, err
	}

	result := &order.Checkout{
		OrderID:        b.orderID,
		Items:          append([]order.Item(nil), b.items...),
		CustomerEmail:  b.customerEmail,
		PaymentMethod:  method,
		Shipping:       copyAddress(b.shipping),
		Subtotal:       subtotal,
		DiscountAmount: discountAmount,
		TotalAmount:    subtotal - discountAmount,
	}
	if card != nil {
		details := *card
		result.CardDetails = &details
	}
	if pix != nil {
		details := *pix
		result.PixDetails = &details
	}

	b.result = result
	return result, nil
}

func (b *common) getResult() (*order.Checkout, error) {
	if b.result == nil {
		return nil, fmt.Errorf("Build must be called before GetResult")
	}
	return copyCheckout(b.result), nil
}

func calculateDiscount(subtotal float64, discount *order.Discount) (float64, error) {
	if discount == nil {
		return 0, nil
	}
	if discount.Amount < 0 {
		return 0, fmt.Errorf("discount cannot be negative")
	}

	var amount float64
	switch discount.Type {
	case order.Percentage:
		if discount.Amount > 100 {
			return 0, fmt.Errorf("percentage discount cannot exceed 100")
		}
		amount = subtotal * discount.Amount / 100
	case order.Fixed:
		amount = discount.Amount
	default:
		return 0, fmt.Errorf("unsupported discount type %q", discount.Type)
	}
	if amount > subtotal {
		return 0, fmt.Errorf("discount cannot exceed subtotal")
	}
	return amount, nil
}

func copyAddress(address *order.Address) *order.Address {
	if address == nil {
		return nil
	}
	copy := *address
	return &copy
}

func copyCheckout(checkout *order.Checkout) *order.Checkout {
	copy := *checkout
	copy.Items = append([]order.Item(nil), checkout.Items...)
	copy.Shipping = copyAddress(checkout.Shipping)
	if checkout.CardDetails != nil {
		details := *checkout.CardDetails
		copy.CardDetails = &details
	}
	if checkout.PixDetails != nil {
		details := *checkout.PixDetails
		copy.PixDetails = &details
	}
	return &copy
}
