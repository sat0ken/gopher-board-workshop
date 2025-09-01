package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var (
	neo  machine.Pin = machine.D2
	leds [54]color.RGBA
)

func main() {
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws := ws2812.NewWS2812(neo)

	colors := []color.RGBA{
		{R: 255, G: 0, B: 0},   // 赤
		{R: 0, G: 255, B: 0},   // 緑
		{R: 0, G: 0, B: 255},   // 青
		{R: 255, G: 255, B: 0}, // 黄色
		{R: 0, G: 255, B: 255}, // シアン
		{R: 255, G: 0, B: 255}, // マゼンタ
	}

	offset := 0 // 色のオフセットを管理
	for {
		for i := range leds {
			// 各LEDに対してオフセットを適用して色を設定
			leds[i] = colors[(i+offset)%len(colors)]
		}
		ws.WriteColors(leds[:]) // LEDストリップに色を送信

		offset = (offset + 1) % len(colors) // オフセットをシフト
		time.Sleep(100 * time.Millisecond) // 遅延を挿入
	}
}

