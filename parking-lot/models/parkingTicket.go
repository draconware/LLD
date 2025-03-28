package models

import (
	"math/rand"
	"strconv"
	"time"
)

type ParkingTicket struct {
	ticketNumber  string
	entryTime     time.Time
	exitTime      time.Time
	amount        float64
	paymentStatus PaymentStatus
	vehicle       IVehicle
	payment       IPayment
	entrance      *EntranceGate
	exit          *ExitGate
}

func newParkingTicket(vehicle IVehicle, entranceGate *EntranceGate) *ParkingTicket {
	ticketNumber := rand.Int()
	ticketNumberStr := strconv.Itoa(ticketNumber)
	return &ParkingTicket{
		ticketNumber:  ticketNumberStr,
		entryTime:     time.Now(),
		vehicle:       vehicle,
		entrance:      entranceGate,
		paymentStatus: PaymentUnPaid,
	}
}
func (pt *ParkingTicket) GetTicketNumber() string {
	return pt.ticketNumber
}
func (pt *ParkingTicket) GetEntryTime() time.Time {
	return pt.entryTime
}
func (pt *ParkingTicket) SetPayment(payment IPayment) {
	pt.payment = payment
}
func (pt *ParkingTicket) SetExit(exit *ExitGate) {
	pt.exit = exit
}
func (pt *ParkingTicket) GetExit() *ExitGate {
	return pt.exit
}
func (pt *ParkingTicket) SetExitTime(time time.Time) {
	pt.exitTime = time
}
func (pt *ParkingTicket) GetExitTime() time.Time {
	return pt.exitTime
}
func (pt *ParkingTicket) SetAmount(amount float64) {
	pt.amount = amount
}
func (pt *ParkingTicket) GetAmount() float64 {
	return pt.amount
}
func (pt *ParkingTicket) SetPaymentStatus(status PaymentStatus) {
	pt.paymentStatus = status
}
func (pt *ParkingTicket) GetPaymentStatus() PaymentStatus {
	return pt.paymentStatus
}

type IParkingRate interface {
	calculate(float64) float64
}

type ParkingRate struct {
	hours float64
	rate  float64
}

func (pr *ParkingRate) Calculate(totalHours float64) float64 {
	return (totalHours / pr.hours) * pr.rate
}

func NewParkingRate() *ParkingRate {
	return &ParkingRate{
		hours: 1.0,
		rate:  50.0,
	}
}

type ParkVehicleResponse struct {
	TicketNumber          string
	EntryTime             time.Time
	VehicleLicenseNumber  string
	ParkingChargesPerHour float64
}

func NewParkingTicketResponse(parkingTicket *ParkingTicket, parkingRate *ParkingRate) *ParkVehicleResponse {
	return &ParkVehicleResponse{
		TicketNumber:          parkingTicket.ticketNumber,
		EntryTime:             parkingTicket.entryTime,
		VehicleLicenseNumber:  parkingTicket.vehicle.GetLicenseNumber(),
		ParkingChargesPerHour: parkingRate.rate / parkingRate.hours,
	}
}

type ExitParkingResponse struct {
	TicketNumber          string
	VehicleLicenseNumber  string
	TotalAmount           float64
	PaymentStatus         string
	EntryTime             time.Time
	ExitTime              time.Time
	ParkingChargesPerHour float64
}

func NewExitParkingResponse(parkingTicket *ParkingTicket, parkingRate *ParkingRate) *ExitParkingResponse {
	return &ExitParkingResponse{
		TicketNumber:          parkingTicket.GetTicketNumber(),
		EntryTime:             parkingTicket.GetEntryTime(),
		ExitTime:              parkingTicket.GetExitTime(),
		TotalAmount:           parkingTicket.GetAmount(),
		PaymentStatus:         string(parkingTicket.GetPaymentStatus()),
		VehicleLicenseNumber:  parkingTicket.vehicle.GetLicenseNumber(),
		ParkingChargesPerHour: parkingRate.rate / parkingRate.hours,
	}
}
