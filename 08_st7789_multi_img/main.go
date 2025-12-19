package main

import (
	_ "embed"
	"machine"
	"time"

	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/st7789"
)

//go:embed output.raw
var imgData1 []byte

//go:embed tinygo-logo.raw
var imgData2 []byte

func main() {
	button1 := machine.D3
	button1.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

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

	img1 := pixel.NewImageFromBytes[pixel.RGB565BE](240, 240, imgData1)
	img2 := pixel.NewImageFromBytes[pixel.RGB565BE](240, 240, imgData2)

	// ▼ スライスも値型で統一
	images := []pixel.Image[pixel.RGB565BE]{img1, img2}
	current := 0

	// 最初の画像表示
	display.DrawBitmap(0, 0, images[current])

	for {
		// プルアップなので LOW = 押された
		if !button1.Get() {
			// 切り替え
			current = (current + 1) % len(images)

			display.DrawBitmap(0, 0, images[current])

			// チャタリング防止 & 長押し対策
			time.Sleep(300 * time.Millisecond)
		}

		time.Sleep(10 * time.Millisecond)
	}
}
