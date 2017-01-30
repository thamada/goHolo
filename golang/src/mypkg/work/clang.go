//Time-stamp: <2017-01-31 01:36:01 hamada>
package work

/*
#cgo LDFLAGS: -lm
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func Clang() {
	p := C.CString("Hello from C language!")
	fmt.Printf("%v, %T\n", p, p)
	C.puts(p)
	C.free(unsafe.Pointer(p))

	n := float64(16.0)
	x := C.double(n)
	z := C.sqrt(x)
	fmt.Println(n, x, z)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", z, z)
}
