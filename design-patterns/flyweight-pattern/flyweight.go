package main

import "fmt"

const (
	TerroristDressType        = "tdress"
	CounterTerroristDressType = "ctdress"
)

var flyweightFactorySingleInstance = &FlyWeightFactory{
	dressMap: make(map[string]Dress),
}

func GetFlyWeightFactorySingleInstance() *FlyWeightFactory {
	return flyweightFactorySingleInstance
}

type FlyWeightFactory struct {
	dressMap map[string]Dress
}

func (fwf *FlyWeightFactory) getDressByType(dressType string) Dress {
	if fwf != nil && fwf.dressMap[dressType] != nil {
		return fwf.dressMap[dressType]
	}
	switch dressType {
	case TerroristDressType:
		fwf.dressMap[dressType] = newTerroristDress()
		return fwf.dressMap[dressType]
	case CounterTerroristDressType:
		fwf.dressMap[dressType] = newCounterTerroristDress()
		return fwf.dressMap[dressType]
	default:
		fmt.Println("dress type not supported")
	}
	return nil
}

type Dress interface {
	getColor() string
}

type TerroristDress struct {
	color string
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "Red"}
}

func (td *TerroristDress) getColor() string {
	return td.color
}

type CounterTerroristDress struct {
	color string
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "Blue"}
}

func (ctd *CounterTerroristDress) getColor() string {
	return ctd.color
}
