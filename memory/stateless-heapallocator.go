package memory

import (
	"unsafe"

	"github.com/junyupL/sys/windows"
)

type StatelessHeapAllocator struct{}

var HeapAllocator StatelessHeapAllocator

func (a StatelessHeapAllocator) Alloc(size uintptr) unsafe.Pointer {
	ptr, _ := windows.HeapAlloc(HEAP, 0, uint32(size))
	return unsafe.Pointer(ptr)
}

func (a StatelessHeapAllocator) ReAlloc(ptr unsafe.Pointer, size uintptr) unsafe.Pointer {
	p, _ := windows.HeapReAlloc(HEAP, 0, uintptr(ptr), uint32(size))
	return unsafe.Pointer(p)
}

func (a StatelessHeapAllocator) Free(ptr unsafe.Pointer) {
	windows.HeapFree(HEAP, 0, uintptr(ptr))
}

// Exponential Never Free Heap Allocator
type ENFHeapAllocator struct {
	beginPtr unsafe.Pointer
	ptr      unsafe.Pointer
	pos      uint32
	length   uint32
}

func InitENF(size uint32) (a ENFHeapAllocator) {
	ptr, _ := windows.HeapAlloc(HEAP, 0, size)
	a.ptr = unsafe.Pointer(ptr)
	a.beginPtr = a.ptr
	a.length = size
	return
}
func (a ENFHeapAllocator) Destruct() {
	windows.HeapFree(HEAP, 0, uintptr(a.beginPtr))
}

func (a *ENFHeapAllocator) Alloc(size uintptr) unsafe.Pointer {
	if a.pos+uint32(size) > a.length {
		for a.pos+uint32(size) > a.length {
			a.length *= 2
		}
		p, _ := windows.HeapReAlloc(HEAP, 0, uintptr(a.beginPtr), a.length)
		a.beginPtr = unsafe.Pointer(p)

		a.ptr = unsafe.Add(a.beginPtr, a.pos)
	}
	retPtr := a.ptr
	a.ptr = unsafe.Add(a.ptr, size)
	a.pos += uint32(size)
	return retPtr
}

func (a *ENFHeapAllocator) ReAlloc(ptr unsafe.Pointer, size uintptr) unsafe.Pointer {
	if a.pos+uint32(size) > a.length {
		for a.pos+uint32(size) > a.length {
			a.length *= 2
		}
		p, _ := windows.HeapReAlloc(HEAP, 0, uintptr(a.beginPtr), a.length)
		a.beginPtr = unsafe.Pointer(p)

		a.ptr = unsafe.Add(a.beginPtr, a.pos)
	}
	retPtr := a.ptr
	a.ptr = unsafe.Add(a.ptr, size)
	a.pos += uint32(size)
	return retPtr
}

func (a *ENFHeapAllocator) Free(ptr unsafe.Pointer) {
}
