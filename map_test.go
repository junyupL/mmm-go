package mmm

import (
	"testing"

	"github.com/junyupL/mmm-go/memory"
)

func BenchmarkMap(b *testing.B) {
	memory.HInit()
	var hashtable = MakeMap[int](memory.HeapAllocator)
	key := "a"
	for i := 0; i < 50; i++ {
		hashtable.Set([]byte(key), i)
		key += "a"
	}
	hashtable.For(func(pData *Pair[int]) {
		pData.value *= 2
	})
	_ = hashtable.Get([]byte("aaa"))
	hashtable.Destruct()
}

func BenchmarkAVL(b *testing.B) {
	memory.HInit()
	avlTree := HeightTree[int, memory.StatelessHeapAllocator]{nil, memory.HeapAllocator}
	for i := 0; i < 50; i++ {
		avlTree.AVLInsert(i)

	}
	//avlTree.PrintTree()
	avlTree.Destruct()
}
