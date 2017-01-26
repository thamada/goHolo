//Time-stamp: <2017-01-27 00:24:13 hamada>
package tut

import (
	"fmt"
	"time"
)

func fibon(c, quit, exit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // --- send to c ---
			x, y = y, x+y
			time.Sleep(500 * time.Millisecond)
		case <-quit: // --- recv from quit ---
			fmt.Println("quit")
			exit <- 0
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
	exit := make(chan int)
	if true {
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println(i, <-c)
			}
			quit <- 0
		}()
	} else {
		go print(c, quit)
	}

	go fibon(c, quit, exit)

	<-exit // sync thread fibon

}
