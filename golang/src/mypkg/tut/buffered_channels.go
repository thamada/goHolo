//Time-stamp: <2017-01-26 04:57:24 hamada>
package tut

import "fmt"

func send(n int, c chan int) {
	c <- n
}

func Buffered_channels() {
	ch := make(chan int, 6)

	for i := 0; i < 100; i++ {
		go send(1, ch)
		go send(2, ch)
		go send(3, ch)
		go send(4, ch)
		go send(5, ch)
		ch <- 6
		fmt.Println(<-ch, <-ch, <-ch, <-ch, <-ch, <-ch)
	}
}
