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

func insert[T any](p_list **Node[T], index int, d T) {
	if index < 0 {
		//index out of bounds
		return
	}

	if index == 0 {
		temp := *p_list
		*p_list = new[Node[T]]()
		(*p_list).data = d
		(*p_list).next = temp
		return
	}
	//pointer to node to add the new node after
	var to_add *Node[T] = *p_list
	if to_add == nil {
		//index out of bounds
		return
	}
	for i := 1; i < index; i++ {
		to_add = to_add.next
		if to_add == nil {
			//index out of bounds
			return
		}
	}
	p_node := new[Node[T]]()
	p_node.data = d
	p_node.next = to_add.next
	to_add.next = p_node
	

}
func (p_node *Node[T]) destruct() {
	for curr := p_node; curr != nil; {
		temp := curr.next
		delete(curr)
		curr = temp
		
	}
}


func main() {
	var linkedl *Node[int] = nil
	insert(&linkedl, 0, 1)
	insert(&linkedl, 0, 4)
	insert(&linkedl, 2, 7)
	insert(&linkedl, 1, 10)
	
	for p_node := linkedl; p_node != nil; p_node = p_node.next {
		fmt.Print(p_node.data, " ")
	}
	
	linkedl.destruct()
}
