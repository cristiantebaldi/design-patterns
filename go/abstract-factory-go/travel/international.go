package travel

import "fmt"

type InternationalFlight struct {
	Origin      string
	Destination string
	Airline     string
}

func (f InternationalFlight) Details() string {
	return fmt.Sprintf("Flight from %s to %s with %s", f.Origin, f.Destination, f.Airline)
}

type InternationalHotel struct {
	Location string
	City     string
}

func (h InternationalHotel) Details() string {
	return fmt.Sprintf("Hotel in %s in %s", h.Location, h.City)
}

type InternationalTransfer struct {
	Type string
}

func (t InternationalTransfer) Details() string {
	return fmt.Sprintf("Transfer of type %s", t.Type)
}

type InternationalFactory struct{}

var _ PackageFactory = InternationalFactory{}

func (InternationalFactory) CreateFlight(origin, destination, airline string) Flight {
	return InternationalFlight{Origin: origin, Destination: destination, Airline: airline}
}

func (InternationalFactory) CreateHotel(location, city string) Hotel {
	return InternationalHotel{Location: location, City: city}
}

func (InternationalFactory) CreateTransfer(transferType string) Transfer {
	return InternationalTransfer{Type: transferType}
}
