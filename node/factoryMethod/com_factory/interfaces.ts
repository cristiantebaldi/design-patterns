namespace WithFactoryMethod {
    export interface PaymentProcessor {
        processPayment(amount: number): void;
        getPaymentInfo(): string;
    }

    export interface PaymentProcessorFactory {
        createProcessor(): PaymentProcessor;
    }

    export type PaymentProvider = "pagarme" | "mercadopago" | "pix";
}