package main

import (
	"machine"
	"time"
)

func main() {

	button1 := machine.D3
	button1.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		if !button1.Get() {
			println("button up is pressed!!")
		}

		time.Sleep(time.Millisecond * 100)
	}
}
