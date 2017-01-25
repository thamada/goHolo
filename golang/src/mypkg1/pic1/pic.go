//Time-stamp: <2017-01-25 12:57:07 hamada>
package pic1

import (
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {

	p := make([][]uint8, 0)

	for iy := 0; iy < dy; iy++ {
		pline := make([]uint8, dx)
		p = append(p, pline)
	}

	for iy := 0; iy < dy; iy++ {
		for ix := 0; ix < ix; ix++ {
//			p[ix][iy] = uint8(0xff & (ix + iy))
			p[ix][iy] = uint8(177)
		}
	}

	fmt.Printf("(%v, %T)\n", p, p)

	return p
}
