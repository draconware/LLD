package main

func main() {
	stationMaster := newStationMaster()

	passengerTrain := &PassengerTrain{
		mediator: stationMaster,
	}
	freightTrain := &FreightTrain{
		mediator: stationMaster,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
