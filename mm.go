package main

/*
#include <stdlib.h>
*/
import "C"
import "unsafe"

func sizeof[T any]() C.ulonglong {
	var a T
	return C.ulonglong(unsafe.Sizeof(a))
}

func mm_new[T any](sizes ...C.ulonglong) *T {
	var factor C.ulonglong = 1
	for _, size := range sizes {
		factor *= size
	}
	return (*T)(C.malloc(factor * sizeof[T]()))
}

func delete[T any](ptr *T) {
	C.free(unsafe.Pointer(ptr))
}

func renew[T any](ptr unsafe.Pointer, sizes ...C.ulonglong) *T {

	var factor C.ulonglong = 1
	for _, size := range sizes {
		factor *= size
	}
	return (*T)(C.realloc(ptr, factor*sizeof[T]()))
}
