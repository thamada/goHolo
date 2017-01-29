//Time-stamp: <2017-01-30 03:12:46 hamada>
package work

/*
#include <stdio.h>
#include <stdlib.h>
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
}
