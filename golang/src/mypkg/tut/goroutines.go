//Time-stamp: <2017-01-26 03:18:29 hamada>
package tut

import (
	"fmt"
	"time"
)

func say(s string) {
	time.Sleep(10 * time.Millisecond)
	fmt.Println(s)
}

func say5(s string) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(s)
}

func Goroutines() {
	for i := 0; i < 100; i++ {
		go say(fmt.Sprintf("%v: -X-", i))
		go say(fmt.Sprintf("%v: O", i))
	}

	time.Sleep(10000 * time.Millisecond)	
}
