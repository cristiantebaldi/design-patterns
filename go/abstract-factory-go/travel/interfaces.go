package travel

type Flight interface {
	Details() string
}

type Hotel interface {
	Details() string
}

type Transfer interface {
	Details() string
}

type PackageFactory interface {
	CreateFlight(origin, destination, airline string) Flight
	CreateHotel(location, city string) Hotel
	CreateTransfer(transferType string) Transfer
}
