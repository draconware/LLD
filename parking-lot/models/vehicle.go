package models

type IVehicle interface {
	GetLicenseNumber() string
	GetParkingTicket() *ParkingTicket
	AssignParkingTicket(*ParkingTicket)
}

type Vehicle struct {
	licenseNumber string
	ticket        *ParkingTicket
}

func (v *Vehicle) GetLicenseNumber() string {
	return v.licenseNumber
}
func (v *Vehicle) AssignParkingTicket(ticket *ParkingTicket) {
	v.ticket = ticket
}
func (v *Vehicle) GetParkingTicket() *ParkingTicket {
	return v.ticket
}

type VanVehicle struct {
	Vehicle
}
type TruckVehicle struct {
	Vehicle
}
type CarVehicle struct {
	Vehicle
}
type BikeVehicle struct {
	Vehicle
}

type VehicleType string

const (
	VanVehicleType   VehicleType = "vanvehicle"
	TruckVehicleType VehicleType = "truckvehicle"
	CarVehicleType   VehicleType = "carvehicle"
	BikeVehicleType  VehicleType = "bikevehicle"
)

type VehicleFactory struct{}

func NewVehicleFactory() *VehicleFactory {
	return &VehicleFactory{}
}

func (vf *VehicleFactory) GetVehicle(licenseNumber string, vehicleType VehicleType) IVehicle {
	vehicle := Vehicle{
		licenseNumber: licenseNumber,
	}
	switch vehicleType {
	case VanVehicleType:
		return &VanVehicle{Vehicle: vehicle}
	case TruckVehicleType:
		return &TruckVehicle{Vehicle: vehicle}
	case CarVehicleType:
		return &CarVehicle{Vehicle: vehicle}
	case BikeVehicleType:
		return &BikeVehicle{Vehicle: vehicle}
	default:
		return nil
	}
}
