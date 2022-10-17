package main

import (
	"fmt"
)

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

	slice := mm_make[int](0, 1)
	mm_append(&slice, 5)
	fmt.Println(slice, len(slice), cap(slice))
	mm_append(&slice, 8)
	fmt.Println(slice, len(slice), cap(slice))
	mm_append(&slice, 13)
	fmt.Println(slice, len(slice), cap(slice))
	slice_destruct(&slice)

}
