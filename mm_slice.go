package mmm

import (
	"unsafe"

	"github.com/junyupL/mmm-go/memory"
)

type sliceHeader[T any] struct {
	Data *T
	Len  int
	Cap  int
}

type DynArray[T any, A memory.Allocator] struct {
	Slice []T
	a     A
}
type DDynArray[T Destructable, A memory.Allocator] DynArray[T, A]

func SliceToHeader[T any](pSlice *[]T) *sliceHeader[T] {
	return (*sliceHeader[T])(unsafe.Pointer(pSlice))
}

func HeaderToSlice[T any](pSlice *sliceHeader[T]) *[]T {
	return (*[]T)(unsafe.Pointer(pSlice))
}

func Make[T any, A memory.Allocator](a A, length int, capacities ...int) DynArray[T, A] {
	var capacity int
	if len(capacities) > 0 {
		capacity = 1
		for _, cap := range capacities {
			capacity *= cap
		}
	} else {
		capacity = length
	}
	return DynArray[T, A]{*HeaderToSlice(&sliceHeader[T]{memory.NewA[T](a, capacity), length, capacity}), a}
}

func DMake[T Destructable, A memory.Allocator](a A, length int, capacities ...int) DDynArray[T, A] {
	return DDynArray[T, A](Make[T](a, length, capacities...))
}

func (pS *DynArray[T, A]) Append(data T) {
	if len(pS.Slice) < cap(pS.Slice) {
		h := SliceToHeader(&pS.Slice)
		*memory.Advanced(h.Data, h.Len) = data
		h.Len++
		return
	}

	pSlice := SliceToHeader(&pS.Slice)
	if pSlice.Cap == 0 {
		pSlice.Cap = 1
		pSlice.Data = memory.New[T](pS.a)
	} else {
		pSlice.Cap *= 2
		pSlice.Data = memory.RenewA(pS.a, pSlice.Data, pSlice.Cap)
	}
	*memory.Advanced(pSlice.Data, pSlice.Len) = data
	pSlice.Len += 1

}

func (pS *DDynArray[T, A]) Append(data T) {
	(*DynArray[T, A])(pS).Append(data)
}

func (pS *DynArray[T, A]) ResizeCap(capacity int) {
	pSlice := SliceToHeader(&pS.Slice)
	pSlice.Cap = capacity

	pSlice.Data = memory.RenewA(pS.a, pSlice.Data, capacity)

}
func (pS *DDynArray[T, A]) ResizeCap(capacity int) {
	(*DynArray[T, A])(pS).ResizeCap(capacity)
}

func (pS *DynArray[T, A]) Destruct() {
	pSlice := SliceToHeader(&pS.Slice)
	memory.Delete(pS.a, pSlice.Data)
}

func (pS *DDynArray[T, A]) Destruct() {
	for _, element := range pS.Slice {
		element.Destruct()
	}

	(*DynArray[T, A])(pS).Destruct()
}
