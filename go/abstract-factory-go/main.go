package main

import (
	"fmt"

	"github.com/cristiantebaldi/abstract-factory-go/travel"
)

func main() {
	domesticBooking := travel.NewBookingSystem(travel.DomesticFactory{})
	internationalBooking := travel.NewBookingSystem(travel.InternationalFactory{})

	fmt.Println("=== Domestic travel package ===")
	printPackage(domesticBooking.CreateTravelPackage(travel.PackageRequest{
		Origin:        "New York",
		Destination:   "Los Angeles",
		HotelCity:     "Los Angeles",
		HotelLocation: "Downtown",
		Airline:       "Delta",
		TransferType:  "Bus",
	}))

	fmt.Println("\n=== International travel package ===")
	printPackage(internationalBooking.CreateTravelPackage(travel.PackageRequest{
		Origin:        "Paris",
		Destination:   "London",
		HotelCity:     "London",
		HotelLocation: "City Centre",
		Airline:       "Air France",
		TransferType:  "Train",
	}))
}

func printPackage(pkg travel.Package) {
	fmt.Println(pkg.Flight.Details())
	fmt.Println(pkg.Hotel.Details())
	fmt.Println(pkg.Transfer.Details())
}
