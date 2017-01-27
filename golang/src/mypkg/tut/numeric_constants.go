//Time-stamp: <2017-01-28 03:20:10 hamada>
package tut

import "fmt"

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func Numeric_constants() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	/*
		if false {
			fmt.Println(needInt(Big))
		}
	*/

}
