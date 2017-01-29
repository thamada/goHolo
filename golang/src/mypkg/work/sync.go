//Time-stamp: <2017-01-30 00:26:52 hamada>
package work

import (
	log "fmt"
	"sync"
	"time"
)

func Sync() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {

		wg.Add(1)

		// In this case,
		// 'i' on shared memory needs to be locked.
		/*
			go func() {
				time.Sleep(1500 * time.Millisecond)
				log.Printf("%03d\n", i)
				//			log.Println("finish", id)
				wg.Done()
			}()
		*/

		// In this case,
		// 'i' will be copied from shared to local memory,
		// so it doesn't need to be locked.
		go func(i int) {
			time.Sleep(1500 * time.Millisecond)
			log.Printf("%02d\n", i)
			wg.Done()
		}(i)

	}
	// Wait() waits until WaitGroup value will become 0.
	wg.Wait()
}
