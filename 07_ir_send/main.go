package main

import (
	"machine"
	"time"
)

// NEC Format
const (
	leaderPulse time.Duration = 9000 * time.Microsecond
	leaderSpace time.Duration = 4500 * time.Microsecond
	bitPulse    time.Duration = 560 * time.Microsecond
	zeroSpace   time.Duration = 560 * time.Microsecond
	oneSpace    time.Duration = 1690 * time.Microsecond
	carrierFreq uint64        = 38000 // 38kHz
)

var (
	pwm     = machine.PWM5
	channel uint8
	duty    uint32
)

// 指定された時間、38kHzの赤外線パルスを送信する
func pulse(d time.Duration) {
	pwm.Set(channel, duty) // PWM出力をON
	time.Sleep(d)
}

// 指定された時間、信号を停止する
func space(d time.Duration) {
	pwm.Set(channel, 0) // PWM出力をOFF
	time.Sleep(d)
}

// 1バイト(8ビット)のデータを送信する
func sendByte(data uint8) {
	for i := 0; i < 8; i++ {
		pulse(bitPulse)
		if (data>>i)&1 == 1 {
			space(oneSpace) // '1'
		} else {
			space(zeroSpace) // '0'
		}
	}
}

// NECフォーマットの1フレームを送信する
func sendNECFrame(address uint8, command uint8) {
	// 1. リーダーコード
	pulse(leaderPulse)
	space(leaderSpace)

	// 2. カスタムコード (アドレス)
	sendByte(address)
	sendByte(^address) // 反転

	// 3. データコード (コマンド)
	sendByte(command)
	sendByte(^command) // 反転

	// 4. ストップビット
	pulse(bitPulse)
	space(0) // 送信終了
}

// プログラムのエントリーポイント
func main() {
	// --- ハードウェアの初期設定 ---
	pin := machine.D26 // 赤外線LEDを接続するピン

	button1 := machine.D3
	button1.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	// PWMグループを設定
	// Period = 1秒(ナノ秒) / 周波数 = 1e9 / 38000
	err := pwm.Configure(machine.PWMConfig{
		Period: 1e9 / carrierFreq,
	})
	if err != nil {
		println("failed to configure PWM:", err)
		return
	}

	// 使用するピンに対応するPWMチャンネルを取得
	channel, err = pwm.Channel(pin)
	if err != nil {
		println("failed to get PWM channel:", err)
		return
	}

	// 50%のデューティ比を計算
	duty = pwm.Top() / 2

	for {

		if !button1.Get() {
			// NECフレームを送信 (カスタムコード: 0x04, データコード: 0x0C)
			println("Sending IR signal...")
			sendNECFrame(0x04, 0x0C)
			time.Sleep(time.Millisecond * 100)
		}
	}
}
