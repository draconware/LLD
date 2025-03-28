package main

import "fmt"

type Train interface {
	arrive()
	depart()
	permitArrival()
}

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}

type StationMaster struct {
	isPlatformFree bool
	trainQueue     []Train
}

func newStationMaster() *StationMaster {
	return &StationMaster{
		isPlatformFree: true,
	}
}

func (sm *StationMaster) canArrive(t Train) bool {
	if sm.isPlatformFree {
		sm.isPlatformFree = false
		return true
	}
	sm.trainQueue = append(sm.trainQueue, t)
	return false
}

func (sm *StationMaster) notifyAboutDeparture() {
	if !sm.isPlatformFree {
		sm.isPlatformFree = true
	}
	if len(sm.trainQueue) > 0 {
		train := sm.trainQueue[0]
		sm.trainQueue = sm.trainQueue[1:]
		train.permitArrival()
	}
}

type PassengerTrain struct {
	mediator Mediator
}

func (pt *PassengerTrain) arrive() {
	if !pt.mediator.canArrive(pt) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting!")
		return
	}
	fmt.Println("PassengerTrain arrived.")
}

func (pt *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Departed!")
	pt.mediator.notifyAboutDeparture()
}

func (pt *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted!")
	pt.arrive()
}

type FreightTrain struct {
	mediator Mediator
}

func (ft *FreightTrain) arrive() {
	if !ft.mediator.canArrive(ft) {
		fmt.Println("FreightTrain: Arrival blocked, waiting!")
		return
	}
	fmt.Println("FreightTrain arrived.")
}

func (ft *FreightTrain) depart() {
	fmt.Println("FreightTrain: Arrival permitted!")
	ft.arrive()
}

func (ft *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted!")
	ft.arrive()
}
