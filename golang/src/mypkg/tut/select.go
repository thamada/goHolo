//Time-stamp: <2017-01-26 06:04:38 hamada>
package tut

import (
	"fmt"
)

func fibon(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func print(c, quit chan int) {
	for i := 0; i < 20; i++ {
		fmt.Println(i, <-c)
	}
	quit <- 0
}

func Select() {
	c := make(chan int)
	quit := make(chan int)
	if true {
		fn := func() {
			for i := 0; i < 10; i++ {
				fmt.Println(<-c)
			}
			quit <- 0
		}
		go fn()
	} else {
		go print(c, quit)
	}
	fibon(c, quit)
}
