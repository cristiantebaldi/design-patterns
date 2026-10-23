namespace WithDirector {
    export class CheckoutDirector {
        createBasicCreditCardCheckout(
            builder: OrderCheckoutBuilder,
            orderId: string,
            customerEmail: string,
            productName: string,
            productPrice: number      
        ): void {
            builder.reset();

            builder
                .setOrderId(orderId)
                .setCustomerEmail(customerEmail)
                .addItem(productName, productPrice, 1);
            
            builder.construct();
        }


        createBasicPixCheckout(
            builder: OrderCheckoutBuilder,
            orderId: string,
            customerEmail: string,
            productName: string,
            productPrice: number,
        ): void {
            builder.reset();

            builder
                .setOrderId(orderId)
                .setCustomerEmail(customerEmail)
                .addItem(productName, productPrice, 1);
            
            builder.construct();
        }

        createSmartphoneBundle(
            builder: OrderCheckoutBuilder,
            orderId: string,
            customerEmail: string,
            shippingAddress: Address,
        ): void {
            builder.reset();

            builder
                .setOrderId(orderId)
                .setCustomerEmail(customerEmail)
                .setShippingAddress(shippingAddress);
            
            builder
                .addItem("Smartphone", 1000, 1)
                .addItem("Case", 100, 1)
                .addItem("Screen Protector", 10, 1)
                .addItem("Charging Cable", 20, 1)
                .addItem("Headphones", 50, 1)
                .addItem("Earbuds", 30, 1)
                .addItem("Power Bank", 100, 1)
                .addItem("Wireless Charger", 100, 1)
                .addItem("Bluetooth Speaker", 100, 1);
            
            builder.applyDiscount(15, "percentage");

            builder.construct();
        }

        createGamerPackage(
            orderId: string,
            customerEmail: string,
            shippingAddress: Address,
            builder: OrderCheckoutBuilder,
        ): void {
            builder.reset();

            builder
                .setOrderId(orderId)
                .setCustomerEmail(customerEmail)
                .setShippingAddress(shippingAddress);
                
            builder
                .addItem("Gaming Mouse", 100, 1)
                .addItem("Gaming Keyboard", 100, 1)
                .addItem("Gaming Headset", 100, 1)
                .addItem("Gaming Mousepad", 100, 1)
                .addItem("Gaming Chair", 100, 1)
                .addItem("Gaming Table", 100, 1);

            builder.applyDiscount(15, "percentage");

            builder.construct();
        }
    }
}