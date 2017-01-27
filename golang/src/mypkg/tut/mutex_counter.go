//Time-stamp: <2017-01-28 00:36:19 hamada>
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
	fmt.Println(i, ": incr: ", c.v[key])
	c.seq = append(c.seq, i)
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
	c.seq = make([]int, 0)

	for i := 0; i < 10; i++ {
		go c.Write("mutex", i+0)
		go c.Write("mutex", i+100)
		go c.Read("mutex", i)
		go c.Write("mutex", i+200)
		go c.Write("mutex", i+300)
	}

	fmt.Println(c.Read("mutex", -1))
	time.Sleep(time.Second)
	fmt.Println(c.Read("mutex", -2))
	fmt.Printf("seq: %v\n", c.seq)
}

