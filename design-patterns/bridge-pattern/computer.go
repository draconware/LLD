package main

import "fmt"

type Computer interface {
	SetPrinter(Printer)
	Print()
}

type MacPrinter struct {
	printer Printer
}

func (m *MacPrinter) SetPrinter(printer Printer) {
	fmt.Printf("%s printer is attached to Mac.\n", printer.GetName())
	m.printer = printer
}

func (m *MacPrinter) Print() {
	fmt.Println("Print request recieved by Mac.")
	m.printer.PrintFile()
}

type WindowsPrinter struct {
	printer Printer
}

func (w *WindowsPrinter) SetPrinter(printer Printer) {
	fmt.Printf("%s printer is attached to Windows.\n", printer.GetName())
	w.printer = printer
}

func (w *WindowsPrinter) Print() {
	fmt.Println("Print request recieved by Windows.")
	w.printer.PrintFile()
}
