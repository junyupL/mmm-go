package memory

import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

func sizeof[T any]() uintptr {
	var a T
	return unsafe.Sizeof(a)
}

func Advanced[T any, I constraints.Integer](ptr *T, count I) *T {
	return (*T)(unsafe.Add(unsafe.Pointer(ptr), uintptr(count)*sizeof[T]()))
}

func New[T any, A Allocator](a A) *T {
	return (*T)(a.Alloc(sizeof[T]()))
}
func NewA[T any, A Allocator, I constraints.Integer](a A, size I) *T {
	return (*T)(a.Alloc(uintptr(size) * sizeof[T]()))
}

func Delete[T any, A Allocator](a A, ptr *T) {
	a.Free(unsafe.Pointer(ptr))
}
func Renew[T any, A Allocator](a A, ptr *T) *T {
	return (*T)(a.ReAlloc(unsafe.Pointer(ptr), sizeof[T]()))
}

func RenewA[T any, A Allocator, I constraints.Integer](a A, ptr *T, size I) *T {
	return (*T)(a.ReAlloc(unsafe.Pointer(ptr), uintptr(size)*sizeof[T]()))
}
