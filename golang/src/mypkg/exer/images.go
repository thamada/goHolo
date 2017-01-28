//Time-stamp: <2017-01-26 02:43:10 hamada>
package exer

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (recv *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (recv *Image) Bounds() image.Rectangle {
	return image.Rectangle{
		image.Point{0, 0},
		image.Point{100, 100},
	}
}

func (recv *Image) At(x, y int) color.Color {
	r := ((x + y) / 2) % 256
	g := (x * y) % 256
	b := ((x + y) / 2) % 256
	a := 255
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func Images() {
	m := Image{}
	pic.ShowImage(&m)
}
