/**
 * Esse código aparentemente está correto e não é que está errado
 * na verdade ele está bem simples, hoje o que modifica entre um objeto e outro são apenas parâmetros
 * o grande problemas que podem existir projetos que para instanciar um objeto precisa declarar N objetos ou necessecitam de variáveis de ambiente, por isso tomar cuidado para não utilizar o fatctory method toda hora.
 */

const checkout = new WithoutFactoryMethod.Checkout();
const pagarMeProcessor = new WithoutFactoryMethod.PagarMeProcessor(
    "çsdkljfalskflkasdjf",
    3,
    true
);
checkout.processPayment(1000, pagarMeProcessor);

const mercadoPagoProcessor = new WithoutFactoryMethod.MercadoPagoProcessor(
    "lajdalskjdlaskjdals",
    2
);
checkout.processPayment(1500, mercadoPagoProcessor);

const pixProcessor = new WithoutFactoryMethod.PIXProcessor(
    "email@exemplo.com",
    "Azambuja"
);
checkout.processPayment(2000, pixProcessor);