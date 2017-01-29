//Time-stamp: <2017-01-30 03:10:47 hamada>
package work

/*
#include <stdio.h>
#include <stdlib.h>
*/
import (
	"C"
	"unsafe"
)

func Clang() {
	p := C.CString("Hello from C language!")
	C.puts(p)
	C.free(unsafe.Pointer(p))
}
