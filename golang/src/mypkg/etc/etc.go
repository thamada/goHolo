//Time-stamp: <2017-01-25 19:03:22 hamada>
package etc

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!
// GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!
// GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!GREAT!!
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func forever_loop() {
	// unreachable code same as while(true)
	/*
		for {
		}
		fmt.Println("Forever")
	*/
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

func switch_test() {

	fmt.Println("GOOS: ", runtime.GOOS)
	fmt.Println("GOARCH: ", runtime.GOARCH)
	fmt.Println("GOROOT: ", runtime.GOROOT())
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

func struct_test() {

	type Vertex struct {
		X int
		Y int
	}

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

func slice_of_slice_test() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Printf("%T\n", board)

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func appending_to_a_slice_test() {

	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	// cannot use v (type []int) as type int in append
	/*
		v := []int{5, 6, 7, 8, 9}
		s = append(s, v)
		printSlice(s)
	*/
	fmt.Println("more details about slices, ")
	fmt.Println("\tsee: https://blog.golang.org/go-slices-usage-and-internals")
}

func range_test() {

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 777}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for _, v := range pow {
		fmt.Printf("value only: %d\n", v)
	}

	for i, _ := range pow {
		fmt.Printf("index only: %d\n", i)
	}

	for i := range pow {
		fmt.Printf("index only: %d\n", i)
	}

}

func map_test() {

	type Vertex struct {
		Lat, Long float64
	}

	m := make(map[string]Vertex)

	fmt.Printf("%T\n", m)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

}

func map_literals_test() {
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

	var m2 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m2)

	var m3 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)

	var m4 = map[string]Vertex{
		"Bell Labs": Vertex{40.68433, -74.39967},
		"Google":    Vertex{37.42202, -122.08408},
	}
	fmt.Println(m4)

}

func maintain_map_test() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func func_value_test() {
	compute := func(fn func(float64, float64) float64) float64 {
		return fn(3, 4)
	}

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	fmt.Printf("hypot:    %T\n", hypot)
	fmt.Printf("math.Pow: %T\n", math.Pow)
	fmt.Printf("compute:  %T\n", compute)
}

func closures_test() {

	adder := func() func(int) int {
		sum := 0 // bind to closure
		return func(x int) int {
			sum += x
			return sum
		}
	}

	pos, neg := adder(), adder()

	fmt.Printf("%T\n", adder)
	fmt.Printf("%T\n", neg)
	fmt.Printf("%T\n", neg)

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fb_test() {

	fibo := func() func() int {
		z0 := 0
		z1 := 1
		fmt.Println(z0)
		fmt.Println(z1)
		return func() int {
			z2 := z1 + z0
			z0 = z1
			z1 = z2
			return z2
		}
	}

	f := fibo() // f is a closure

	for i := 0; i < 30; i++ {
		fmt.Println(f())
	}

}

func pointer_game() {

	//	type ZEON float64

	type ZEON struct {
		x float64
	}

	z := &ZEON{1.23}
	fmt.Printf("z:  (%v, %T) (%v, %T)\n", z, z, *z, *z)
	fmt.Printf("z:  (%v, %T) (%v, %T)\n", z, z, &z, &z)
	z2 := &z
	fmt.Printf("z2: (%v, %T) (%v, %T)\n", z2, z2, &z2, &z2)
	z3 := &z2
	fmt.Printf("z3: (%v, %T) (%v, %T)\n", z3, z3, &z3, &z3)
	z4 := &z3
	fmt.Printf("z4: (%v, %T) (%v, %T)\n", z4, z4, &z4, &z4)
	fmt.Printf("(%v, %T) (%v, %T)\n", ****z4, ****z4, ***z4, ***z4)

	var z5 *****ZEON
	fmt.Printf("z5: (%v, %T) (%v, %T)\n", z5, z5, &z5, &z5)
	z5 = &z4
	fmt.Printf("z5: (%v, %T) (%v, %T)\n", z5, z5, &z5, &z5)
	fmt.Printf("(%v, %T)\n", ****z5, ****z5)

	/*
	     // cannot take the address of ZEON(7.77)
	   	w := &ZEON(7.77)
	   	fmt.Printf("(%v, %T)\n", w, w)
	*/

	w := &ZEON{7.77}
	fmt.Printf("(%v, %T)\n", w, w)

}

//------------------------ interface_test
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func interface_test() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())
}

//------------------------ implicit_interface_test
type I interface {
	M()
}

type II interface {
	MM()
}

type T struct {
	S string
}

// type T implements the interface I,
func (t T) M() {
	fmt.Println("T.M(): ", t.S)
}

// type T implements the interface II,
func (t *T) MM() {
	fmt.Println("&T.MM():", t.S)
}

type TT struct {
	S string
}

// type TT implements the interface I,
func (t TT) M() {
	fmt.Println("TT.M(): ", t.S)
}

// type TT implements the interface II,
func (t *TT) MM() {
	fmt.Println("&TT.M(): ", t.S)
}

type F struct {
	x float64
}

func (f F) M() {
	fmt.Println(f.x)
}

func (f *F) MM() {
	fmt.Println(f.x)
}

type F64 float64

func (f F64) M() {
	fmt.Println(f)
}

func (f *F64) MM() {
	fmt.Println(f)
}

type F64STR struct {
	x float64
}

func (f F64STR) M() {
	fmt.Println(f)
}

func (f *F64STR) MM() {
	if f == nil {
		fmt.Println("receiver is nil object: nil comes from nihil.", f)
		return
	}
	fmt.Println("(*F64STR).x: ", f.x)
}

func ducktyping_test() {
	var a I
	var b II

	fmt.Println(a, b)
	fmt.Printf("(%v, %T)\n", a, a)
	fmt.Printf("(%v, %T)\n", b, b)

	a = T{"Interface I"}
	b = &T{"Interface II"}
	a.M()
	b.MM()
	fmt.Printf("(%v, %T)\n", a, a)
	fmt.Printf("(%v, %T)\n", b, b)

	a = TT{"Interface I"}
	b = &TT{"Interface II"}
	a.M()
	b.MM()
	fmt.Printf("(%v, %T)\n", a, a)
	fmt.Printf("(%v, %T)\n", b, b)

	a = F{math.Pi}
	b = &F{-math.Pi}
	a.M()
	b.MM()
	fmt.Printf("(%v, %T)\n", a, a)
	fmt.Printf("(%v, %T)\n", b, b)

	/*
		a = F64(math.Pi)
		b = &F64(-math.Pi)
		a.M()
		b.MM()
		fmt.Printf("(%v, %T)\n", a, a)
		fmt.Printf("(%v, %T)\n", b, b)
	*/

	a = F64(math.Pi)
	a.M()
	fmt.Printf("(%v, %T)\n", a, a)
	fmt.Printf("(%v, %T)\n", &a, &a)

	// cannot use F64(math.Pi) (type F64) as type II in assignment:
	// F64 does not implement II (MM method has pointer receiver)
	/*
		b = F64(math.Pi)
		fmt.Printf("(%v, %T)\n", &b, &b)
	*/

	// cannot use &a (type *I) as type II in assignment:
	// *I is pointer to interface, not interface
	/*
		b = &a
		fmt.Printf("(%v, %T)\n", b, b)
	*/

	//cannot take the address of F64(math.Pi)
	/*
		b = &F64(math.Pi)
		fmt.Printf("(%v, %T)\n", b, b)
	*/

	// no problem
	b = &F64STR{math.Pi}
	b.MM()
	fmt.Printf("(%v, %T)\n", b, b)

	// no problem
	{
		var c *F64
		d := F64(math.Pi)
		c = &d
		c.MM()
		fmt.Printf("(%v, %T)\n", c, c)
	}

	{
		var c *F64STR
		b = c
		b.MM()
		fmt.Printf("(%v, %T)\n", b, b)
	}

	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal SIGSEGV: segmentation violation code=0x1 addr=0x20 pc=0x3799]
	*/
	if false {
		var c I
		c.M()
		fmt.Printf("(%v, %T)\n", c, c)
	}

	{
		var c interface{}
		fmt.Printf("(%v, %T)\n", c, c)
		c = math.Pi
		fmt.Printf("(%v, %T)\n", c, c)
		c = int64(777)
		fmt.Printf("(%v, %T)\n", c, c)
		c = F64(math.Pi)
		fmt.Printf("(%v, %T)\n", c, c)
		//		c.M() // c.M undefined (type interface {} is interface with no methods)
	}

	{
		fmt.Println("--- Type Assertion ---")
		var i interface{} = "hello"

		s, ok := i.(string)
		fmt.Println(s, ok)

		f, ok := i.(float64)
		fmt.Println(f, ok)

		// f = i.(float64) // panic
		// fmt.Println(f)
	}

	{
		do := func(i interface{}) {
			switch v := i.(type) {
			case int:
				fmt.Printf("Twice %v is %v\n", v, v*2)
			case string:
				fmt.Printf("%q is %v bytes long\n", v, len(v))
			default:
				fmt.Printf("I don't know about type %T!\n", v)
			}
		}
		do(21)
		do("hello")
		do(true)
		do(do)
		do(nil)
	}

	{
		var a interface{} = 4649
		var b int64
		fmt.Println(a)
		b = reflect.ValueOf(a).Int()
		fmt.Println(b)
	}

	{
		type Autonomous interface {
			drive() bool
		}

		type Airplane interface {
			fly() bool
		}

		car := Honda{"Nbox", false}
		/*
			fmt.Printf("(%v, %T): %v\n", car, car, car.drive())
			fmt.Printf("(%v, %T): %v\n", car, car, car.fly())
		*/
		check_auto := func(c Autonomous) {
			fmt.Printf("Autonomous: (%v, %T)\n", c, c)
		}

		/*
			If Honda doesn't have fly() method, the following error happens:
					  cannot use car (type Honda) as type Airplane in argument to check_plane:
						Honda does not implement Airplane (missing fly method)
		*/
		check_plane := func(c Airplane) {
			fmt.Printf("Airplane: (%v, %T)\n", c, c)
		}

		check_auto(car)
		check_plane(car)
	}

	{
		Sqrt := func(x float64) (float64, error) {
			err := ErrNegativeSqrt(x)
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
			return z, err
		}

		fmt.Println(Sqrt(2))
		fmt.Println(Sqrt(-2))
	}

}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprint(float64(e))
	} else {
		return ""
	}
}

type Honda struct {
	name      string
	autodrive bool
}

func (h Honda) drive() bool {
	return true
}

func (h Honda) fly() bool {
	return true
}

func Main() {

	ducktyping_test()

	if false {
		pointer_game()
		interface_test()
		fb_test()
		closures_test()
		func_value_test()
		maintain_map_test()
		map_literals_test()
		map_test()
		range_test()
		appending_to_a_slice_test()
		slice_of_slice_test()
		slice_defaults_test()
		slice_literals_test()
		array_test()
		struct_test()
		pointer_test()
		time_test()
		switch_test()
	}

	if false {

		Sqrt := func(x float64) float64 {
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

	if false {
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
