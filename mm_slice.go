package main

/*
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Slice[T any] struct {
	addr *T
	len  int
	cap  int
}

type DynArray[T any] []T
type DDynArray[T Destructable] DynArray[T]

func Make[T any](length int, capacities ...int) DynArray[T] {
	var capacity int
	if len(capacities) > 0 {
		capacity = 1
		for _, cap := range capacities {
			capacity *= cap
		}
	} else {
		capacity = length
	}
	var sliceObject = Slice[T]{New[T](C.ulonglong(capacity)), length, capacity}
	return *(*[]T)(unsafe.Pointer(&sliceObject))
}
func DMake[T Destructable](length int, capacities ...int) DDynArray[T] {
	return DDynArray[T](Make[T](length, capacities...))
}

func (pS *DynArray[T]) Append(data T) {
	if len(*pS) < cap(*pS) {
		*pS = append(*pS, data)
		return
	}

	pSlice := (*Slice[T])(unsafe.Pointer(pS))
	if cap(*pS) == 0 {
		pSlice.cap = 1
		pSlice.addr = New[T](C.ulonglong(1))
	} else {
		pSlice.cap *= 2
		pSlice.addr = Renew[T](unsafe.Pointer(pSlice.addr), C.ulonglong(cap(*pS)))
	}
	*Advanced(pSlice.addr, len(*pS)) = data
	pSlice.len += 1

}

func (pS *DDynArray[T]) Append(data T) {
	(*DynArray[T])(pS).Append(data)
}

func (pS *DynArray[T]) ResizeCap(capacity int) {
	pSlice := (*Slice[T])(unsafe.Pointer(pS))
	pSlice.cap = capacity

	pSlice.addr = Renew[T](unsafe.Pointer(pSlice.addr), C.ulonglong(cap(*pS)))

}
func (pS *DDynArray[T]) ResizeCap(capacity int) {
	(*DynArray[T])(pS).ResizeCap(capacity)
}

func (pS *DynArray[T]) Destruct() {
	pSlice := (*Slice[T])(unsafe.Pointer(pS))
	Delete(pSlice.addr)
}

func (pS *DDynArray[T]) Destruct() {
	for _, element := range *pS {
		element.Destruct()
	}

	(*DynArray[T])(pS).Destruct()
}
