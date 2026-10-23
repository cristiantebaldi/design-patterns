namespace WithoutDirector {

    let creditCardBuilder = new CreditCardCheckoutBuilder();

    try {
        creditCardBuilder
            .setOrderId("1234567890")
            .addItem("Product 1", 100, 1)
            .addItem("Product 2", 200, 2)
            .setCustomerEmail("test@example.com")
            .setShippingAddress({ address: "123 Main St", city: "Anytown" })
            .setCardNumber("1234567890")
            .setCardHolderName("John Doe")
            .setCardExpiration("12/2026")
            .setCardCvv("123")
        
        creditCardBuilder.construct();

        const creditCardOrder = creditCardBuilder.getResult();

        console.log("Order created successfully:");
        console.log(`Order ID: ${creditCardOrder.orderId}`);
        console.log(`Customer: ${creditCardOrder.customerEmail}`);
        console.log(`Total amount: $${creditCardOrder.getTotalAmount()}`);
        console.log(`Payment method: ${creditCardOrder.paymentMethod}`);
        console.log(`Items: ${creditCardOrder.items.length}`);
    } catch (error: any) {
        console.error(`Error creating order: ${error.message}`);
    }
    
    let pixBuilder = new PixCheckoutBuilder();

    try {
        pixBuilder
            .setOrderId("1234567890")
            .addItem("Product 1", 100, 1)
            .addItem("Product 2", 200, 2)
            .setCustomerEmail("test@example.com")
            .setPixKey("1234567890")
            .setPayerName("John Doe")
        
        pixBuilder.construct();

        const pixOrder = pixBuilder.getResult();

        console.log("Order created successfully!");
        console.log(`Order ID: ${pixOrder.orderId}`);
        console.log(`Customer: ${pixOrder.customerEmail}`);
        console.log(`Total amount: $${pixOrder.getTotalAmount()}`);
        console.log(`Payment method: ${pixOrder.paymentMethod}`);
        console.log(`Items: ${pixOrder.items.length}`);
    } catch (error: any) {
        console.error(`Error creating order: ${error.message}`);
    } 
}