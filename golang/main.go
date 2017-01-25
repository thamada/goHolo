//Time-stamp: <2017-01-25 14:31:08 hamada>
package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"mypkg1"
	// qualified identifiers
	//  see:  https://golang.org/ref/spec#Import_declarations
	mpic "mypkg1/pic"
	"mypkg1/exer"

)

func main() {

	exer.Stringer()

	if false {
		exer.Maps()
		mpic.DEBUG = false
		pic.Show(mpic.Pic)
		mypkg1.Main()
		fmt.Printf("%v, %T\n", mypkg1.Main, mypkg1.Main)
	}

}
