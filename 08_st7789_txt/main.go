package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/st7789"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 16000000,
		Mode:      0,
	})

	display := st7789.New(machine.SPI1,
		machine.GPIO9,  // TFT_RESET
		machine.GPIO12, // TFT_DC
		machine.GPIO13, // TFT_CS
		machine.GPIO14) // TFT_LITE

	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_90,
		Height:   240,
		Width:    240,
	})

	// Clear the screen to black
	display.FillScreen(color.RGBA{0, 0, 0, 255})

	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 00, 50, "Hello", color.RGBA{R: 255, G: 255, B: 0, A: 255})
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 00, 80, "Gophers!", color.RGBA{R: 255, G: 0, B: 255, A: 255})
}
