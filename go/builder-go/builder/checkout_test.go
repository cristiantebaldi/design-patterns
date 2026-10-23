package builder

import (
	"strings"
	"testing"

	"github.com/cristiantebaldi/builder-go/order"
)

func TestCreditCardBuilderBuildsOrderAndAppliesPercentageDiscount(t *testing.T) {
	b := NewCreditCardCheckoutBuilder()
	b.SetOrderID("order-1")
	b.AddItem("Notebook", 1000, 1)
	b.AddItem("Mouse", 100, 2)
	b.SetCustomerEmail("customer@example.com")
	b.SetCardDetails(order.CardDetails{
		Number: "4111111111111111", HolderName: "Customer", Expiration: "12/2026", CVV: "123",
	})
	b.ApplyDiscount(15, order.Percentage)

	result, err := b.Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	if result.Subtotal != 1200 || result.DiscountAmount != 180 || result.TotalAmount != 1020 {
		t.Fatalf("unexpected totals: subtotal=%v discount=%v total=%v", result.Subtotal, result.DiscountAmount, result.TotalAmount)
	}
	if result.PaymentMethod != order.CreditCard || result.CardDetails == nil {
		t.Fatalf("unexpected payment details: %+v", result)
	}
}

func TestPixBuilderRequiresPaymentDetails(t *testing.T) {
	b := NewPixCheckoutBuilder()
	b.SetOrderID("order-1")
	b.AddItem("Headphones", 200, 1)
	b.SetCustomerEmail("customer@example.com")

	_, err := b.Build()
	if err == nil || !strings.Contains(err.Error(), "PIX details") {
		t.Fatalf("Build() error = %v, want missing PIX details error", err)
	}
}

func TestGetResultRequiresSuccessfulBuildAndResetClearsState(t *testing.T) {
	b := NewPixCheckoutBuilder()
	if _, err := b.GetResult(); err == nil {
		t.Fatal("GetResult() should fail before Build()")
	}

	b.SetOrderID("order-1")
	b.AddItem("Headphones", 200, 1)
	b.SetCustomerEmail("customer@example.com")
	b.SetPixDetails(order.PixDetails{Key: "demo-key", PayerName: "Customer"})
	if _, err := b.Build(); err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	b.Reset()
	if _, err := b.GetResult(); err == nil {
		t.Fatal("GetResult() should fail after Reset()")
	}
	if _, err := b.Build(); err == nil {
		t.Fatal("Build() should fail after Reset()")
	}
}

func TestBuilderRejectsDiscountGreaterThanSubtotal(t *testing.T) {
	b := NewPixCheckoutBuilder()
	b.SetOrderID("order-1")
	b.AddItem("Headphones", 200, 1)
	b.SetCustomerEmail("customer@example.com")
	b.SetPixDetails(order.PixDetails{Key: "demo-key", PayerName: "Customer"})
	b.ApplyDiscount(201, order.Fixed)

	if _, err := b.Build(); err == nil {
		t.Fatal("Build() should reject discount greater than subtotal")
	}
}
