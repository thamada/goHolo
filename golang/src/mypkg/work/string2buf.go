//Time-stamp: <2017-01-29 02:33:08 hamada>
package work

import (
	"fmt"
	"time"
)

func String2buf() {
	var s0 = fmt.Sprintf("%v", time.Nanosecond)
	var s1 = []byte(s0)
	fmt.Println(s0)
	fmt.Println(s1)
}
