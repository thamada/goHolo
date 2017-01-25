//Time-stamp: <2017-01-25 14:39:19 hamada>
package exer

import (
	// qualified identifiers
	//  see:  https://golang.org/ref/spec#Import_declarations
	mpic "mypkg1/pic"
	"fmt"
	"golang.org/x/tour/pic"
)

var DEBUG bool = false

func Pic(dx, dy int) [][]uint8 {

	p := make([][]uint8, 0)

	for iy := 0; iy < dy; iy++ {
		pline := make([]uint8, dx)
		p = append(p, pline)
	}

	for iy := 0; iy < dy; iy++ {
		pline := make([]uint8, dx)
		for ix := 0; ix < dx; ix++ {
			v := uint8(0xff & ((ix ^ iy) | (ix - 1)))
			pline[ix] = v
		}
		p[iy] = pline
	}

	if DEBUG {
		fmt.Printf("(%v, %T)\n", p, p)
		fmt.Printf("(%v, %T)\n", dx, dx)
		fmt.Printf("(%v, %T)\n", dy, dy)
	}

	return p
}

func Slices() {
	if false {
		mpic.DEBUG = false
		pic.Show(mpic.Pic)
	}else{
		DEBUG = false
		pic.Show(Pic)
	}
}
