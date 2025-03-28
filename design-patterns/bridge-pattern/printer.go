package main

import "fmt"

type Printer interface {
	PrintFile()
	GetName() string
}

type Epson struct {
	name string
}

func (e *Epson) GetName() string {
	return e.name
}

func (e *Epson) PrintFile() {
	fmt.Println("Printing file from Epson printer.")
}

type HP struct {
	name string
}

func (hp *HP) GetName() string {
	return hp.name
}

func (hp *HP) PrintFile() {
	fmt.Println("Printing file from HP printer.")
}
