package travel

import "fmt"

type DomesticFlight struct {
	Origin      string
	Destination string
	Airline     string
}

func (f DomesticFlight) Details() string {
	return fmt.Sprintf("Flight from %s to %s with %s", f.Origin, f.Destination, f.Airline)
}

type DomesticHotel struct {
	Location string
	City     string
}

func (h DomesticHotel) Details() string {
	return fmt.Sprintf("Hotel in %s in %s", h.Location, h.City)
}

type DomesticTransfer struct {
	Type string
}

func (t DomesticTransfer) Details() string {
	return fmt.Sprintf("Transfer of type %s", t.Type)
}

type DomesticFactory struct{}

var _ PackageFactory = DomesticFactory{}

func (DomesticFactory) CreateFlight(origin, destination, airline string) Flight {
	return DomesticFlight{Origin: origin, Destination: destination, Airline: airline}
}

func (DomesticFactory) CreateHotel(location, city string) Hotel {
	return DomesticHotel{Location: location, City: city}
}

func (DomesticFactory) CreateTransfer(transferType string) Transfer {
	return DomesticTransfer{Type: transferType}
}
