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

	for pNode := linkedl; pNode != nil; pNode = pNode.next {
		fmt.Print(pNode.data, " ")
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

	var hashtable = MakeMap[int]()
	hashtable.Set([]byte("asdf"), 4)
	hashtable.Set([]byte("qwe"), 77)
	hashtable.Set([]byte("asdf"), 2)

	key := "a"
	for i := 0; i < 50; i++ {
		hashtable.Set([]byte(key), i)
		key += "a"
	}

	fmt.Println(hashtable.Get([]byte("asdf")))

	for i := 0; i < hashtable.cap; i++ {
		for pNode := *Advanced(hashtable.pA, i); pNode != nil; pNode = pNode.next {
			fmt.Println(string(pNode.data.key), pNode.data.value)
		}
	}

	hashtable.Destruct()

}
