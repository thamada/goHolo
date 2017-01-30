# C references to Go

Go functions can be exported for use by C code in the following way:

'''
//export MyFunction
func MyFunction(arg1, arg2 int, arg3 string) int64 {...}

//export MyFunction2
func MyFunction2(arg1, arg2 int, arg3 string) (int64, *C.char) {...}
'''
They will be available in the C code as:
'''
extern int64 MyFunction(int arg1, int arg2, GoString arg3);
extern struct MyFunction2_return MyFunction2(int arg1, int arg2, GoString arg3);
'''
found in the _cgo_export.h generated header, after any preambles copied from the cgo input files. Functions with multiple return values are mapped to functions returning a struct. Not all Go types can be mapped to C types in a useful way.

Using //export in a file places a restriction on the preamble: since it is copied into two different C output files, it must not contain any definitions, only declarations. If a file contains both definitions and declarations, then the two output files will produce duplicate symbols and the linker will fail. To avoid this, definitions must be placed in preambles in other files, or in C source files.

---
- https://golang.org/cmd/cgo/
- https://blog.golang.org/c-go-cgo

