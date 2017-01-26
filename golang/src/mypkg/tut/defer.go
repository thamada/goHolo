//Time-stamp: <2017-01-27 03:37:13 hamada>
package tut

import (
	"fmt"
)

func Defer() {
	var func_a func (x int) string
	var func_b func (s string) string 

	func_a = func (x int) string {
		var s string
		for i :=0; i<10; i++ { 
			s = fmt.Sprintf("msg %v", x+i)
			defer fmt.Println("sub_a: Defer: ", s)
		}
		return func_b(s)
	}

	func_b = 	func (s string) string {
		fmt.Println("sub_b:", s)
		return s
	}

	for i := 0; i < 2; i++ {
		fmt.Println("main:", func_a(i))
		fmt.Println()
	}

	fmt.Println("end.")
}
