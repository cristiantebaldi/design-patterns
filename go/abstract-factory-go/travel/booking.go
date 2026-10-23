package travel

type PackageRequest struct {
	Origin        string
	Destination   string
	HotelCity     string
	HotelLocation string
	Airline       string
	TransferType  string
}

type Package struct {
	Flight   Flight
	Hotel    Hotel
	Transfer Transfer
}

type BookingSystem struct {
	factory PackageFactory
}

func NewBookingSystem(factory PackageFactory) *BookingSystem {
	return &BookingSystem{factory: factory}
}

func (b *BookingSystem) CreateTravelPackage(request PackageRequest) Package {
	return Package{
		Flight:   b.factory.CreateFlight(request.Origin, request.Destination, request.Airline),
		Hotel:    b.factory.CreateHotel(request.HotelLocation, request.HotelCity),
		Transfer: b.factory.CreateTransfer(request.TransferType),
	}
}
