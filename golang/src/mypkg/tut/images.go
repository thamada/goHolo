//Time-stamp: <2017-01-26 02:14:14 hamada>
package tut

import (
	"fmt"
	"image"
)

func Images() {
	// https://golang.org/pkg/image/
	m := image.NewRGBA(image.Rect(0, 0, 10, 10))
	fmt.Printf("%v, %T\n", m, m)
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
	fmt.Println(m.RGBAAt(0, 0))
	fmt.Printf("%v, %T\n", m.ColorModel(), m.ColorModel())
}
