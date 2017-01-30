// Created by cgo - DO NOT EDIT

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:2
package toC
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:8

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:7
import (
	log "fmt"
	"sync"
	"time"
)
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:14

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:13
func Waitgroup() {
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:21

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:20
	var wg sync.WaitGroup
									for i, ni := 0, 100; i < ni; i++ {
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:28

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:27
		wg.Add(1)
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:33

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:32
		if false {
			go func() {
				time.Sleep(1500 * time.Millisecond)
				log.Printf("%03d\n", i)
				wg.Done()
			}()
		}
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:44

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:43
		go func(i int) {
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:46

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:45
			defer wg.Done()
											sleep := time.Duration((ni-i)*10) * time.Millisecond
											time.Sleep(sleep)
											log.Printf("%02d: %v\n", i, sleep)
		}(i)
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:52

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:51
	}
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:54

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:53
	if false {
		wg.Add(1)
									}
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:62

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:61
	wg.Wait()
}
