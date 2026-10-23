namespace WithoutDirector {
    export class CreditCardCheckoutBuilder implements OrderCheckoutBuilder {
        private orderId: string = "";
        private items: Array<{ name: string, price: number, quantity: number }> = [];
        private customerEmail: string = "";
        private shippingAddress?: Address;
        private cardNumber: string = "";
        private cardHolderName: string = "";
        private cardExpiration: string = "";
        private cardCvv: string = "";

        private result: OrderCheckout | null = null;

        reset(): void {
            this.orderId = "";
            this.items = [];
            this.customerEmail = "";
            this.shippingAddress = undefined;
            this.cardNumber = "";
            this.cardHolderName = "";
            this.cardExpiration = "";
            this.cardCvv = "";
            this.result = null;
        }

        setOrderId(orderId: string): this {
            this.orderId = orderId;
            return this;
        }

        addItem(name: string, price: number, quantity: number): this {
            this.items.push({ name, price, quantity });
            return this;
        }

        setCustomerEmail(email: string): this {
            this.customerEmail = email;
            return this;
        }

        setShippingAddress(address: Address): this {
            this.shippingAddress = address;
            return this;
        }
        
        setCardNumber(number: string): this {
            this.cardNumber = number;
            return this;
        }

        setCardHolderName(name: string): this {
            this.cardHolderName = name;
            return this;
        }
        
        setCardExpiration(expiration: string): this {
            this.cardExpiration = expiration;
            return this;
        }

        setCardCvv(cvv: string): this {
            this.cardCvv = cvv;
            return this;
        }
        
        construct(): void {
            if (!this.orderId) {
                throw new Error("Order ID is required");
            }
            if (this.items.length === 0) {
                throw new Error("At least one item is required");
            }
            if (!this.customerEmail) {
                throw new Error("Customer email is required");
            }

            if (!this.cardNumber || !this.cardHolderName || !this.cardExpiration || !this.cardCvv) {
                throw new Error("Card details are incomplete");
            }

            this.result = new OrderCheckout(
                this.orderId, 
                this.items, 
                this.customerEmail, 
                "credit_card",
                {
                    cardNumber: this.cardNumber,
                    cardHolderName: this.cardHolderName,
                    cardExpiration: this.cardExpiration,
                    cardCvv: this.cardCvv,
                },
                this.shippingAddress
            );
        }

        getResult(): OrderCheckout {
            if (!this.result) {
                throw new Error("You must call construct() before getResult()");
            }
            return this.result;
        }
    }
}