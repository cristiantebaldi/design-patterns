namespace WithFactoryMethod {
    export class PagarMeFactory implements PaymentProcessorFactory {
        private readonly apiKey = "alksdjaslkdjaldkaçsldk";
        private readonly installments = 3;
        private readonly captureAutomatically = true;

        createProcessor(): PaymentProcessor {
            return new PagarMeProcessor(
                this.apiKey,
                this.installments,
                this.captureAutomatically,
            );
        }
    }
}