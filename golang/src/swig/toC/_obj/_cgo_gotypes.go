// Created by cgo - DO NOT EDIT

package toC

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, ...interface{}) interface{}

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})

//go:cgo_export_dynamic Waitgroup
//go:linkname _cgoexp_09a449653019_Waitgroup _cgoexp_09a449653019_Waitgroup
//go:cgo_export_static _cgoexp_09a449653019_Waitgroup
//go:nosplit
//go:norace
func _cgoexp_09a449653019_Waitgroup(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_09a449653019_Waitgroup
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_09a449653019_Waitgroup() {
	Waitgroup()
}
