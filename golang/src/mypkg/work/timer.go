//Time-stamp: <2017-01-29 01:56:01 hamada>
package work

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const PARALLEL = true
const INSERT_BUG = true

func Timer() {
	t := 2000
	ni := 10
	c := make(chan int, ni)

	for i := 0; i < ni; i++ {

		calc := func(id int, t int, ch chan int) {
			fmt.Printf("start calc id=%v, i=%v\n", id, i)
			time.Sleep(20 * time.Millisecond)
			fmt.Printf("debug calc id=%v, i=%v\n", id, i) 
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
				ch <- i // 'i' lives on shared memory after this loop ends
			} else {
				ch <- id
			}
			runtime.GC()
		}

		if PARALLEL {
			go calc(i, t, c)
		} else {
			calc(i, t, c)
		}

		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("done for\n")

	for i := 0; i < ni; i++ {
		fmt.Printf("%v: i=%v\n", runtime.NumGoroutine(), <-c)
		if false {
			outStackProf()
		}
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
