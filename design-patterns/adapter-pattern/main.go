package main

func main() {
	user1 := &User{}

	macComputer := &Mac{}
	user1.InsertLightningConnectorIntoComputer(macComputer)

	windowsComputer := &Windows{}
	windowsAdapter := &WindowsAdapter{
		windowsMachine: *windowsComputer,
	}

	user1.InsertLightningConnectorIntoComputer(windowsAdapter)
}
