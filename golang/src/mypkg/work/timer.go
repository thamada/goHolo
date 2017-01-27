//Time-stamp: <2017-01-28 04:24:14 hamada>
package work

import "fmt"

import "time"

const PARALLEL = true

func Timer() {
	ni := 100
	c := make(chan int, ni)

	for i := 0; i < ni; i++ {
		t := 200 * time.Millisecond

		fn := func(id int, d time.Duration, ch chan int) {
			time.Sleep(d)
			if false {
				ch <- i
			}else{
				ch <- id
			}
		}

		if PARALLEL {
			go fn(i, t, c)
		} else {
			fn(i, t, c)
		}

		fmt.Printf("done %v\n", i)
	}
	fmt.Printf("done for\n")

	for i := 0; i < ni; i++ {
		fmt.Printf("%v, ", <-c)
	}

}
