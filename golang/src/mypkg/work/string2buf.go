//Time-stamp: <2017-01-29 02:47:39 hamada>
package work

import (
	"fmt"
	"time"
)

func String2buf() {
	var s0 = fmt.Sprintf("%T", time.Nanosecond)
	var s1 = []byte(s0)
	var s2 = uint32(s1[0])
	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("%T, %T, %T\n", s0, s1, s2)
}
