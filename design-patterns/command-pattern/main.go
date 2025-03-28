package main

func main() {
	tv := &TV{
		isRunning: false,
	}

	onCommand := &OnCommand{
		device: tv,
	}
	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &OnButton{
		command: onCommand,
	}
	offButton := &OffButton{
		command: offCommand,
	}

	onButton.press()
	offButton.press()
}
