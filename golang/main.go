//Time-stamp: <2017-01-25 11:28:34 hamada>
package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"mypkg1"
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

func main() {

	pic.Show(Pic)

	if false {
		mypkg1.Main()
		fmt.Printf("%v, %T\n", mypkg1.Main, mypkg1.Main)
	}

}
