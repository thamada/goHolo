//Time-stamp: <2017-01-29 02:20:04 hamada>
package work

import (
	"fmt"
	"time"
)

func Race_conditions() {
	ni := 100

	for i := 0; i < ni; i++ {

		// calc() is a closure which has 'i' on the shared memory.
		calc := func(id int) {
			ii := i // i is on shared memory
			time.Sleep(100 * time.Nanosecond)
			var s = fmt.Sprintf("%v%v", id, i)
			if ii != i {
				fmt.Printf("id=%v, i=%v\n", id, i)
				fmt.Println(s)
				fmt.Println([]byte(s))
			}
		}

		if true {
			go calc(i)
		} else {
			calc(i)
		}
		time.Sleep(10000 * time.Nanosecond)
	}
	time.Sleep(10 * time.Millisecond)
}
