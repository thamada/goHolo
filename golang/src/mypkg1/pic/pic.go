//Time-stamp: <2017-01-25 11:34:50 hamada>
package pic

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	p := make([][]uint8, 0)

	for iy := 0; iy < dy; iy++ {
		pline := make([]uint8, dx)
		for ix := 0; ix < ix; ix++ {
			pline[ix] = uint8(0xff & (ix+iy))
		}
		p = append(p, pline)
	}

	return p
}

func Pic_main() {
	pic.Show(Pic)
}
