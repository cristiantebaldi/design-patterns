/// <reference path="../interfaces.ts" />

namespace WithFactoryMethod {
    export class PIXProcessor implements PaymentProcessor {
        private pixKey: string;
        private merchantName: string; 
        
        constructor(pixKey: string, merchantName: string) {
            if (!pixKey) {
                throw new Error("PIX: Chave API é obrigatório");
            }

            if (!merchantName) {
                throw new Error("PIX: Nome do comerciante é obrigatório");
            }

            this.pixKey = pixKey;
            this.merchantName = merchantName;


            console.log("PIX incializado com sucesso!");
        }

        processPayment(amount: number): void {
            console.log(
                `\nProcessando pgto de R$ ${amount} via PIX`
            );
            console.log(`Chave PIX: ${this.pixKey.substring(0, 4)}...`);
            console.log(`Recebedor: ${this.merchantName}`);
            console.log("Gerando QR Code...");
            console.log("Status: QR Code gerado com sucesso");
        }

        getPaymentInfo(): string {
            return "PIX - Taxa: 0% - Processamento instantâneo";
        }
    }
}