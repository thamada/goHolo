//Time-stamp: <2017-01-25 18:05:33 hamada>
package exer

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	ret := fmt.Sprintf("cannot Sqrt negative number: -7")
	return ret
}

func Sqrt(x float64) (float64, error) {

	return 0, ErrNegativeSqrt(x)
}

func Errors() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
