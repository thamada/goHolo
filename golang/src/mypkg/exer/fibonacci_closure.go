//Time-stamp: <2017-01-25 17:59:27 hamada>
package exer

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	ncall := 0
	prev := 0
	curr := 1

	f := func() int {
		ret := 0
		if ncall == 0 {
			ret = prev
		} else if ncall == 1 {
			ret = curr
		} else {
			next := curr + prev
			prev = curr
			curr = next
			ret = next
		}
		ncall++
		return ret
	}

	return f
}

func Fibonacci_closure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
