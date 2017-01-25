//Time-stamp: <2017-01-25 14:36:33 hamada>
package main

import (
	"fmt"
	"mypkg1"
	"mypkg1/exer"
)

func main() {

	exer.Slices()

	if false {
		exer.Stringer()
		exer.Maps()
		mypkg1.Main()
		fmt.Printf("%v, %T\n", mypkg1.Main, mypkg1.Main)
	}

}
