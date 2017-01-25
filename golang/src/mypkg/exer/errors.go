//Time-stamp: <2017-01-25 18:19:19 hamada>
package exer

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	var ret string
	if e < 0 {
		v := float64(e)
		ret = fmt.Sprintf("cannot Sqrt negative number: %v", v)
	}
	return ret
}

func Sqrt(x float64) (float64, error) {

	w := x
	u := float64(0)

	for i := 0; i < 10; i++ {
		u = w - (w*w-x)/(2*w)
		if u == w {
			break
		} else {
			w = u
		}
	}

	return u, ErrNegativeSqrt(x)
}

func Errors() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
