//Time-stamp: <2017-01-27 03:31:18 hamada>
package tut

import (
	"fmt"
)


func Defer() {

	var func_a func (x int) string
	var func_b func (s string) string 

	func_b = 	func (s string) string {
		fmt.Println("sub_b:", s)
		return s
	}

	func_a = func (x int) string {
		s := fmt.Sprintf("hello %v", x)
		defer fmt.Println("sub_a: Defer!")
		return func_b(s)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("main:", func_a(i))
		fmt.Println()
	}

	fmt.Println("end.")
}
