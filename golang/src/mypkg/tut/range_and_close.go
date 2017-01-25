//Time-stamp: <2017-01-26 05:56:03 hamada>
package tut

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	ni := n * 5
	for i := 0; i < ni; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func Range_and_close() {
	c := make(chan int, 5)
	go fibonacci(cap(c), c)

	for v := range c {
		fmt.Printf("%v, %v, %v, %v\n", v, len(c), cap(c), c)
	}

}
