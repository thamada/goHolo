//Time-stamp: <2017-01-29 01:57:15 hamada>
package main

import (
	"fmt"
	"mypkg/etc"
	"mypkg/exer"
	"mypkg/tut"
	"mypkg/work"
)

func main() {

	work.Race_conditions()

	if false {
		work.Timer()
		tut.Numeric_constants()
		exer.Web()
		tut.Mutex_counter()
		tut.Defer()
		exer.Equivalent_binary_trees()
		tut.Default_selection()
		tut.Select()
		tut.Range_and_close()
		tut.Buffered_channels()
		tut.Channels()
		tut.Goroutines()
		exer.Images()
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
