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

type dyn_array[T any] []T

func mm_make[T any](length int, capacities ...int) dyn_array[T] {
	var capacity int
	if len(capacities) > 0 {
		capacity = 1
		for _, cap := range capacities {
			capacity *= cap
		}
	} else {
		capacity = length
	}
	var sliceObject = Slice[T]{mm_new[T](C.ulonglong(capacity)), length, capacity}
	return *(*[]T)(unsafe.Pointer(&sliceObject))
}

func (p_s *dyn_array[T]) append(data T) {
	if len(*p_s) < cap(*p_s) {
		*p_s = append(*p_s, data)
		return
	}

	p_slice := (*Slice[T])(unsafe.Pointer(p_s))
	if cap(*p_s) == 0 {
		p_slice.cap = 1
	} else {
		p_slice.cap *= 2
	}
	p_slice.addr = renew[T](unsafe.Pointer(p_slice.addr), C.ulonglong(cap(*p_s)))
	p_slice.len += 1
	(*p_s)[len(*p_s)-1] = data

}

func (p_s *dyn_array[T]) destruct() {
	p_slice := (*Slice[T])(unsafe.Pointer(p_s))
	delete(p_slice.addr)
}

type rdyn_array[T Destructable] []T

func rmm_make[T Destructable](length int, capacities ...int) rdyn_array[T] {
	var capacity int
	if len(capacities) > 0 {
		capacity = 1
		for _, cap := range capacities {
			capacity *= cap
		}
	} else {
		capacity = length
	}
	var sliceObject = Slice[T]{mm_new[T](C.ulonglong(capacity)), length, capacity}
	return *(*[]T)(unsafe.Pointer(&sliceObject))
}

func (p_s *rdyn_array[T]) append(data T) {
	if len(*p_s) < cap(*p_s) {
		*p_s = append(*p_s, data)
		return
	}

	p_slice := (*Slice[T])(unsafe.Pointer(p_s))
	if cap(*p_s) == 0 {
		p_slice.cap = 1
	} else {
		p_slice.cap *= 2
	}
	p_slice.addr = renew[T](unsafe.Pointer(p_slice.addr), C.ulonglong(cap(*p_s)))
	p_slice.len += 1
	(*p_s)[len(*p_s)-1] = data

}

func (p_s *rdyn_array[T]) destruct() {
	for _, element := range *p_s {
		element.destruct()
	}

	p_slice := (*Slice[T])(unsafe.Pointer(p_s))
	delete(p_slice.addr)
}
