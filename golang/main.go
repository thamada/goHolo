//Time-stamp: <2017-01-17 23:28:02 hamada>
// A Tour of Go
package main

import (
	"fmt"
	"math"
)

// GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!
// GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!
// GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100 
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func forever_loop() {
	// while(true)
	for {
	}
	fmt.Println("Forever")
}

func sqrt(x float64) string {
	if x < 0 {
		x = -x
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-4))
	fmt.Printf("%T\n", sqrt(2))
	fmt.Printf("%T\n", sqrt(-4))

	if 0==1 {
		fmt.Println(Small)
		fmt.Println(float64(Big))
		fmt.Println(needInt(Small))
		fmt.Println(needFloat(Small))
		fmt.Println(needFloat(Big))
		fmt.Printf("%T\n", Small)
		fmt.Printf("%T\n", Big>>38)
		fmt.Printf("0x%x\n", Big>>37 -1)
	}
}
