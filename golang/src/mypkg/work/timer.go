//Time-stamp: <2017-01-29 00:01:04 hamada>
package work

import (
	"fmt"
	"math/rand"
	"time"
)

const PARALLEL = true
const INSERT_BUG = false

func Timer() {
	t := 200
	ni := 100
	c := make(chan int, ni)

	for i := 0; i < ni; i++ {

		calc := func(id int, t int, ch chan int) {
			td := time.Duration(rand.Intn(t))
			d := time.Millisecond * td // Time.Duration
			time.Sleep(d)
			if INSERT_BUG {
				ch <- i // be carefull ;-)
			} else {
				ch <- id
			}
		}

		if PARALLEL {
			go calc(i, t, c)
		} else {
			calc(i, t, c)
		}

		time.Sleep(time.Duration(t/100) * time.Millisecond)

		fmt.Printf("done %v\n", i)
	}
	fmt.Printf("done for\n")

	for i := 0; i < ni; i++ {
		fmt.Printf("%v, ", <-c)
	}

	fmt.Println("")
}
