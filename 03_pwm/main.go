package main

import (
	"machine"
	"time"
)

var period uint64 = 2e9

func main() {
	pin := machine.D26
	pwm := machine.PWM5

	pwm.Configure(machine.PWMConfig{})

	ch, err := pwm.Channel(pin)
	if err != nil {
		println(err.Error())
		return
	}

	for {
		// Fade the LED in.
		for percentOn := 0; percentOn <= 100; percentOn++ {
			pwm.Set(ch, pwm.Top()*uint32(percentOn)/100)
			time.Sleep(time.Second / 100)
		}

		// Fade the LED out.
		for percentOn := 100; percentOn >= 0; percentOn-- {
			pwm.Set(ch, pwm.Top()*uint32(percentOn)/100)
			time.Sleep(time.Second / 100)
		}
	}
}
