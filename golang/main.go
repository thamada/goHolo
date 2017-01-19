//Time-stamp: <2017-01-19 12:56:45 hamada>
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

	const layout = "Now, Mon Jan. 02 15:04:05 JST 2006"
	fmt.Println(t.Format(layout))
	const layout2 = "Jan 02 15:04:05 UTC 2006"
	fmt.Println(t.Format(layout2))

	t = time.Now()
	s := t.String()
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", t)
	fmt.Printf("%T\n", t.String())

	today := time.Now().Weekday()
	fmt.Printf("%T\t", today)
	fmt.Println(today)
	fmt.Printf("%T\t", time.Saturday)
	fmt.Println(time.Saturday)
	/*
		  for i := 0; i < 10; i++ {
				fmt.Println(today + i)
			}
	*/

	fmt.Println(today + 0)
	fmt.Println(today + 1)
	defer fmt.Println("defer: today+2:", today+2)

	x := 1
	defer fmt.Println("defer: x:", x)
	fmt.Println("x:", x)
	x = 1111

	x = add(x)
	defer fmt.Println("x:", x)
	x = add(x)

	fmt.Println(today + 3)
	fmt.Println(today + 4) // runtime error
	fmt.Println(today - 0)
	fmt.Println(today - 1)
	fmt.Println(today - 2)
	fmt.Println(today - 3)
	fmt.Println(today - 4) // runtime error

	t = time.Now()
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}

func add(x int) int {
	return x + 1
}

func pointer_test() {
	i, j := 42, 2701

	var p *int
	fmt.Printf("%T\n", p)
	fmt.Println(p)
	p = &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	fmt.Printf("%T\n", p)

	p = &j
	*p = *p / 37
	fmt.Println(j)
	fmt.Printf("%T\n", p)
}

type Vertex struct {
	X int
	Y int
}

func struct_test() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
	v.Y = 1.2e9
	fmt.Println(v)

	p := &v
	p.X = 1e3
	fmt.Println(v)

	(*p).X = 1e3
	fmt.Println(v)

	p = &Vertex{20, 30}
	fmt.Println(*p)

	p = &Vertex{}
	fmt.Println(*p)

	fmt.Printf("%T\n", v)
	fmt.Printf("%T\n", p)
}

func array_test() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println(primes)

	for i := 0; i < 6; i++ {
		var s []int = primes[0 : i+1]
		fmt.Println(i+1, s)
	}

	for i := 0; i < 6; i++ {
		var s []int = primes[i:6]
		fmt.Println(i, s)
	}

	for i := 0; i < 6; i++ {
		var s []int = primes[i : i+1]
		fmt.Println(i, s)
	}

	{
		s := primes[1:3]
		fmt.Printf("%T\n", s)

		s[0] = 111
		fmt.Println(s)
		fmt.Println(primes)
	}

}

func slice_literals_test() {
	q := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("q: array: %T\n", q)
	fmt.Println(q)

	ql := []int{2, 3, 5, 7, 11, 13} // slice literal
	fmt.Printf("ql: slice: %T\n", ql)
	fmt.Println(ql)

	r := []bool{true, false, true, true, false, true} // slice literal
	fmt.Printf("r: slice: %T\n", r)
	fmt.Println(r)

	s := []struct { // slice literal
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Printf("s: slice: %T\n", s)
	fmt.Println(s)
}

func slice_defaults_test() {

	// len 6, cap 6
	s := []int{700, 701, 702, 703, 704, 705}
	printSlice(s)

	// len 0, cap 6
	s = s[:0]
	printSlice(s)

	// len 6, cap 6
	s = s[:6]
	printSlice(s)

	// len 4, cap 4
	s = s[2:]
	printSlice(s)

	// len 3, cap 3
	s = s[1:]
	printSlice(s)

	//	s = s[:4] // Error

	// len 0, cap 0
	s = s[3:]
	printSlice(s)

	if s != nil {
		fmt.Println("s is NOT nil!")
	}

	var ss []int
	s = ss
	printSlice(s)

	if s == nil {
		fmt.Println("s is nil!")
	}

	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:3]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)

	d[0] = 777
	fmt.Println(a, b, c, d)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(txt string, s []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", txt, len(s), cap(s), s)
}


func main() {

	slice_defaults_test()

	if false {
		slice_literals_test()
		array_test()
		struct_test()
		pointer_test()
		time_test()
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
