interface Flight {
    origin: string;
    destination: string;
    airline: string;
    getDetails(): string;
}

interface Hotel {
    location: string;
    city: string;
    getDetails(): string;
}

interface Transfer {
    type: string;
    getDetails(): string;
}

interface TravelPackageFactory {
    createFlight(origin: string, destination: string, airline: string): Flight;
    createHotel(location: string, city: string): Hotel;
    createTransfer(type: string): Transfer;
}

class DomesticFlight implements Flight {
    constructor(
        public origin: string, 
        public destination: string, 
        public airline: string) {}
    getDetails(): string {
        return `Flight from ${this.origin} to ${this.destination} with ${this.airline}`;
    }
}

class InternationalFlight implements Flight {
    constructor(
        public origin: string, 
        public destination: string, 
        public airline: string) {}
    getDetails(): string {
        return `Flight from ${this.origin} to ${this.destination} with ${this.airline}`;
    }
}

class DomesticHotel implements Hotel {
    constructor(
        public location: string, 
        public city: string) {}
    getDetails(): string {
        return `Hotel in ${this.location} in ${this.city}`;
    }
}

class InternationalHotel implements Hotel {
    constructor(
        public location: string, 
        public city: string) {}
    getDetails(): string {
        return `Hotel in ${this.location} in ${this.city}`;
    }
}

class DomesticTransfer implements Transfer {
    constructor(public type: string) {}
    getDetails(): string {
        return `Transfer of type ${this.type}`;
    }
}

class InternationalTransfer implements Transfer {
    constructor(public type: string) {}
    getDetails(): string {
        return `Transfer of type ${this.type}`;
    }
}

class DomesticTravelPackageFactory implements TravelPackageFactory {
    createFlight(origin: string, destination: string, airline: string): Flight {
        return new DomesticFlight(origin, destination, airline);
    }
    createHotel(location: string, city: string): Hotel {
        return new DomesticHotel(location, city);
    }
    createTransfer(type: string): Transfer {
        return new DomesticTransfer(type);
    }
}

class InternationalTravelPackageFactory implements TravelPackageFactory {
    createFlight(origin: string, destination: string, airline: string): Flight {
        return new InternationalFlight(origin, destination, airline);
    }
    createHotel(location: string, city: string): Hotel {
        return new InternationalHotel(location, city);
    }
    createTransfer(type: string): Transfer {
        return new InternationalTransfer(type);
    }
}

class BookingSystem {
    private factory: TravelPackageFactory;
    constructor(factory: TravelPackageFactory) {
        this.factory = factory;
    }

    createTravelPackage(
        origin: string, 
        destination: string, 
        HotelCity: string, 
        HotelLocation: string, 
        airline: string, 
        type: string
    ): void {
        const flight = this.factory.createFlight(origin, destination, airline);
        const hotel = this.factory.createHotel(HotelLocation, HotelCity);
        const transfer = this.factory.createTransfer(type);
        console.log(flight.getDetails());
        console.log(hotel.getDetails());
        console.log(transfer.getDetails());
    }
}

const domesticFactory = new DomesticTravelPackageFactory();
const internationalFactory = new InternationalTravelPackageFactory();

const domesticBookingSystem = new BookingSystem(domesticFactory);

const internationalBookingSystem = new BookingSystem(internationalFactory);

domesticBookingSystem.createTravelPackage(
    "New York", 
    "Los Angeles", 
    "New York", 
    "Los Angeles", 
    "Delta", 
    "Bus"
);

internationalBookingSystem.createTravelPackage(
    "Paris", 
    "London", 
    "Paris", 
    "London", 
    "Air France", 
    "Train"
);