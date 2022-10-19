package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

func sizeof[T any]() C.ulonglong {
	var a T
	return C.ulonglong(unsafe.Sizeof(a))
}

func New[T any](sizes ...C.ulonglong) *T {
	var factor C.ulonglong = 1
	for _, size := range sizes {
		factor *= size
	}
	return (*T)(C.malloc(factor * sizeof[T]()))
}

func Delete[T any](ptr *T) {
	C.free(unsafe.Pointer(ptr))
}

func Renew[T any](ptr unsafe.Pointer, sizes ...C.ulonglong) *T {

	var factor C.ulonglong = 1
	for _, size := range sizes {
		factor *= size
	}
	return (*T)(C.realloc(ptr, factor*sizeof[T]()))
}

func Advanced[T any, I constraints.Integer](ptr *T, count I) *T {
	return (*T)(unsafe.Add(unsafe.Pointer(ptr), C.ulonglong(count)*sizeof[T]()))
}
