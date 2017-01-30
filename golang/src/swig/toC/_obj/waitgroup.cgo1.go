// Created by cgo - DO NOT EDIT

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:2
package toC
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:5

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:4
import (
	log "fmt"
	"sync"
	"time"
)
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:13

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:12
func Waitgroup() {
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:20

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:19
	var wg sync.WaitGroup
									for i, ni := 0, 100; i < ni; i++ {
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:27

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:26
		wg.Add(1)
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:32

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:31
		if false {
			go func() {
				time.Sleep(1500 * time.Millisecond)
				log.Printf("%03d\n", i)
				wg.Done()
			}()
		}
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:43

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:42
		go func(i int) {
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:45

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:44
			defer wg.Done()
											sleep := time.Duration((ni-i)*10) * time.Millisecond
											time.Sleep(sleep)
											log.Printf("%02d: %v\n", i, sleep)
		}(i)
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:51

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:50
	}
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:53

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:52
	if false {
		wg.Add(1)
									}
//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:61

//line /Users/hamada/git/goHolo/golang/src/swig/toC/waitgroup.go:60
	wg.Wait()
}
