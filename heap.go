package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func sizeof[T any]() C.ulonglong {
	var a T
	return C.ulonglong(unsafe.Sizeof(a))
}

func new[T any](sizes ...C.ulonglong) *T {
	var factor C.ulonglong = 1
	for _, size := range sizes {
		factor *= size
	}
	return (*T)(C.malloc(factor * sizeof[T]()))
}

func delete[T any](ptr *T) {
	C.free(unsafe.Pointer(ptr))
}

type Node[T any] struct {
    data T
	next *Node[T]
}

func (base *Node[T])insert(index int, d T) {
	if index == 0 {
		p_node := new[Node[T]]()
		*p_node = *base
		base.data = d
		base.next = p_node
		return
	}
	var curr *Node[T] = base
	for i := 1; i < index; i++ {
		curr = curr.next
		if curr == nil {
			//index out of bounds
			return
		}
	}
	p_node := new[Node[T]]()
	p_node.data = d
	if curr.next == nil {
		
		p_node.next = nil
		curr.next = p_node

	} else {
		p_node.next = curr.next
		curr.next = p_node
	}

}
func (base Node[T]) destruct() {
	for curr := base.next; curr != nil; {
		temp := curr.next
		delete(curr)
		curr = temp
		
	}
}


func main() {
	var l Node[int]
	l.data = 1
	l.insert(0, 4)
	l.insert(2, 7)
	l.insert(1, 10)
	
	fmt.Print(l.data, l.next.data, l.next.next.data, l.next.next.next.data)
	l.destruct()
}
