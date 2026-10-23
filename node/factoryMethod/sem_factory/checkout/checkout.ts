/// <reference path="../interfaces.ts" />

namespace WithoutFactoryMethod {
    export class Checkout {
        processPayment(amount: number, processor: PaymentProcessor): void {
            try {
                console.log("\nInformações do pagamento:");
                console.log(processor.getPaymentInfo());

                processor.processPayment(amount);
            } catch (error: any) {
                console.error(`Erro ao processar pagamento: ${error.message}`);
            }
        }
    }
}