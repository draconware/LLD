package main

import (
	"fmt"

	"github.com/mastik5h/LLD/parking-lot/models"
	"github.com/mastik5h/LLD/parking-lot/services"
)

func main() {
	// Create new parking lot

	parkingLot := services.GetParkingLotInstance()

	// Park vehicle of type car with vehicle number: abc

	response1 := parkingLot.ParkVehicle("abc1", models.CarVehicleType, "entryGate1")
	if response1 == nil {
		return
	}
	fmt.Printf("Parked vehicle of type car with response: %v\n", response1)

	response2 := parkingLot.ParkVehicle("abc2", models.BikeVehicleType, "entryGate1")
	if response2 == nil {
		return
	}
	fmt.Printf("Parked vehicle of type bike with response: %v\n", response2)

	response3 := parkingLot.ParkVehicle("abc3", models.TruckVehicleType, "entryGate2")
	if response3 == nil {
		return
	}
	fmt.Printf("Parked vehicle of type truck with response: %v\n", response3)

	exitResponse1 := parkingLot.ExitParking(response1.TicketNumber, "exitGate1", models.CashPaymentMode)
	fmt.Printf("Exit vehicle of type car with response: %v\n", exitResponse1)

	exitResponse2 := parkingLot.ExitParking(response2.TicketNumber, "exitGate1", models.CardPaymentMode)
	fmt.Printf("Exit vehicle of type bike with response: %v\n", exitResponse2)

	exitResponse3 := parkingLot.ExitParking(response3.TicketNumber, "exitGate2", models.CardPaymentMode)
	fmt.Printf("Exit vehicle of type truck with response: %v\n", exitResponse3)
}
