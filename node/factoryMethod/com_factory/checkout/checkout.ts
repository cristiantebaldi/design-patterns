/// <reference path="../interfaces.ts" />

namespace WithFactoryMethod {
    export class CheckoutFactory {
        constructor(private processorFactory: PaymentProcessorFactory) {}

        processPayment(amount: number): void {
            try {
                const processor = this.processorFactory.createProcessor();
                
                console.log("\nInformações do pagamento:");
                console.log(processor.getPaymentInfo());

                processor.processPayment(amount);
            } catch (error: any) {
                console.error(`Erro ao processar pagamento: ${error.message}`);
            }
        }
    }
}