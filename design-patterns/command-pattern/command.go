package main

import "fmt"

type Device interface {
	on()
	off()
}

type Command interface {
	execute()
}

type OnCommand struct {
	device Device
}

func (on *OnCommand) execute() {
	on.device.on()
}

type OffCommand struct {
	device Device
}

func (off *OffCommand) execute() {
	off.device.off()
}

type TV struct {
	isRunning bool
}

func (tv *TV) on() {
	tv.isRunning = true
	fmt.Println("TV is on.")
}

func (tv *TV) off() {
	tv.isRunning = false
	fmt.Println("TV is off.")
}

type Button interface {
	press()
}

type OnButton struct {
	command Command
}

func (on *OnButton) press() {
	on.command.execute()
}

type OffButton struct {
	command Command
}

func (off *OffButton) press() {
	off.command.execute()
}
