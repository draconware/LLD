package models

type IParkingSpot interface {
	getIsSpotFree() bool
	assignVehicle(IVehicle)
}

type ParkingSpot struct {
	isSpotFree bool
	vehicle    IVehicle
}

func (ps *ParkingSpot) getIsSpotFree() bool {
	return ps.isSpotFree
}

type HandiCappedParkingSpot struct {
	ParkingSpot
}

func (hcps *HandiCappedParkingSpot) assignVehicle(vehicle IVehicle) {
	hcps.vehicle = vehicle
}

type LargeParkingSpot struct {
	ParkingSpot
}

func (lps *LargeParkingSpot) assignVehicle(vehicle IVehicle) {
	lps.vehicle = vehicle
}

type CompactParkingSpot struct {
	ParkingSpot
}

func (cps *CompactParkingSpot) assignVehicle(vehicle IVehicle) {
	cps.vehicle = vehicle
}

type BikeParkingSpot struct {
	ParkingSpot
}

func (bps *BikeParkingSpot) assignVehicle(vehicle IVehicle) {
	bps.vehicle = vehicle
}

type ParkingSpotType string

const (
	HandiCappedParkingSpotType ParkingSpotType = "handicappedspot"
	LargeParkingSpotType       ParkingSpotType = "largespot"
	CompactParkingSpotType     ParkingSpotType = "compactspot"
	BikeParkingSpotType        ParkingSpotType = "bikespot"
)

type ParkingSpotFactory struct{}

func (psf *ParkingSpotFactory) GetParkingSpot(spotType ParkingSpotType) IParkingSpot {
	parkingSpot := ParkingSpot{
		isSpotFree: true,
	}
	switch spotType {
	case HandiCappedParkingSpotType:
		return &HandiCappedParkingSpot{ParkingSpot: parkingSpot}
	case LargeParkingSpotType:
		return &LargeParkingSpot{ParkingSpot: parkingSpot}
	case CompactParkingSpotType:
		return &CompactParkingSpot{ParkingSpot: parkingSpot}
	case BikeParkingSpotType:
		return &BikeParkingSpot{ParkingSpot: parkingSpot}
	default:
		return nil
	}
}
