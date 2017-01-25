//Time-stamp: <2017-01-25 14:24:42 hamada>
package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"mypkg1"
	// qualified identifiers
	//  see:  https://golang.org/ref/spec#Import_declarations
	test "mypkg1/exer"
	mpic "mypkg1/pic"
)

func main() {

	test.Maps()

	if false {
		mpic.DEBUG = false
		pic.Show(mpic.Pic)
		mypkg1.Main()
		fmt.Printf("%v, %T\n", mypkg1.Main, mypkg1.Main)
	}

}
