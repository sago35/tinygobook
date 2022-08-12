package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"
)

func main() {
	for _, file := range os.Args[1:] {
		fmt.Printf("%s\n", file)
		r, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()

		img, err := png.Decode(r)
		if err != nil {
			log.Fatal(err)
		}

		to := strings.Replace(file, ".png", ".bin", -1)
		w, err := os.Create(to)
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
