package memory

/*
#include <stdlib.h>
*/
import "C"
import "unsafe"

type StatelessMAllocator struct{}

var MAllocator StatelessMAllocator

func (a StatelessMAllocator) Alloc(size uintptr) unsafe.Pointer {
	return C.malloc(C.ulonglong(size))
}

func (a StatelessMAllocator) ReAlloc(ptr unsafe.Pointer, size uintptr) unsafe.Pointer {
	return C.realloc(ptr, C.ulonglong(size))
}

func (a StatelessMAllocator) Free(ptr unsafe.Pointer) {
	C.free(ptr)
}

type Allocator interface {
	Alloc(uintptr) unsafe.Pointer
	ReAlloc(unsafe.Pointer, uintptr) unsafe.Pointer
	Free(unsafe.Pointer)
}
