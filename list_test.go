package mmm

import (
	"container/list"
	"testing"

	"github.com/junyupL/mmm-go/memory"
)

func BenchmarkGCList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var linkedl list.List

		after := linkedl.PushFront(1)
		linkedl.PushFront(4)

		linkedl.PushBack(7)
		linkedl.InsertAfter(10, after)
		for pNode := linkedl.Front(); pNode != nil; pNode = pNode.Next() {
			pNode.Value = pNode.Value.(int) + 1
			if pNode.Next() != nil {
				pNode.Next().Value = pNode.Value.(int) + pNode.Next().Value.(int)
			}
		}
	}
}

func BenchmarkENFList(b *testing.B) {
	memory.HInit()

	enf := memory.InitENF(1000)

	for i := 0; i < b.N; i++ {
		linkedl := List[int, *memory.ENFHeapAllocator]{nil, &enf}
		linkedl.PushFront(1)

		after := linkedl.Node
		linkedl.PushFront(4)
		linkedl.PushBack(7)
		linkedl.InsertAfter(10, after)

		linkedl.For(func(pNode *Node[int]) {
			pNode.Value++
			if pNode.next != nil {
				pNode.next.Value += pNode.Value
			}
		})
		linkedl.Destruct()
	}
	enf.Destruct()
}

func BenchmarkHeapList(b *testing.B) {
	memory.HInit()
	for i := 0; i < b.N; i++ {
		linkedl := List[int, memory.StatelessHeapAllocator]{nil, memory.HeapAllocator}
		linkedl.PushFront(1)

		after := linkedl.Node
		linkedl.PushFront(4)
		linkedl.PushBack(7)
		linkedl.InsertAfter(10, after)

		linkedl.For(func(pNode *Node[int]) {
			pNode.Value++
			if pNode.next != nil {
				pNode.next.Value += pNode.Value
			}
		})
		linkedl.Destruct()
	}
}

func BenchmarkMList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linkedl := List[int, memory.StatelessMAllocator]{nil, memory.MAllocator}
		linkedl.PushFront(1)

		after := linkedl.Node
		linkedl.PushFront(4)
		linkedl.PushBack(7)
		linkedl.InsertAfter(10, after)

		linkedl.For(func(pNode *Node[int]) {
			pNode.Value++
			if pNode.next != nil {
				pNode.next.Value += pNode.Value
			}
		})
		linkedl.Destruct()
	}
}
