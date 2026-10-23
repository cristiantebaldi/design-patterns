package travel_test

import (
	"reflect"
	"testing"

	"github.com/cristiantebaldi/abstract-factory-go/travel"
)

func TestBookingSystemCreatesMatchingProductFamilies(t *testing.T) {
	tests := []struct {
		name     string
		factory  travel.PackageFactory
		flight   travel.Flight
		hotel    travel.Hotel
		transfer travel.Transfer
	}{
		{
			name:     "domestic",
			factory:  travel.DomesticFactory{},
			flight:   travel.DomesticFlight{Origin: "Origin", Destination: "Destination", Airline: "Airline"},
			hotel:    travel.DomesticHotel{Location: "Location", City: "City"},
			transfer: travel.DomesticTransfer{Type: "Bus"},
		},
		{
			name:     "international",
			factory:  travel.InternationalFactory{},
			flight:   travel.InternationalFlight{Origin: "Origin", Destination: "Destination", Airline: "Airline"},
			hotel:    travel.InternationalHotel{Location: "Location", City: "City"},
			transfer: travel.InternationalTransfer{Type: "Bus"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			booking := travel.NewBookingSystem(tt.factory)
			pkg := booking.CreateTravelPackage(travel.PackageRequest{
				Origin:        "Origin",
				Destination:   "Destination",
				HotelCity:     "City",
				HotelLocation: "Location",
				Airline:       "Airline",
				TransferType:  "Bus",
			})

			if !reflect.DeepEqual(pkg.Flight, tt.flight) {
				t.Errorf("Flight = %#v, want %#v", pkg.Flight, tt.flight)
			}
			if !reflect.DeepEqual(pkg.Hotel, tt.hotel) {
				t.Errorf("Hotel = %#v, want %#v", pkg.Hotel, tt.hotel)
			}
			if !reflect.DeepEqual(pkg.Transfer, tt.transfer) {
				t.Errorf("Transfer = %#v, want %#v", pkg.Transfer, tt.transfer)
			}
			if got := pkg.Flight.Details(); got != "Flight from Origin to Destination with Airline" {
				t.Errorf("Flight.Details() = %q", got)
			}
			if got := pkg.Hotel.Details(); got != "Hotel in Location in City" {
				t.Errorf("Hotel.Details() = %q", got)
			}
			if got := pkg.Transfer.Details(); got != "Transfer of type Bus" {
				t.Errorf("Transfer.Details() = %q", got)
			}
		})
	}
}
