//Time-stamp: <2017-01-29 03:08:01 hamada>
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
		/*
			go func(i int) {
				time.Sleep(1500 * time.Millisecond)
				id := i
				log.Printf("%03d\n", id)
				//			log.Println("finish", id)
				wg.Done()
			}(i)
		*/
		go func() {
			time.Sleep(1500 * time.Millisecond)
			log.Printf("%03d\n", i)
			//			log.Println("finish", id)
			wg.Done()
		}()

	}
	wg.Wait()
}
