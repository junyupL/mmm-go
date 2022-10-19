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
	fmt.Print("\n")
	linkedl.Destruct()

	slice := Make[int](1, 3)
	slice.Append(5)
	fmt.Println(slice, len(slice), cap(slice))
	slice.Append(8)
	slice[0] = 23
	fmt.Println(slice, len(slice), cap(slice))
	slice.Append(13)
	slice.Append(16)
	slice.Append(234)
	fmt.Println(slice, len(slice), cap(slice))
	slice.Destruct()

	var hashtable = make_map[int]()
	hashtable.set([]byte("asdf"), 4)
	hashtable.set([]byte("qwe"), 77)
	hashtable.set([]byte("asdf"), 2)

	fmt.Println(hashtable.get([]byte("asdf")))

	for i := 0; i < hashtable.cap; i++ {
		for p_node := *Advanced(hashtable.pA, i); p_node != nil; p_node = p_node.next {
			fmt.Println(string(p_node.data.key), p_node.data.value)
		}
	}

	hashtable.Destruct()

}
