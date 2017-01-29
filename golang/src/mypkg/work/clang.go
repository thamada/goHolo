//Time-stamp: <2017-01-30 03:03:15 hamada>
package work

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func Clang() {
	p := C.CString("Hello from C language!")
	C.puts(p)
	C.free(unsafe.Pointer(p))
}
