//Time-stamp: <2017-01-25 11:36:18 hamada>
package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"mypkg1"
	"mypkg1/pic1"
)

func main() {

	pic.Show(pic1.Pic)

	if false {
		mypkg1.Main()
		fmt.Printf("%v, %T\n", mypkg1.Main, mypkg1.Main)
	}

}
