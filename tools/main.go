package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"
)

// RGB888 → RGB565 (Big Endian)
func rgbToRGB565BE(r, g, b uint32) (byte, byte) {
	r5 := uint16((r >> 11) & 0x1F)
	g6 := uint16((g >> 10) & 0x3F)
	b5 := uint16((b >> 11) & 0x1F)
	val := (r5 << 11) | (g6 << 5) | b5
	return byte(val >> 8), byte(val & 0xFF)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: convert240 input.(png|jpg) output.raw")
		return
	}

	inFile := os.Args[1]
	outFile := os.Args[2]

	// 入力画像を開く
	f, err := os.Open(inFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var img image.Image
	switch filepath.Ext(inFile) {
	case ".png":
		img, err = png.Decode(f)
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(f)
	default:
		panic("unsupported format")
	}
	if err != nil {
		panic(err)
	}

	// 240x240 にリサイズ
	width, height := 240, 240
	resized := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.ApproxBiLinear.Scale(resized, resized.Bounds(), img, img.Bounds(), draw.Over, nil)

	// 出力ファイル
	out, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// RGB565BE に変換して書き込み
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := resized.At(x, y).RGBA()
			hi, lo := rgbToRGB565BE(r, g, b)
			out.Write([]byte{hi, lo})
		}
	}

	fmt.Printf("done: %s (%dx%d -> %d bytes)\n", outFile, width, height, width*height*2)
}
