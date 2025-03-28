package main

import "fmt"

func main() {
	hpPrinter := &HP{
		name: "HP",
	}
	epsonPrinter := &Epson{
		name: "Epson",
	}

	mac := &MacPrinter{}
	windows := &WindowsPrinter{}

	mac.SetPrinter(hpPrinter)
	mac.Print()
	fmt.Println()

	windows.SetPrinter(epsonPrinter)
	windows.Print()
	fmt.Println()

	mac.SetPrinter(epsonPrinter)
	mac.Print()
	fmt.Println()

	mac.SetPrinter(hpPrinter)
	mac.Print()
	fmt.Println()
}
