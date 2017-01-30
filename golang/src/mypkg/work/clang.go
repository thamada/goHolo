//Time-stamp: <2017-01-31 01:34:14 hamada>
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

	n := float64(8.0)
	x := C.double(n)
	z := C.sqrt(x)
	C.puts(n)
	C.puts(x)
	C.puts(z)
	C.free(unsafe.Pointer(n))
	C.free(unsafe.Pointer(x))
	C.free(unsafe.Pointer(z))
}
