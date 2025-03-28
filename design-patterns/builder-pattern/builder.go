package main

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() *House
}

type NormalHouseBuilder struct {
	windowType string
	doorType   string
	numFloor   int
}

func NewNormalHouseBuilder() *NormalHouseBuilder {
	return &NormalHouseBuilder{}
}

func (nhb *NormalHouseBuilder) setWindowType() {
	nhb.windowType = "Wooden Window"
}
func (nhb *NormalHouseBuilder) setDoorType() {
	nhb.doorType = "Wooden Door"
}
func (nhb *NormalHouseBuilder) setNumFloor() {
	nhb.numFloor = 2
}
func (nhb *NormalHouseBuilder) getHouse() *House {
	return &House{
		windowType: nhb.windowType,
		doorType:   nhb.doorType,
		numFloor:   nhb.numFloor,
	}
}

type IglooHouseBuilder struct {
	windowType string
	doorType   string
	numFloor   int
}

func NewIglooHouseBuilder() *IglooHouseBuilder {
	return &IglooHouseBuilder{}
}

func (ihb *IglooHouseBuilder) setWindowType() {
	ihb.windowType = "Snow Window"
}
func (ihb *IglooHouseBuilder) setDoorType() {
	ihb.doorType = "Snow Door"
}
func (ihb *IglooHouseBuilder) setNumFloor() {
	ihb.numFloor = 1
}
func (ihb *IglooHouseBuilder) getHouse() *House {
	return &House{
		windowType: ihb.windowType,
		doorType:   ihb.doorType,
		numFloor:   ihb.numFloor,
	}
}
