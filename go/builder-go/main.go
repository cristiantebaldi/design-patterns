package main

import (
	"fmt"
	"log"

	"github.com/cristiantebaldi/builder-go/builder"
	"github.com/cristiantebaldi/builder-go/director"
	"github.com/cristiantebaldi/builder-go/order"
)

func main() {
	fmt.Println("=== Builder sem Director ===")
	runWithoutDirector()

	fmt.Println("\n=== Builder com Director ===")
	runWithDirector()
}

func runWithoutDirector() {
	cardBuilder := builder.NewCreditCardCheckoutBuilder()
	cardBuilder.SetOrderID("1234567890")
	cardBuilder.AddItem("Smartphone", 1000, 1)
	cardBuilder.AddItem("Case", 100, 2)
	cardBuilder.SetCustomerEmail("test@example.com")
	cardBuilder.SetShippingAddress(order.Address{Address: "123 Main St", City: "Anytown"})
	cardBuilder.SetCardDetails(demoCardDetails())
	cardOrder, err := cardBuilder.Build()
	if err != nil {
		log.Printf("erro ao criar pedido com cartão: %v", err)
		return
	}
	printOrder(cardOrder)

	pixBuilder := builder.NewPixCheckoutBuilder()
	pixBuilder.SetOrderID("9876543210")
	pixBuilder.AddItem("Headphones", 200, 1)
	pixBuilder.SetCustomerEmail("test@example.com")
	pixBuilder.SetPixDetails(order.PixDetails{Key: "1234567890", PayerName: "John Doe"})
	pixOrder, err := pixBuilder.Build()
	if err != nil {
		log.Printf("erro ao criar pedido PIX: %v", err)
		return
	}
	printOrder(pixOrder)
}

func runWithDirector() {
	checkoutDirector := director.NewCheckoutDirector()
	cardBuilder := builder.NewCreditCardCheckoutBuilder()

	basicCardOrder, err := checkoutDirector.CreateBasicCreditCardCheckout(
		cardBuilder,
		"1234567890",
		"john.doe@example.com",
		"Smartphone",
		1000,
		demoCardDetails(),
	)
	if err != nil {
		log.Printf("erro ao criar pedido básico com cartão: %v", err)
		return
	}
	printOrder(basicCardOrder)

	pixBuilder := builder.NewPixCheckoutBuilder()
	pixOrder, err := checkoutDirector.CreateBasicPixCheckout(
		pixBuilder,
		"9876543210",
		"jane.smith@example.com",
		"Smartphone",
		1000,
		order.PixDetails{Key: "demo-pix-key", PayerName: "Jane Smith"},
	)
	if err != nil {
		log.Printf("erro ao criar pedido PIX: %v", err)
		return
	}
	printOrder(pixOrder)

	bundle, err := checkoutDirector.CreateSmartphoneBundle(
		cardBuilder,
		"1122334455",
		"john.doe@example.com",
		order.Address{Address: "123 Main St", City: "Anytown"},
		demoCardDetails(),
	)
	if err != nil {
		log.Printf("erro ao criar pacote de smartphone: %v", err)
		return
	}
	printOrder(bundle)
	fmt.Printf("Subtotal: $%.2f\nDesconto: $%.2f\n", bundle.Subtotal, bundle.DiscountAmount)
}

func demoCardDetails() order.CardDetails {
	return order.CardDetails{
		Number:     "4111111111111111",
		HolderName: "John Doe",
		Expiration: "12/2026",
		CVV:        "123",
	}
}

func printOrder(checkout *order.Checkout) {
	fmt.Printf("Resumo: %s\n", checkout.Summary())
	fmt.Printf("Cliente: %s\n", checkout.CustomerEmail)
	fmt.Printf("Itens: %d\n", len(checkout.Items))
}
