//Time-stamp: <2017-01-26 04:12:47 hamada>
package tut

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	//	fmt.Println(s)
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel c
}

func Channels() {
	s := make([]int, 0)
	z := int(0)
	ni := 100000

	for i := 0; i < ni; i++ {
		s = append(s, i)
		z += i
	}

	c := make(chan int)
	for i := 0; i < 100; i++ {
		n := ni/2
		go sum(s[:n], c)
		go sum(s[n:], c)
		x, y := <-c, <-c // sync threads
		fmt.Println(x, y, x+y, x+y-z)
	}
}
