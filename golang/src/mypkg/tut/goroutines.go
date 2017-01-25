//Time-stamp: <2017-01-26 03:44:29 hamada>
package tut

import (
	"fmt"
	"time"
)

func say(s string, c chan int) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(s)
	c <- 1
}

func say5(s string, c chan int) {
	time.Sleep(999 * time.Millisecond)
	fmt.Println(s)
	c <- 5
}

func Goroutines() {

	c := make(chan int)

	for i := 0; i < 100; i++ {
		fmt.Println(fmt.Sprintf("--------: %v", i))
		go say(fmt.Sprintf("%v: oops", i), c)
		go say(fmt.Sprintf("%v: ^_^;", i), c)
		go say5(fmt.Sprintf("%v: five!", i), c)
		x, y, z := <-c, <-c, <-c
		fmt.Println(x, y, z)
	}

	//	time.Sleep(10000 * time.Millisecond)
}
