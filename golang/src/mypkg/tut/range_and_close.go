//Time-stamp: <2017-01-26 05:42:06 hamada>
package tut

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < 0xffffff; i++ {
		c <- x
		//		x, y = y, x+y
		x, y = y, y+1
	}
	close(c)
}

func Range_and_close() {
	c := make(chan int, 1)
	go fibonacci(cap(c), c)

	sum := int(0)
	for v := range c {
		sum = sum + v
//		fmt.Printf("%v, %v, %v, %v\n", v, len(c),cap(c), c)
	}
	fmt.Println(sum)

}

