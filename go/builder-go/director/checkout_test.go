package director

import (
	"testing"

	"github.com/cristiantebaldi/builder-go/builder"
	"github.com/cristiantebaldi/builder-go/order"
)

func TestCreateSmartphoneBundleAppliesDiscountAndResetsBuilder(t *testing.T) {
	d := NewCheckoutDirector()
	b := builder.NewCreditCardCheckoutBuilder()
	b.SetOrderID("old-order")
	b.AddItem("Old item", 50, 1)
	b.SetCustomerEmail("old@example.com")
	b.SetCardDetails(order.CardDetails{
		Number: "4111111111111111", HolderName: "Customer", Expiration: "12/2026", CVV: "123",
	})

	result, err := d.CreateSmartphoneBundle(
		b,
		"bundle-1",
		"customer@example.com",
		order.Address{Address: "123 Main St", City: "Anytown"},
		order.CardDetails{Number: "4111111111111111", HolderName: "Customer", Expiration: "12/2026", CVV: "123"},
	)
	if err != nil {
		t.Fatalf("CreateSmartphoneBundle() error = %v", err)
	}
	if len(result.Items) != 9 {
		t.Fatalf("got %d items, want 9", len(result.Items))
	}
	if result.OrderID != "bundle-1" || result.Subtotal != 1510 || result.DiscountAmount != 226.5 || result.TotalAmount != 1283.5 {
		t.Fatalf("unexpected bundle result: %+v", result)
	}
	if result.Shipping == nil || result.Shipping.City != "Anytown" {
		t.Fatalf("shipping address was not copied: %+v", result.Shipping)
	}
}
