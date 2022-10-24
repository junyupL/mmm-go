package mmm

import (
	"bytes"
	"hash/fnv"

	"github.com/junyupL/mmm-go/memory"
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

type Map[T any, A memory.Allocator] struct {
	pA  **Node[Pair[T]]
	len int
	cap int
	a   A
}

func (hashtable Map[T, A]) For(loop func(pData *Pair[T])) {
	for i := 0; i < hashtable.cap; i++ {
		for pNode := *memory.Advanced(hashtable.pA, i); pNode != nil; pNode = pNode.next {
			loop(&pNode.Value)

		}
	}
}

func MakeMap[T any, A memory.Allocator](a A) Map[T, A] {
	hashtable := Map[T, A]{memory.NewA[*Node[Pair[T]]](a, 16), 0, 16, a}
	for i := 0; i < hashtable.cap; i++ {
		*memory.Advanced(hashtable.pA, i) = nil
	}
	return hashtable
}

func (hasht *Map[V, A]) Set(key []byte, value V) {
	if hasht.len == hasht.cap {
		//resize the hashtable
		newTable := memory.NewA[*Node[Pair[V]]](hasht.a, hasht.cap*2)
		for i := 0; i < hasht.cap*2; i++ {
			*memory.Advanced(newTable, i) = nil
		}
		for i := 0; i < hasht.cap; i++ {
			for pNode := *memory.Advanced(hasht.pA, i); pNode != nil; {
				//set node to new hashtable
				hash := uint32(hasht.cap*2-1) & FNV32a(pNode.Value.key)
				tempPNode := pNode
				pNode = pNode.next
				InsertNode(memory.Advanced(newTable, hash), 0, tempPNode)

			}
		}
		memory.Delete(hasht.a, hasht.pA)
		hasht.cap *= 2
		hasht.pA = newTable

	}
	hash := uint32(hasht.cap-1) & FNV32a(key)

	for pNode := *memory.Advanced(hasht.pA, hash); pNode != nil; pNode = pNode.next {
		if bytes.Equal(pNode.Value.key, key) {

			pNode.Value.value = value
			return
		}
	}
	PushFront(hasht.a, memory.Advanced(hasht.pA, hash), Pair[V]{key, value})
	hasht.len++
}

func (hasht *Map[V, A]) Get(key []byte) V {
	hash := uint32(hasht.cap-1) & FNV32a(key)
	for pNode := *memory.Advanced(hasht.pA, hash); pNode != nil; pNode = pNode.next {

		if bytes.Equal(pNode.Value.key, key) {

			return pNode.Value.value
		}
	}
	panic("no value for key provided")
}

func (hasht Map[V, A]) Destruct() {
	for i := 0; i < hasht.cap; i++ {
		NodeDestruct(hasht.a, *memory.Advanced(hasht.pA, i))
	}
	memory.Delete(hasht.a, hasht.pA)
}
