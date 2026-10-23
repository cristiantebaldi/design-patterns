namespace WithoutDirector {
    export class OrderCheckout {
        constructor(
            public readonly orderId: string,
            public readonly items: Array<{
                name: string,
                price: number,
                quantity: number,
            }>,
            public readonly customerEmail: string,
            public readonly paymentMethod: PaymentMethod,
            public readonly paymentDetails: Record<string, any>,
            public readonly shippingAddress?: Address
        ) {}

        getTotalAmount(): number {
            return this.items.reduce(
                (sum, item) => sum + item.price * item.quantity,
                0
            );
        }

        getSummary(): string {
            return `Order #${this.orderId} - ${
                this.paymentMethod
            } - $${this.getTotalAmount().toFixed(2)}`;
        }
    }
}