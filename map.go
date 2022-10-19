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

func make_map[T any]() Map[T] {
	hashtable := Map[T]{New[*Node[Pair[T]]](16), 0, 16}
	for i := 0; i < hashtable.cap; i++ {
		*Advanced(hashtable.pA, i) = nil
	}
	return hashtable
}

func (hasht *Map[V]) set(key []byte, value V) {
	if hasht.len == hasht.cap {
		//resize the hashtable
		newTable := New[*Node[Pair[V]]](C.ulonglong(hasht.cap * 2))
		for i := 0; i < hasht.cap*2; i++ {
			*Advanced(newTable, i) = nil
		}
		for i := 0; i < hasht.cap; i++ {
			for p_node := *Advanced(hasht.pA, i); p_node != nil; p_node = p_node.next {
				//set node to new hashtable
				hash := uint32(hasht.cap-1) & FNV32a(p_node.data.key)

				insertNode(Advanced(newTable, hash), 0, p_node)

			}
		}
		Delete(hasht.pA)
		hasht.cap *= 2
		hasht.pA = newTable

	}
	hash := uint32(hasht.cap-1) & FNV32a(key)
	for p_node := *Advanced(hasht.pA, int(hash)); p_node != nil; p_node = p_node.next {
		if bytes.Equal(p_node.data.key, key) {

			p_node.data.value = value
			return
		}
	}
	insert(Advanced(hasht.pA, hash), 0, Pair[V]{key, value})
	hasht.len++
}

func (hasht *Map[V]) get(key []byte) V {
	hash := uint32(hasht.cap-1) & FNV32a(key)
	for p_node := *Advanced(hasht.pA, hash); p_node != nil; p_node = p_node.next {
		if bytes.Equal(p_node.data.key, key) {

			return p_node.data.value
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
