//Time-stamp: <2017-01-26 03:53:45 hamada>
package tut

import (
	"fmt"
	"time"
)

func say(s string, n int, c chan int) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(s)
	c <- n
}

func say5(s string, n int, c chan int) {
	time.Sleep(999 * time.Millisecond)
	fmt.Println(s)
	c <- n
}

func Goroutines() {

	c := make(chan int)

	for i := 0; i < 100; i++ {
		fmt.Println(fmt.Sprintf("--------: %v", i))
		go say(fmt.Sprintf("%v: 1.oops!!", i), 1, c)
		go say(fmt.Sprintf("%v: 2.^_^;", i), 2, c)
		go say5(fmt.Sprintf("%v: 3.five!", i), 3, c)
		if true {
			x, y, z := <-c, <-c, <-c // --- sync threads ---
			fmt.Println(x, y, z)
		}
	}

	time.Sleep(1000 * time.Millisecond)
}
