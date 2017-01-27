//Time-stamp: <2017-01-28 00:15:56 hamada>
package tut

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	defer c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	fmt.Println(c.Value("somekey"))
	defer c.mux.Unlock()
	return c.v[key]
}

func Mutex_counter() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 10; i++ {
		go c.Value("somekey")
		go c.Inc("somekey")
	}

	fmt.Println(c.Value("somekey"))
	time.Sleep(time.Second)
}

