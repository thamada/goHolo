//Time-stamp: <2017-01-28 04:07:48 hamada>
package work

import "fmt"

import "time"

func timer(id int, d time.Duration) <-chan int {
	c := make(chan int)

	if true {
		go func() {
			time.Sleep(d)
			c <- id
		}()
	} else {
		time.Sleep(d)
		c <- id
	}

	return c
}

func Timer() {

	for i := 0; i < 24; i++ {
		c := timer(i, 100 * time.Millisecond)
		fmt.Printf("[%v]: ", i)
		fmt.Printf("%v\n", <-c)

	}
}
