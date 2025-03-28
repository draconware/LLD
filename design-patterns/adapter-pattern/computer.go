package main

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector plugged into Mac computer.")
}

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector plugged into Windows computer.")
}

type User struct{}

func (u *User) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("User inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

type WindowsAdapter struct {
	windowsMachine Windows
}

func (wa *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts lightning signals to USB.")
	wa.windowsMachine.insertIntoUSBPort()
}
