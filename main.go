package main

import (
	"fmt"
)

type Map[T any] struct {
	rdyn_array[*Node[Pair[T]]]
	len int
}

func make_map[T any]() Map[T] {
	hashtable := Map[T]{rmm_make[*Node[Pair[T]]](16), 0}
	for i, _ := range hashtable.rdyn_array {
		hashtable.rdyn_array[i] = nil
	}
	return hashtable
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
	fmt.Print("\n")
	linkedl.destruct()

	slice := mm_make[int](1, 500)
	slice.append(5)
	fmt.Println(slice, len(slice), cap(slice))
	slice.append(8)
	slice[0] = 23
	fmt.Println(slice, len(slice), cap(slice))
	slice.append(13)
	slice.append(16)
	slice.append(234)
	fmt.Println(slice, len(slice), cap(slice))
	slice.destruct()

	var hashtable = make_map[int]()
	hashtable.set([]byte("asdf"), 4)
	hashtable.set([]byte("qwe"), 77)
	hashtable.set([]byte("asdf"), 2)

	fmt.Println(hashtable.get([]byte("asdf")))

	for i := 0; i < 16; i++ {
		for p_node := hashtable.rdyn_array[i]; p_node != nil; p_node = p_node.next {
			fmt.Println(string(p_node.data.key), p_node.data.value)
		}
	}

	hashtable.destruct()

}
