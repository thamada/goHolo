//Time-stamp: <2017-01-25 13:23:08 hamada>
package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"mypkg1"
	mpic "mypkg1/pic"
)

func main() {

	mpic.DEBUG = false
	pic.Show(mpic.Pic)

	if false {
		mypkg1.Main()
		fmt.Printf("%v, %T\n", mypkg1.Main, mypkg1.Main)
	}

}
