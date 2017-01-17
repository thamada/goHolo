//Time-stamp: <2017-01-18 04:54:15 hamada>
// A Tour of Go
package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
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

func __pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// can use v here
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func Sqrt(x float64) float64 {
	z := x
	zz := float64(0)

	for i := 0; i < 10; i++ {
		zz = z - (z*z-x)/(2*z)
		//		fmt.Printf("%d: %g\n", i, zz)
		if zz == z {
			break
		} else {
			z = zz
		}
	}

	return z
}

func switch_test() {

	fmt.Println("GOOS: ", runtime.GOOS)
	fmt.Println("GOARCH: ", runtime.GOARCH)
	fmt.Println("GOROOT: ", runtime.GOROOT)
	fmt.Println("NumCPU: ", runtime.NumCPU())
	fmt.Println("NumCgoCall: ", runtime.NumCgoCall())
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	fmt.Println("ReadTrace: ", runtime.ReadTrace())
	fmt.Println("Version: ", runtime.Version())
	fmt.Println("CPUProfile: ", runtime.CPUProfile())
	if 0 == 1 {
		runtime.Breakpoint()
	}

	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

}

func time_test() {
	t := time.Now()

	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Weekday())

	t = time.Date(2016, 12, 31, 23, 59, 60, 0, time.UTC)
	fmt.Println(t)
	t = time.Date(2016, 12, 31, 23, 59, 60, 0, time.Local)
	fmt.Println(t)

}

func main() {

	time_test()

	if 0 == 1 {
		switch_test()
	}

	if 0 == 1 {
		e_max := 0.0
		for i := 0; i < 0x1000; i++ {
			x := float64(i)
			z0 := Sqrt(x)
			z1 := math.Sqrt(x)
			e_rel := (z1 - z0) / z1
			e_rel = math.Abs(e_rel)
			if e_rel > e_max {
				e_max = e_rel
				fmt.Println(i, z0, z1, e_max)
			}
		}
	}

	if 0 == 1 {
		fmt.Println(
			pow(3, 3, 10),
			pow(3, 4, 20),
		)
		fmt.Println(sqrt(2))
		fmt.Println(sqrt(-4))
		fmt.Printf("%T\n", sqrt(2))
		fmt.Printf("%T\n", sqrt(-4))
		fmt.Println(Small)
		fmt.Println(float64(Big))
		fmt.Println(needInt(Small))
		fmt.Println(needFloat(Small))
		fmt.Println(needFloat(Big))
		fmt.Printf("%T\n", Small)
		fmt.Printf("%T\n", Big>>38)
		fmt.Printf("0x%x\n", Big>>37-1)
	}
}
