//Time-stamp: <2017-01-29 01:03:23 hamada>
package work

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const PARALLEL = true
const INSERT_BUG = false //true

func Timer() {
	t := 2000
	ni := 100
	c := make(chan int, ni)

	for i := 0; i < ni; i++ {

		calc := func(id int, t int, ch chan int) {
			{
				// (pc uintptr, file string, line int, ok bool)
				pc, file, line, ok := runtime.Caller(0)
				v := runtime.GOROOT()
				if false {
					fmt.Println(id, pc, file, line, ok, v)
				}
			}
			// d Time.Duration
			d := time.Millisecond * time.Duration(rand.Intn(t))
			time.Sleep(d)
			if INSERT_BUG {
				// Be carefull ;-)
				ch <- i // 'i' lives after this loop ends
			} else {
				ch <- id
			}
			runtime.GC()
		}

		fmt.Printf("invoke %v\n", i)

		if PARALLEL {
			go calc(i, t, c)
		} else {
			calc(i, t, c)
		}

		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("done for\n")

	for i := 0; i < ni; i++ {
		fmt.Printf("%v: %v\n", runtime.NumGoroutine(), <-c)
		outStackProf()
	}

	fmt.Println("")
}

func outStackProf() {
	buf := make([]byte, 1<<20)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Println("-------------")
	fmt.Println(string(buf))
	fmt.Println("-------------")
}
