package main

import (
	"fmt"
	"log"

	"github.com/cristiantebaldi/factory-method-go/payment"
	"github.com/cristiantebaldi/factory-method-go/providers/mercadopago"
	"github.com/cristiantebaldi/factory-method-go/providers/pagarme"
	"github.com/cristiantebaldi/factory-method-go/providers/pix"
)

func main() {
	fmt.Println("=== Sem Factory Method ===")
	runWithoutFactoryMethod()

	fmt.Println("\n=== Com Factory Method ===")
	runWithFactoryMethod()
}

func runWithoutFactoryMethod() {
	checkout := payment.NewCheckout()

	pagarMeProcessor, err := pagarme.NewProcessor("demo-pagarme-key", 3, true)
	if err != nil {
		log.Printf("erro ao criar processador PagarMe: %v", err)
		return
	}
	if err := checkout.ProcessPayment(1000, pagarMeProcessor); err != nil {
		log.Printf("erro ao processar pagamento: %v", err)
	}

	mercadoPagoProcessor, err := mercadopago.NewProcessor("demo-mercadopago-token", 2)
	if err != nil {
		log.Printf("erro ao criar processador MercadoPago: %v", err)
		return
	}
	if err := checkout.ProcessPayment(1500, mercadoPagoProcessor); err != nil {
		log.Printf("erro ao processar pagamento: %v", err)
	}

	pixProcessor, err := pix.NewProcessor("email@exemplo.com", "Azambuja")
	if err != nil {
		log.Printf("erro ao criar processador PIX: %v", err)
		return
	}
	if err := checkout.ProcessPayment(2000, pixProcessor); err != nil {
		log.Printf("erro ao processar pagamento: %v", err)
	}
}

func runWithFactoryMethod() {
	checkout := payment.NewFactoryCheckout(pagarme.NewFactory())
	if err := checkout.ProcessPayment(1000); err != nil {
		log.Printf("erro ao processar pagamento PagarMe: %v", err)
	}

	checkout = payment.NewFactoryCheckout(mercadopago.NewFactory())
	if err := checkout.ProcessPayment(1500); err != nil {
		log.Printf("erro ao processar pagamento MercadoPago: %v", err)
	}

	checkout = payment.NewFactoryCheckout(pix.NewFactory())
	if err := checkout.ProcessPayment(2000); err != nil {
		log.Printf("erro ao processar pagamento PIX: %v", err)
	}
}
