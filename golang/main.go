//Time-stamp: <2017-01-25 14:43:07 hamada>
package main

import (
	"fmt"
	"mypkg1"
	"mypkg1/exer"
)

func main() {

	exer.Fibonacci_closure()

	if false {
		exer.Slices()
		exer.Stringer()
		exer.Maps()
		mypkg1.Main()
		fmt.Printf("%v, %T\n", mypkg1.Main, mypkg1.Main)
	}

}
