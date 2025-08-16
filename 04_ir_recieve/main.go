package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/irremote"
)

var (
	pinLED  = machine.D27
	pinIRIn = machine.D15
	ir      irremote.ReceiverDevice
	ch      chan irremote.Data
)

func setupPins() {
	pinLED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ir = irremote.NewReceiver(pinIRIn)
	ir.Configure()
}

func irCallback(data irremote.Data) {
	ch <- data
}

func blinkLED() {
	pinLED.High()
	time.Sleep(time.Millisecond * 100)
	pinLED.Low()
}

func main() {
	setupPins()
	ch = make(chan irremote.Data, 100)
	ir.SetCommandHandler(irCallback)
	for {
		// Read data from the channel
		data := <-ch
		// Blink the LED as feedback that we received something
		blinkLED()
		strRepeat := ""
		if data.Flags&irremote.DataFlagIsRepeat != 0 {
			strRepeat = "RPT"
		}
		strTop := fmt.Sprintf("Code: %08X", data.Code)
		println(strTop)
		strBot := fmt.Sprintf("A: %04X C:%02X %s", data.Address, data.Command, strRepeat)
		println(strBot)
	}
}
