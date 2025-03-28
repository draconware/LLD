package models

import (
	"math/rand"
	"strconv"
)

type EntranceGate struct {
	id string
}

func (eng *EntranceGate) GiveTicket(vehicle IVehicle) {
	parkingTicket := newParkingTicket(vehicle, eng)
	vehicle.AssignParkingTicket(parkingTicket)
}
func (eng *EntranceGate) GetId() string {
	return eng.id
}

func NewEntranceGate() *EntranceGate {
	return &EntranceGate{
		id: strconv.Itoa(rand.Int()),
	}
}

type ExitGate struct {
	id string
}

func (exg *ExitGate) ValidateTicket(parkingTicket *ParkingTicket) {
	// validate payment status

	paymentStatus := parkingTicket.payment.InitiateTransaction()

	parkingTicket.SetPaymentStatus(paymentStatus)
	parkingTicket.SetExit(exg)
}
func (exg *ExitGate) GetId() string {
	return exg.id
}

func NewExitGate() *ExitGate {
	return &ExitGate{
		id: strconv.Itoa(rand.Int()),
	}
}
