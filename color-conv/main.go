package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	files := []struct{ src, dst string }{
		{
			src: "./wioterminal/fukuwarai/png/01_body.png",
			dst: "./wioterminal/fukuwarai/img/01_body.bin",
		},
		{
			src: "./wioterminal/fukuwarai/png/02_eye_left_1.png",
			dst: "./wioterminal/fukuwarai/img/02_eye_left_1.bin",
		},
		{
			src: "./wioterminal/fukuwarai/png/02_eye_left_2.png",
			dst: "./wioterminal/fukuwarai/img/02_eye_left_2.bin",
		},
		{
			src: "./wioterminal/fukuwarai/png/02_eye_right_1.png",
			dst: "./wioterminal/fukuwarai/img/02_eye_right_1.bin",
		},
		{
			src: "./wioterminal/fukuwarai/png/02_eye_right_2.png",
			dst: "./wioterminal/fukuwarai/img/02_eye_right_2.bin",
		},
		{
			src: "./wioterminal/fukuwarai/png/03_mouse.png",
			dst: "./wioterminal/fukuwarai/img/03_mouse.bin",
		},
		{
			src: "./wioterminal/fukuwarai/png/qr_and_license.png",
			dst: "./wioterminal/fukuwarai/img/qr_and_license.bin",
		},
	}
	for _, file := range files {
		fmt.Printf("%s\n", file.src)
		r, err := os.Open(file.src)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()

		img, err := png.Decode(r)
		if err != nil {
			log.Fatal(err)
		}

		w, err := os.Create(file.dst)
		if err != nil {
			log.Fatal(err)
		}
		defer w.Close()

		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				b := COLORTo555a(img.At(x, y))
				w.Write([]byte{byte(b >> 8), byte(b)})
			}
		}
	}
}

// r:5bit + g:5bit + a:1bit + b:5bit
func COLORTo555a(c color.Color) uint16 {
	r, g, b, a := c.RGBA()
	if a > 0 {
		a = 0x0020
	}
	out := uint16((r & 0xF800) +
		((g & 0xF800) >> 5) +
		((b & 0xF800) >> 11) +
		a)

	return out
}
