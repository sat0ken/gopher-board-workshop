# gopher-board-workshop

[gopher-board](https://github.com/sat0ken/gopher-board) は Go のマスコットキャラクターの Gopher を型取ったオリジナル基板です。

![](./img/gopher-board.jpg)

![](./img/gopher-board2.jpg)

様々なデザインの基板がありますが、構成は同じです。

## 環境設定

- TinyGo のインストール

以下より TinyGo をインストールしてください  
https://tinygo.org/getting-started/install/

## 基板について

基板には以下の部品がついています。

- マイコン
- スイッチ
- フルカラーLED
- 液晶
- ブザー
- Grove端子

## ワークショップ

- ブレットボード

まずワークショップで利用するブレットボードについて説明します。
ブレットボードはパーツを差し込んで簡単に電子回路を作成して試すことができるものです。

外側の+-は縦方向に、A~EとF~Jは横方向に接続されています。

- Lチカ

LEDを29とGNDに挿します。LEDの長い方を29に短い方をGNDにして向きを間違えないように挿してください。
LEDを挿したら以下のコマンドでプログラムを書き込みます。

```
$ tinygo flash --target waveshare-rp2040-zero --size short ./00_blink/main.go
```

- デジタル入力とシリアル通信

基板には6つのスイッチがついています。押されたスイッチを読み取るデジタル入力のプログラムを書き込みます。

```
$ tinygo flash --target waveshare-rp2040-zero --size short ./01_switch/main.go
```

`tinygo monitor`を実行してUpボタンを押すとメッセージが出力されます。

```
$ tinygo monitor
Connected to /dev/ttyACM0. Press Ctrl-C to exit.
button up is pressed!!
button up is pressed!!
```

Goのprintln関数は標準出力にメッセージを出力しますが、TinyGoのprintln関数はシリアルに出力します。

- アナログ入力

- アナログ出力

PWM(Pulse Width Modulation: パルス幅変調)を利用したアナログ出力でLEDを光らせます。
プログラムを書くとLチカの時とはLEDの光り方が異なり、ホタルのよう暗くなったり明るくなったりと光ります。

```
$ tinygo flash --target waveshare-rp2040-zero --size short ./03_pwm/main.go
```

PWMによるアナログ出力は、一定の周波数で高速でHIGHとLOWを切り替えることで供給する電力を制御します。
供給される電力が高いときはLEDが明るく光り、電力が低いときはLEDが暗くなります。

https://tinygo.org/tour/pwm/fade/

- フルカラーLEDを光らす

Gopherくんの目の部分にはフルカラーLEDのWS2812がついています。これを光らせてみましょう。
WS2812自体に小さなマイコンがついていてマイコンにRGBの信号を送ると光るようになっています。
TinyGoではWS2812用のドライバがあるのでそれを利用します。

- 赤外線リモコン

テレビやエアコンなど家電で利用される赤外線リモコンの挙動をTinyGoで体験してみましょう。
リモコンは送信と受信に分かれます。

まず送信側のプログラムを書き込みます。

```
$ tinygo flash --target waveshare-rp2040-zero --size short ./04_ir_recieve/main.go
```

次に受信側のプログラムを書き込みます。

```
$ tinygo flash --target waveshare-rp2040-zero --size short ./04_ir_send/main.go
```

送信側はUpボタンを押すと赤外線LEDが光り、データを送信します。
受信側はデータを受信したらLEDが光ります。

- ブザーを鳴らす


- 液晶画面に文字を出す
- [koebiten](https://github.com/sago35/koebiten)でゲームを遊んでみる
