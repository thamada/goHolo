//Time-stamp: <2017-01-28 00:30:18 hamada>
package tut

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
	seq [15]int
}

func (c *SafeCounter) Write(key string, i int) {
	c.mux.Lock()
	c.v[key]++
	fmt.Println(i, ": incr: ", c.v[key])
	c.seq[i] = c.v[key]
	defer c.mux.Unlock()
}

func (c *SafeCounter) Read(key string, i int) int {
	c.mux.Lock()
	fmt.Println(i, ": read: ", c.v[key])
	defer c.mux.Unlock()
	return c.v[key]
}

func Mutex_counter() {
	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 10; i++ {
		go c.Read("mutex", i)
		go c.Write("mutex", i)
		go c.Write("mutex", i)
	}

	fmt.Println(c.Read("mutex", -1))
	time.Sleep(time.Second)
	fmt.Println(c.Read("mutex", -2))
	fmt.Printf("seq: %v\n", c.seq)
}

