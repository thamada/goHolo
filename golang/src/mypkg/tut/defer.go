//Time-stamp: <2017-01-27 03:24:49 hamada>
package tut

import (
	"fmt"
)

func sub_201701270317_b(s string) string {
	fmt.Println("sub_b:", s)
	return s
}

func sub_201701270317_a(x int) string {
	s := fmt.Sprintf("hello %v", x)
	defer fmt.Println("sub_a: Defer!")
	return sub_201701270317_b(s)
}

func Defer() {

	for i := 0; i < 10; i++ {
		fmt.Println("main:", sub_201701270317_a(i))
		fmt.Println()
	}

	fmt.Println("end.")
}
