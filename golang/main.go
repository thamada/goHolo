//Time-stamp: <2017-01-25 10:32:39 hamada>
package main

import (
	"fmt"
//	"golang.org/x/tour/pic"
	"mypkg1"
)

/*
func Pic(dx, dy int) [][]uint8 {
	z := [][]uint8
	return z
}
*/

func main() {

//	pic.Show(Pic)
	mypkg1.Main()
	fmt.Println("Hello!!")
	if false {
		mypkg1.Main()
	}
}
