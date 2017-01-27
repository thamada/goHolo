//Time-stamp: <2017-01-28 04:15:38 hamada>
package work

import "fmt"

import "time"

func timer(id int, d time.Duration, ch chan int) {
	time.Sleep(d)
	ch <- id
}

func Timer() {
	ni := 100 
	c := make(chan int, ni)

	for i := 0; i < ni; i++ {
		t := 2000 * time.Millisecond
		go timer(i, t, c)
		fmt.Printf("done %v\n", i)
	}
	fmt.Printf("done for\n")

	for i := 0; i < ni; i++ {
		fmt.Printf("%v, ", <-c)
	}

}
