package main

import "fmt"

func main() {
	normalHouseBuilder := NewNormalHouseBuilder()
	iglooHouseBuilder := NewIglooHouseBuilder()

	director := NewDirector(normalHouseBuilder)
	normalHouse := director.BuildHouse()

	fmt.Println("House door type: ", normalHouse.doorType)
	fmt.Println("House window type: ", normalHouse.windowType)
	fmt.Println("House number of floors: ", normalHouse.numFloor)

	director.SetBuilder(iglooHouseBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Println("House door type: ", iglooHouse.doorType)
	fmt.Println("House window type: ", iglooHouse.windowType)
	fmt.Println("House number of floors: ", iglooHouse.numFloor)
}
