namespace WithDirector {

    const director = new CheckoutDirector();

    const creditCardBuilder = new CreditCardCheckoutBuilder();
    creditCardBuilder
        .setCardNumber("1234567890")
        .setCardHolderName("John Doe")
        .setCardExpiration("12/2026")
        .setCardCvv("123");

    director.createBasicCreditCardCheckout(
        creditCardBuilder,
        "1234567890",
        "john.doe@example.com",
        "Smartphone",
        1000
    );

    const basicCreditCardOrder = creditCardBuilder.getResult();

    console.log(`Order Summary: ${basicCreditCardOrder.getSummary()}`);
    console.log(
        `Total Amount: ${basicCreditCardOrder.getTotalAmount().toFixed(2)}`
    );


    const pixBuilder = new PixCheckoutBuilder();

    pixBuilder.setPixKey("1234567890").setPayerName("Jane Smith");

    director.createBasicPixCheckout(
        pixBuilder,
        "1234567890",
        "jane.smith@example.com",
        "Smartphone",
        1000
    );

    const pixOrder = pixBuilder.getResult();
    
    console.log(`Order Summary: ${pixOrder.getSummary()}`);
    console.log(
        `Total Amount: ${pixOrder.getTotalAmount().toFixed(2)}`
    );


    const shippingAddress = {address: "123 Main St", city: "Anytown"};

    director.createSmartphoneBundle(
        creditCardBuilder,
        "1234567890",
        "john.doe@example.com",
        shippingAddress
    );

    const smartphoneBundle = creditCardBuilder.getResult();

    console.log(`Order Summary: ${smartphoneBundle.getSummary()}`);
    console.log(`Subtotal: ${smartphoneBundle.getSubtotal().toFixed(2)}`);
    console.log(`Discount: ${smartphoneBundle.getDiscountAmount().toFixed(2)}`);
    console.log(
        `Total Amount: ${smartphoneBundle.getTotalAmount().toFixed(2)}`
    );
}