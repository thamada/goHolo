//Time-stamp: <2017-01-27 01:14:38 hamada>
package tut

import (
	"fmt"
	"time"
)

func Default_selection() {

	// Becareful, it "LEAKS" !!
	// Tick is a convenience wrapper for
	// NewTicker providing access to the ticking
	// channel only. While Tick is useful for clients
	// that have no need to shut down the Ticker,
	// be aware that without a way to shut it down the underlying
	// Ticker cannot be recovered by the garbage collector;
	// it "leaks".
	// Unlike NewTicker, Tick will return nil if d <= 0.
	//
	//	func Tick(d Duration) <-chan Time {
	//		if d <= 0 {
	//			return nil
	//		}
	//		return NewTicker(d).C
	//	}
	//
	// https://golang.org/pkg/time/

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Printf("%v, %T: ", <-tick, <-tick)
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
