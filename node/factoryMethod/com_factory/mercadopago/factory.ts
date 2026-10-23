namespace WithFactoryMethod {
    export class MercadoPagoFactory implements PaymentProcessorFactory {
        private readonly accessToken = "adklalsjdlaskjdklas";
        private readonly installments = 2;

        // Aqui é simples mas caso fosse muito difícil instância o objeto ficaria tudo aqui
        createProcessor(): PaymentProcessor {
            return new MercadoPagoProcessor(this.accessToken, this.installments);
        }
    }
}