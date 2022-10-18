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

	slice := mm_make[int](1, 500)
	slice.append(5)
	fmt.Println(slice, len(slice), cap(slice))
	slice.append(8)
	fmt.Println(slice, len(slice), cap(slice))
	slice.append(13)
	slice.append(12)
	slice.append(16)
	slice.append(234)
	fmt.Println(slice, len(slice), cap(slice))
	slice.destruct()

	var hashtable = rmm_make[*Node[Pair[int]]](16)
	for i, _ := range hashtable {
		hashtable[i] = nil
	}
	rset_map(&hashtable, []byte("asdf"), 4)
	rset_map(&hashtable, []byte("qwe"), 77)
	rset_map(&hashtable, []byte("asdf"), 2)

	fmt.Println(rget_map(&hashtable, []byte("asdf")))

	for i := 0; i < 16; i++ {
		for p_node := hashtable[i]; p_node != nil; p_node = p_node.next {
			fmt.Println(string(p_node.data.key), p_node.data.value)
		}
	}

	hashtable.destruct()

}
