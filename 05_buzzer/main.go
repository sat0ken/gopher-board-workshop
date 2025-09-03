package main

import (
	"time"

	"machine"
)

func main() {
	buzzer := machine.D2
	buzzer.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// ドレミファソラシドの周波数（Hz）
	notes := []uint16{
		262, // ド (C4)
		294, // レ (D4)
		330, // ミ (E4)
		349, // ファ (F4)
		392, // ソ (G4)
		440, // ラ (A4)
		494, // シ (B4)
		523, // ド (C5)
	}

	for _, n := range notes {
		playTone(buzzer, n, 300*time.Millisecond)
		time.Sleep(100 * time.Millisecond) // 音の間隔
	}
}

// 指定した周波数と時間で音を鳴らす
func playTone(pin machine.Pin, freq uint16, duration time.Duration) {
	period := time.Second / time.Duration(freq) // 1周期の時間
	cycles := int(duration / period / 2)        // 高低の半周期を繰り返す回数

	for i := 0; i < cycles; i++ {
		pin.High()
		time.Sleep(period / 2)
		pin.Low()
		time.Sleep(period / 2)
	}
}
