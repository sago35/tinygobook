//go:build !baremetal
// +build !baremetal

package main

import (
	"image/color"

	"github.com/sago35/tinydisplay/examples/initdisplay"
)

var (
	display *initdisplay.TinyDisplay
)

func init() {
	display = initdisplay.InitDisplay()
	display.FillScreen(color.RGBA{0xFF, 0xFF, 0xFF, 0xFF})
}

func GetPressedKey() uint16 {
	return display.GetPressedKey()
}
