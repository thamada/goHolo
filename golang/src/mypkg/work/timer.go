//Time-stamp: <2017-01-29 00:26:36 hamada>
package work

import (
	"fmt"
	"math/rand"
	"time"
	"runtime"
)

const PARALLEL = true
const INSERT_BUG = true

func Timer() {
	t := 3000
	ni := 100
	c := make(chan int, ni)

	for i := 0; i < ni; i++ {

		calc := func(id int, t int, ch chan int) {
			{
				// (pc uintptr, file string, line int, ok bool)
				pc, file, line, ok := runtime.Caller(0)
				fmt.Println(id, pc, file, line, ok)
			}
			td := time.Duration(rand.Intn(t))
			d := time.Millisecond * td // Time.Duration
			time.Sleep(d)
			if INSERT_BUG {
				// Be carefull ;-)
				ch <- i // 'i' lives after this loop ends
			} else {
				ch <- id
			}
		}

		if PARALLEL {
			go calc(i, t, c)
		} else {
			calc(i, t, c)
		}

		time.Sleep(10 * time.Millisecond)

		fmt.Printf("done %v\n", i)
	}
	fmt.Printf("done for\n")

	for i := 0; i < ni; i++ {
		fmt.Printf("%v, ", <-c)
	}

	fmt.Println("")
}
