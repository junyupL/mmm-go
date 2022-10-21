package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"bytes"
	"hash/fnv"
)

type Pair[V any] struct {
	key   []byte
	value V
}

func FNV32a(bytes []byte) uint32 {
	algorithm := fnv.New32a()
	algorithm.Write(bytes)
	return algorithm.Sum32()
}

type Map[T any] struct {
	pA  **Node[Pair[T]]
	len int
	cap int
}

func For[T any](hashtable Map[T], loop func(pair Pair[T])) {
	for i := 0; i < hashtable.cap; i++ {
		for pNode := *Advanced(hashtable.pA, i); pNode != nil; pNode = pNode.next {
			loop(pNode.data)

		}
	}
}

func MakeMap[T any]() Map[T] {
	hashtable := Map[T]{New[*Node[Pair[T]]](16), 0, 16}
	for i := 0; i < hashtable.cap; i++ {
		*Advanced(hashtable.pA, i) = nil
	}
	return hashtable
}

func (hasht *Map[V]) Set(key []byte, value V) {
	if hasht.len == hasht.cap {
		//resize the hashtable
		newTable := New[*Node[Pair[V]]](C.ulonglong(hasht.cap * 2))
		for i := 0; i < hasht.cap*2; i++ {
			*Advanced(newTable, i) = nil
		}
		for i := 0; i < hasht.cap; i++ {
			for pNode := *Advanced(hasht.pA, i); pNode != nil; {
				//set node to new hashtable
				hash := uint32(hasht.cap*2-1) & FNV32a(pNode.data.key)
				tempPNode := pNode
				pNode = pNode.next
				insertNode(Advanced(newTable, hash), 0, tempPNode)

			}
		}
		Delete(hasht.pA)
		hasht.cap *= 2
		hasht.pA = newTable

	}
	hash := uint32(hasht.cap-1) & FNV32a(key)

	for pNode := *Advanced(hasht.pA, int(hash)); pNode != nil; pNode = pNode.next {
		if bytes.Equal(pNode.data.key, key) {

			pNode.data.value = value
			return
		}
	}
	insert(Advanced(hasht.pA, hash), 0, Pair[V]{key, value})
	hasht.len++
}

func (hasht *Map[V]) Get(key []byte) V {
	hash := uint32(hasht.cap-1) & FNV32a(key)
	for pNode := *Advanced(hasht.pA, hash); pNode != nil; pNode = pNode.next {

		if bytes.Equal(pNode.data.key, key) {

			return pNode.data.value
		}
	}
	panic("no value for key provided")
}

func (hasht Map[V]) Destruct() {
	for i := 0; i < hasht.cap; i++ {
		(*Advanced(hasht.pA, i)).Destruct()
	}
	Delete(hasht.pA)
}
