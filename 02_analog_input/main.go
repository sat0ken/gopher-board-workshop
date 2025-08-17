package main

import (
	"machine"
	"time"
)

func main() {
	machine.InitADC()
	inputPin := machine.ADC{Pin: machine.D26}
	inputPin.Configure(machine.ADCConfig{})

	for {
		voltage := float32(inputPin.Get()) * (3.3 / 65535.0)
		println(voltage)
		time.Sleep(time.Millisecond * 1000)
	}
}
