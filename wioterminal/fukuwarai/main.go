// 8-2 Gopher福笑い
//
// Wio Terminal の画面に Gopher を表示して各パーツを移動することが出来ます
//
// - LCD に Gopher を表示する
// - 十字キーにより操作対象を移動する
// - BUTTON_3 および十字キー押し込みにより操作対象を切り替える
// - BUTTON_2 により Gopher の体の色を変える
// - BUTTON_1 により QR コード画面へ切り替える
//
// ビルドする時は速度面を改善するために -opt 2 オプションを追加してください
//
// tinygo flash --target wioterminal --size short --opt 2 ./wioterminal/fukuwarai
//
// 色々な見た目の Gopher を作って、写真を Tweet してください
// QR コード画面の QR コードを読ませると Twitter を開くことが出来ます
// ハッシュタグは #tinygobook および #tinygo です

package main

import (
	"embed"
	_ "embed"
	"fmt"
	"image/color"
	"time"
)

var (
	eyeOpen     = false
	colorsIndex = 0
	buf         [102 * 101 * 2]uint8
	fb          framebuffer
)

func main() {
	objects := []Obj{
		{file: "img/02_eye_left_1.bin", x: 54, y: 73, w: 101, h: 101},
		{file: "img/02_eye_left_2.bin", x: 72, y: 86, w: 25, h: 25},
		{file: "img/02_eye_right_1.bin", x: 192, y: 37, w: 102, h: 101},
		{file: "img/02_eye_right_2.bin", x: 217, y: 50, w: 25, h: 24},
		{file: "img/03_mouse.bin", x: 154, y: 88, w: 54, h: 54},
	}

	redraw(objects)

	sel := 0
	for {
		needsRedraw := true
		switch GetPressedKey() {
		case KeyRight:
			objects[sel].Move(1, 0)
		case KeyLeft:
			objects[sel].Move(-1, 0)
		case KeyDown:
			objects[sel].Move(0, 1)
		case KeyUp:
			objects[sel].Move(0, -1)
		case KeyReturn:
			sel = (sel + 1) % len(objects)
			time.Sleep(200 * time.Millisecond)
		case KeySpace:
			changeColor()
			time.Sleep(200 * time.Millisecond)
		case KeyBackspace:
			license()
			time.Sleep(500 * time.Millisecond)
			for GetPressedKey() == 0xFFFF {
			}
			redraw(objects)
			time.Sleep(100 * time.Millisecond)
		default:
			needsRedraw = false
		}

		if needsRedraw {
			redraw(objects)

			for _, o := range objects {
				fmt.Printf("(%d, %d) ", o.x, o.y)
			}
			fmt.Printf("\r\n")
			time.Sleep(10 * time.Millisecond)
		}
	}
}

var colors = [][2]uint16{
	{0xAE3A, 0xCEFC}, // 0xABC3D6 0xCDDBE6 TinyGo
	// https://go.dev/assets/go-brand-book-v1.9.5.pdf
	{0x6E7C, 0x9EFD}, // 0x68CCE7 0x9CDBED Gopher Blue
	{0xC77E, 0xE7BF}, // 0xC5E9F2 0xE3F4F8 Light Blue
	{0xEDB8, 0xF6FC}, // 0xEBB1C1 0xF6DCE3 Fuchsia
	{0xAEFB, 0xD77D}, // 0xADDEDB 0xD7EEED Aqua
	{0xFFB4, 0xFFFA}, // 0xFEF1A4 0xFEF9D5 Yellow
}

func changeColor() {
	colorsIndex = (colorsIndex + 1) % len(colors)
}

func redraw(objects []Obj) {
	r, _ := static.Open("img/01_body.bin")
	r.Read(fb[:])

	for _, o := range objects {
		o.Draw()
	}
	for i := range fb {
		if i%2 == 0 {
			if fb[i] == 0xBD { // BDF7
				fb[i+0] = uint8(colors[colorsIndex][0] >> 8)
				fb[i+1] = uint8(colors[colorsIndex][0])
			} else if fb[i] == 0xDE { // DEFB
				fb[i+0] = uint8(colors[colorsIndex][1] >> 8)
				fb[i+1] = uint8(colors[colorsIndex][1])
			}
		}
	}
	display.DrawRGBBitmap8(0, 0, fb[:], 320, 240)
}

type framebuffer [320 * 240 * 2]uint8

func (f *framebuffer) SetRGB555a(x, y int, b1, b2 uint8) {
	if (b2 & 0x20) > 0 {
		f[(x+y*320)*2] = b1
		f[(x+y*320)*2+1] = b2
	}
}

type Obj struct {
	data []uint8
	file string
	x    int16
	y    int16
	w    int16
	h    int16
}

func (o Obj) Draw() {
	r, _ := static.Open(o.file)
	r.Read(buf[:])

	ws, we := 0, int(o.w)
	hs, he := 0, int(o.h)
	for y := hs; y < he; y++ {
		for x := ws; x < we; x++ {
			fb.SetRGB555a(int(o.x)+x, (int(o.y) + y), buf[(x+y*we)*2], buf[(x+y*we)*2+1])
		}
	}
}

func (o *Obj) Move(vx, vy int16) {
	if 0 <= o.x+vx && o.x+vx+o.w <= 320 {
		o.x += vx
	}

	if 0 <= o.y+vy && o.y+vy+o.h <= 180 {
		o.y += vy
	}
}

func failMessage(err error) {
	for {
		fmt.Printf("%s\r\n", err.Error())
		time.Sleep(5 * time.Second)
	}
}

// The following is a definition of a special key that goes beyond the ASCII
// range.
const (
	KeyReturn    = 0x101
	KeySpace     = 0x020
	KeyBackspace = 0x103
	KeyRight     = 0x106
	KeyLeft      = 0x107
	KeyDown      = 0x108
	KeyUp        = 0x109
)

//go:embed img/*
var static embed.FS

func license() {
	display.FillScreen(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
	r, _ := static.Open("img/qr_and_license.bin")
	r.Read(fb[:])
	display.DrawRGBBitmap8(0, 0, fb[:], 320, 240)
}
