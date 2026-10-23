namespace WithFactoryMethod {
    export class PixFactory implements PaymentProcessorFactory {
        private readonly pixKey = "email@exemplo.com";
        private readonly merchantName = "Azambuja";

        // Esse createProcessor nunca terá parâmetros senão teremos que que criar a fábrica da fábrica porque o método responsável por criar o objeto já está muito complexo, no máximo passar alguma config nos parâmetros...
        createProcessor(): PaymentProcessor {
            return new PIXProcessor(this.pixKey, this.merchantName);
        }
    }
}