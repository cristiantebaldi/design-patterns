/// <reference path="../interfaces.ts" />

namespace WithoutFactoryMethod {
    export class PagarMeProcessor implements PaymentProcessor {
        private apiKey: string;
        private installments: number; // parcelas
        private captureAutomatically: boolean;
        
        constructor(
            apiKey: string, 
            installments: number = 1,
            captureAutomatically: boolean = false
        ) {
            if (!apiKey) {
                throw new Error("PagarMe: API Key é obrigatória");
            }

            if (installments < 1 || installments > 12) {
                throw new Error(
                    "PagarMe: Número de parcelas deve estar entre 1 e 12"
                );
            }

            this.apiKey = apiKey;
            this.installments = installments;
            this.captureAutomatically = captureAutomatically;

            console.log("PagarMe incializado com sucesso!");
        }

        processPayment(amount: number): void {
            console.log(
                `\nProcessando pgto de R$ ${amount} via Cartão de Crédito (PagarMe)`
            );
            console.log(`API Key: ${this.apiKey.substring(0, 4)}...`);
            console.log(`Parcelas: ${this.installments}`);
            console.log(`Captura Automática: ${this.captureAutomatically}`);
            console.log("Status: Pagamento Aprovado");
        }

        getPaymentInfo(): string {
            return "Cartão de Crédito (PagarMe) - Taxa: 2.8%";
        }
    }
}