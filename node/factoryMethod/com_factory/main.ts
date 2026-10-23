/**
 * Esse código aparentemente está correto e não é que está errado
 * na verdade ele está bem simples, hoje o que modifica entre um objeto e outro são apenas parâmetros
 * o grande problemas que podem existir projetos que para instanciar um objeto precisa declarar N objetos ou necessecitam de variáveis de ambiente, por isso tomar cuidado para não utilizar o fatctory method toda hora.
 */

const checkoutPagarme = new WithFactoryMethod.CheckoutFactory(
    new WithFactoryMethod.PagarMeFactory()
);
checkoutPagarme.processPayment(1000);

const checkoutMercadoPago = new WithFactoryMethod.CheckoutFactory(
    new WithFactoryMethod.MercadoPagoFactory()
);
checkoutMercadoPago.processPayment(1500);

const checkoutPix = new WithFactoryMethod.CheckoutFactory(
    new WithFactoryMethod.MercadoPagoFactory()
);
checkoutMercadoPago.processPayment(2000);