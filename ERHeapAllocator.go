package mmm

/*
import (
	"unsafe"

	"github.com/junyupL/mmm-go/memory"
	"github.com/junyupL/sys/windows"
)

// Exponential Reuse Heap Allocator
type ERHeapAllocator struct {
	beginPtr    unsafe.Pointer
	ptr         unsafe.Pointer
	pos         uint32
	length      uint32
	listOfFreed List[Pair[uint32, unsafe.Pointer], memory.StatelessHeapAllocator]
}

func InitER(size uint32) (a ERHeapAllocator) {
	ptr, _ := windows.HeapAlloc(memory.HEAP, 0, size)
	a.ptr = unsafe.Pointer(ptr)
	a.beginPtr = a.ptr
	a.length = size
	a.listOfFreed.a = memory.HeapAllocator
	return
}
func (a ERHeapAllocator) Destruct() {
	windows.HeapFree(HEAP, 0, uintptr(a.beginPtr))
}

func (a *ERHeapAllocator) Alloc(size uintptr) unsafe.Pointer {
	a.listOfFreed.For(func(pNode *Node[unsafe.Pointer]) {
		if size <= pNode.key {

		}
	})
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

func (a *ERHeapAllocator) ReAlloc(ptr unsafe.Pointer, size uintptr) unsafe.Pointer {
	if a.pos+uint32(size) > a.length {
		for a.pos+uint32(size) > a.length {
			a.length *= 2
		}
		p, _ := windows.HeapReAlloc(memory.HEAP, 0, uintptr(a.beginPtr), a.length)
		a.beginPtr = unsafe.Pointer(p)

		a.ptr = unsafe.Add(a.beginPtr, a.pos)
	}
	retPtr := a.ptr
	a.ptr = unsafe.Add(a.ptr, size)
	a.pos += uint32(size)
	return retPtr
}

func (a *ERHeapAllocator) Free(ptr unsafe.Pointer) {
}
*/
