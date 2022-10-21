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
	var avlTree *HeightNode[int]
	key := "a"
	for i := 0; i < 50; i++ {
		hashtable.Set([]byte(key), i)
		AVLInsert(&avlTree, i)
		key += "a"
	}
	For(hashtable, func(pair Pair[int]) {
		fmt.Println(string(pair.key), pair.value)
	})
	fmt.Println(hashtable.Get([]byte("aaa")))

	avlTree.PrintTree()

	avlTree.Destruct()
	hashtable.Destruct()
}
