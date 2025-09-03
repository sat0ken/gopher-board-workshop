package main

import (
	_ "embed"
	"machine"

	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/st7789"
)

//go:embed output.raw
var imgData []byte

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 16000000,
		Mode:      0,
	})

	display := st7789.New(machine.SPI1,
		machine.GPIO9,  // RESET
		machine.GPIO12, // DC
		machine.GPIO13, // CS
		machine.GPIO14) // LITE

	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_90,
		Height:   240,
		Width:    240,
	})

	// display.FillScreen(color.RGBA{0, 0, 0, 255})

	// 生バイトから Image を作成
	img := pixel.NewImageFromBytes[pixel.RGB565BE](240, 240, imgData)
	display.DrawBitmap(0, 0, img)

	for {
	}
}
