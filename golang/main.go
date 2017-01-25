//Time-stamp: <2017-01-26 02:16:44 hamada>
package main

import (
	"fmt"
	"mypkg/etc"
	"mypkg/exer"
	"mypkg/tut"
)

func main() {

	exer.Images()

	if false {
		tut.Images()
		exer.Rot_reader()
		exer.Reader()
		tut.Reader()
		exer.Errors()
		exer.Fibonacci_closure()
		exer.Slices()
		exer.Stringer()
		exer.Maps()
		etc.Main()
		fmt.Printf("%v, %T\n", etc.Main, etc.Main)
	}

}
