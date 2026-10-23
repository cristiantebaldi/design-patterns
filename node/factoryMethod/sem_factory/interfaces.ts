namespace WithoutFactoryMethod {
    export interface PaymentProcessor {
        processPayment(amount: number): void;
        getPaymentInfo(): string;
    }

    export type PaymentProvider = "pagarme" | "mercadopago" | "pix";
}