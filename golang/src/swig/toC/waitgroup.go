//Time-stamp: <2017-01-31 01:03:41 hamada>
package toC

import "C"
//export Waitgroup

import (
	log "fmt"
	"sync"
	"time"
)

func Waitgroup() {

	/* A WaitGroup waits for a collection of goroutines to finish. The
	main goroutine calls Add to set the number of goroutines to wait
	for. Then each of the goroutines runs and calls Done when
	finished. At the same time, Wait can be used to block until all
	goroutines have finished.  */
	var wg sync.WaitGroup
	for i, ni := 0, 100; i < ni; i++ {

		// You must call Add here at the main goroutine
		// to set the number of
		// goroutines to wait for.
		// Don't increment WaitGroup inside a collection of goroutines
		wg.Add(1)

		// In this case,
		// 'i' on shared memory needs to be locked.
		// so without lock, the race condition happens for 'i'.
		if false {
			go func() {
				time.Sleep(1500 * time.Millisecond)
				log.Printf("%03d\n", i)
				wg.Done()
			}()
		}

		// In this case,
		// 'i' will be copied from shared to local memory,
		// so it doesn't need to be locked.
		go func(i int) {
			// Done decrements the WaitGroup counter.
			defer wg.Done() // I recommmend to use defer for Done() ;-)
			sleep := time.Duration((ni-i)*10) * time.Millisecond
			time.Sleep(sleep)
			log.Printf("%02d: %v\n", i, sleep)
		}(i)

	}

	if false {
		wg.Add(1) // deadlock happens
	}

	// WaitGroup don't have a method to know the WaitGroup counter.
	// log.Printf(*wg) // Error: invalid indirect

	// Wait() waits until WaitGroup value will become 0.
	wg.Wait()
}
