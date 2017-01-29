//Time-stamp: <2017-01-30 03:11:19 hamada>
package work

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

func Clang() {
	p := C.CString("Hello from C language!")
	C.puts(p)
	C.free(unsafe.Pointer(p))
}
