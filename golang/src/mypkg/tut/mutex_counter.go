//Time-stamp: <2017-01-28 01:01:42 hamada>
package tut

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
	seq []int
}

func (c *SafeCounter) Write(key string, i int) {
	c.mux.Lock()
	c.v[key]++
	fmt.Printf("%03d: incr: %v\n", i, c.v[key])
	c.seq = append(c.seq, i)
	if false {
		defer c.mux.Unlock()
	} else {
		c.mux.Unlock()
	}
	time.Sleep(100 * time.Millisecond)
}

func (c *SafeCounter) Read(key string, i int) int {
	c.mux.Lock()
	fmt.Printf("%03d: read: %v\n", i, c.v[key])
	defer c.mux.Unlock()
	return c.v[key]
}

func Mutex_counter() {
	c := SafeCounter{v: make(map[string]int)}
	c.seq = make([]int, 0)

	tid := int(0)
	for i := 0; i < 100; i++ {
		go c.Write("mutex", tid)
		tid++

		go c.Write("mutex", tid)
		tid++

		go c.Read("mutex", i+900)

		go c.Write("mutex", tid)
		tid++

		go c.Write("mutex", tid)
		tid++
	}

	fmt.Println(c.Read("mutex", -1))
	time.Sleep(2 * time.Second)
	fmt.Println(c.Read("mutex", -2))
	fmt.Printf("seq: %v\n", c.seq)
}
