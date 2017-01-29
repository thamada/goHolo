//Time-stamp: <2017-01-30 00:45:06 hamada>
package work

import (
	log "fmt"
	"sync"
	"time"
)

func Sync() {

	/* A WaitGroup waits for a collection of goroutines to finish. The
	main goroutine calls Add to set the number of goroutines to wait
	for. Then each of the goroutines runs and calls Done when
	finished. At the same time, Wait can be used to block until all
	goroutines have finished.  */
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {

		// You must call Add here at the main goroutine
		// to set the number of
		// goroutines to wait for.
		// Don't increment WaitGroup inside a collection of goroutines
		wg.Add(1)

		// In this case,
		// 'i' on shared memory needs to be locked.
		// so without lock, the race condition happens for 'i'.
		/*
			go func() {
				time.Sleep(1500 * time.Millisecond)
				log.Printf("%03d\n", i)
				wg.Done()
			}()
		*/

		// In this case,
		// 'i' will be copied from shared to local memory,
		// so it doesn't need to be locked.
		go func(i int) {
			defer wg.Done() // I recommmend to use defer for Done() ;-)
			time.Sleep(1500 * time.Millisecond)
			log.Printf("%02d\n", i)
		}(i)

	}
	// Wait() waits until WaitGroup value will become 0.
	wg.Wait()
}
