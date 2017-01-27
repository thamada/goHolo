//Time-stamp: <2017-01-28 04:31:00 hamada>
package work

import (
	"fmt"
	"time"
)

const PARALLEL = true

func Timer() {
	ni := 100
	c := make(chan int, ni)

	for i := 0; i < ni; i++ {
		t := 200 * time.Millisecond

		calc := func(id int, d time.Duration, ch chan int) {
			time.Sleep(d)
			if true {
				ch <- i  // be carefull ;-)
			}else{
				ch <- id
			}
		}

		if PARALLEL {
			go calc(i, t, c)
		} else {
			calc(i, t, c)
		}

		time.Sleep(t/40)

		fmt.Printf("done %v\n", i)
	}
	fmt.Printf("done for\n")

	for i := 0; i < ni; i++ {
		fmt.Printf("%v, ", <-c)
	}

}
