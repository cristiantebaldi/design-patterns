namespace WithDirector {
    export interface Address {
        address: string;
        city: string;
    }

    export type PaymentMethod = "credit_card" | "pix";

    export type DiscountType = "percentage" | "fixed";

    export interface OrderCheckoutBuilder {
        reset(): void;

        setOrderId(orderId: string): this;

        addItem(name: string, price: number, quantity: number): this;
    
        setCustomerEmail(email: string): this;

        setShippingAddress(address: Address): this;

        construct(): void;

        getResult(): OrderCheckout;

        applyDiscount(amount: number, type: DiscountType): this;

        getSubtotal(): number;

        getDiscountAmount(): number;
    }
}