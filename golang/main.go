//Time-stamp: <2017-01-25 18:02:19 hamada>
package main

import (
	"fmt"
	"mypkg1"
	"mypkg1/exer"
)

func main() {

	exer.Errors()

	if false {
		exer.Fibonacci_closure()
		exer.Slices()
		exer.Stringer()
		exer.Maps()
		mypkg1.Main()
		fmt.Printf("%v, %T\n", mypkg1.Main, mypkg1.Main)
	}

}
