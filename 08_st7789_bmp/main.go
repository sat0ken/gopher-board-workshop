package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/st7789"
)

func main() {
	// SPI 設定
	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 16000000,
		Mode:      0,
	})

	// ST7789 初期化
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

	// 画面クリア（黒）
	display.FillScreen(color.RGBA{0, 0, 0, 255})

	width, height := 240, 240
	img := pixel.NewImage[pixel.RGB565BE](width, height)

	// グラデーション描画
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// color.RGBA から pixel カラー型へ変換
			c := pixel.NewColor[pixel.RGB565BE](uint8(x), uint8(y), uint8((x+y)/2))
			img.Set(x, y, c)
		}
	}

	display.DrawBitmap(0, 0, img)

	for {
	}
}
