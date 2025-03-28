package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mastik5h/LLD/parking-lot/models"
)

type parkingLot struct {
	id             int
	name           string
	parkingRate    *models.ParkingRate
	entranceGates  map[string]*models.EntranceGate
	exitGates      map[string]*models.ExitGate
	parkingTickets map[string]*models.ParkingTicket

	vehicleFactory *models.VehicleFactory
	paymentFactory *models.PaymentFactory
}

func newParkingLot() *parkingLot {
	parkingLot := &parkingLot{
		id:             rand.Int(),
		name:           "ParkingLot1",
		parkingRate:    models.NewParkingRate(),
		entranceGates:  make(map[string]*models.EntranceGate),
		exitGates:      make(map[string]*models.ExitGate),
		parkingTickets: make(map[string]*models.ParkingTicket),
		vehicleFactory: models.NewVehicleFactory(),
		paymentFactory: models.NewPaymentFactory(),
	}
	entranceGate1 := models.NewEntranceGate()
	entranceGate2 := models.NewEntranceGate()

	exitGate1 := models.NewExitGate()
	exitGate2 := models.NewExitGate()

	parkingLot.entranceGates["entryGate1"] = entranceGate1
	parkingLot.entranceGates["entryGate2"] = entranceGate2

	parkingLot.exitGates["exitGate1"] = exitGate1
	parkingLot.exitGates["exitGate2"] = exitGate2

	return parkingLot
}

var parkingLotInstance *parkingLot

func GetParkingLotInstance() *parkingLot {
	if parkingLotInstance == nil {
		parkingLotInstance = newParkingLot()
	}
	return parkingLotInstance
}

func (pl *parkingLot) ParkVehicle(vehicleNumber string, vehicleType models.VehicleType, entranceGateName string) *models.ParkVehicleResponse {
	if len(vehicleNumber) == 0 {
		fmt.Println("Invalid vehiclenumber.")
		return nil
	}
	var entrance *models.EntranceGate
	for k, v := range pl.entranceGates {
		if k == entranceGateName {
			entrance = v
			break
		}
	}
	if entrance == nil {
		fmt.Println("Invalid entrance gate.")
		return nil
	}

	vehicle := pl.vehicleFactory.GetVehicle(vehicleNumber, vehicleType)
	if vehicle == nil {
		fmt.Println("Invalid vehicle type.")
		return nil
	}

	entrance.GiveTicket(vehicle)
	parkingTicket := vehicle.GetParkingTicket()
	pl.parkingTickets[parkingTicket.GetTicketNumber()] = parkingTicket

	return models.NewParkingTicketResponse(parkingTicket, pl.parkingRate)
}

func (pl *parkingLot) ExitParking(parkingTicketNumber string, exitGateName string, paymentMode models.PaymentMode) *models.ExitParkingResponse {
	var parkingTicket *models.ParkingTicket
	for k, v := range pl.parkingTickets {
		if k == parkingTicketNumber {
			parkingTicket = v
			break
		}
	}
	if parkingTicket == nil {
		fmt.Println("Invalid parking ticket number.")
		return nil
	}

	var exitGate *models.ExitGate
	for k, v := range pl.exitGates {
		if k == exitGateName {
			exitGate = v
			break
		}
	}
	if exitGate == nil {
		fmt.Println("Invalid exit gate found.")
		return nil
	}

	exitTime := time.Now()
	totalParkingHours := exitTime.Sub(parkingTicket.GetEntryTime()).Hours()
	totalParkingAmount := pl.parkingRate.Calculate(totalParkingHours)
	paymentMethod := pl.paymentFactory.GetPaymentMethod(totalParkingAmount, paymentMode)

	parkingTicket.SetAmount(totalParkingAmount)
	parkingTicket.SetPayment(paymentMethod)
	parkingTicket.SetExitTime(exitTime)

	exitGate.ValidateTicket(parkingTicket)

	return models.NewExitParkingResponse(parkingTicket, pl.parkingRate)
}
