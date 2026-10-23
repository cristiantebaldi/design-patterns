namespace WithDirector {
    export class PixCheckoutBuilder implements OrderCheckoutBuilder {
        private orderId: string = "";
        private items: Array<{ name: string, price: number, quantity: number }> = [];
        private customerEmail: string = "";
        private shippingAddress?: Address;
        private pixKey: string = "";
        private payerName: string = "";

        private result: OrderCheckout | null = null;

        reset(): void {
            this.orderId = "";
            this.items = [];
            this.customerEmail = "";
            this.shippingAddress = undefined;
            this.pixKey = "";
            this.payerName = "";
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

        setPixKey(key: string): this {
            this.pixKey = key;
            return this;
        }

        setPayerName(name: string): this {
            this.payerName = name;
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
            if (!this.pixKey || !this.payerName) {
                throw new Error("PIX details are incomplete");
            }

            this.result = new OrderCheckout(
                this.orderId, 
                this.items, 
                this.customerEmail, 
                "pix",
                {
                    pixKey: this.pixKey,
                    payerName: this.payerName,
                },
                this.shippingAddress, 
            );
        }

        getResult(): OrderCheckout {
            if (!this.result) {
                throw new Error("You must call construct() before getResult()");
            }
            return this.result;
        }

        applyDiscount(amount: number, type: DiscountType): this {
            if (type === "percentage") {
                this.result?.applyDiscount(amount, "percentage");
            } else {
                this.result?.applyDiscount(amount, "fixed");
            }
            return this;
        }

        getSubtotal(): number {
            return this.result?.getSubtotal() ?? 0;
        }

        getDiscountAmount(): number {
            return this.result?.getDiscountAmount() ?? 0;
        }
    }
}