//Time-stamp: <2017-01-28 04:20:44 hamada>
package work

import "fmt"

import "time"

func Timer() {
	ni := 100
	c := make(chan int, ni)

	for i := 0; i < ni; i++ {
		t := 2000 * time.Millisecond

		go func(id int, d time.Duration, ch chan int) {
			time.Sleep(d)
			ch <- id
		}(i, t, c)

		fmt.Printf("done %v\n", i)
	}
	fmt.Printf("done for\n")

	for i := 0; i < ni; i++ {
		fmt.Printf("%v, ", <-c)
	}

}
