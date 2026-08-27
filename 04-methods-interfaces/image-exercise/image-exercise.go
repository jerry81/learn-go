package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 250, 1000)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8((x + y) / 2), uint8((x + y) / 2), uint8((x + y) / 2), uint8((x + y) / 2)}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
