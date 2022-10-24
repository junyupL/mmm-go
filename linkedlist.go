package mmm

import "github.com/junyupL/mmm-go/memory"

type Node[T any] struct {
	Value T
	next  *Node[T]
}

type List[T any, A memory.Allocator] struct {
	*Node[T]
	a A
}

func (pNode *Node[T]) For(loop func(pNode *Node[T])) {
	for ; pNode != nil; pNode = pNode.next {
		loop(pNode)
	}
}

func (pList *List[T, A]) Insert(index int, d T) {
	if index < 0 {
		//index out of bounds
		return
	}

	if index == 0 {
		temp := pList.Node
		pList.Node = memory.New[Node[T]](pList.a)
		pList.Node.Value = d
		pList.Node.next = temp
		return
	}
	//pointer to node to add the new node after
	var toAdd *Node[T] = pList.Node
	if toAdd == nil {
		//index out of bounds
		return
	}
	for i := 1; i < index; i++ {
		toAdd = toAdd.next
		if toAdd == nil {
			//index out of bounds
			return
		}
	}
	pNode := memory.New[Node[T]](pList.a)
	pNode.Value = d
	pNode.next = toAdd.next
	toAdd.next = pNode

}

func (pList *List[T, A]) InsertAfter(d T, after *Node[T]) {

	if pList.Node == nil {
		if after == nil {
			pList.Node = memory.New[Node[T]](pList.a)
			pList.Node.Value = d
			pList.Node.next = nil
			return
		} else {
			//error
			return
		}
	}
	//pointer to node to add the new node after
	toAdd := pList.Node
	for ; toAdd.next != nil; toAdd = toAdd.next {
		if toAdd == after {
			break
		}
	}
	if toAdd != after {
		//error
		return
	}
	toAdd.next = memory.New[Node[T]](pList.a)
	toAdd.next.Value = d
	toAdd.next.next = nil

}

func (pList *List[T, A]) PushBack(d T) {

	if pList.Node == nil {
		pList.Node = memory.New[Node[T]](pList.a)
		pList.Node.Value = d
		pList.Node.next = nil
		return
	}
	//pointer to node to add the new node after
	toAdd := pList.Node
	for ; toAdd.next != nil; toAdd = toAdd.next {
	}

	toAdd.next = memory.New[Node[T]](pList.a)
	toAdd.next.Value = d
	toAdd.next.next = nil

}

func (pList *List[T, A]) PushFront(d T) {
	temp := pList.Node
	pList.Node = memory.New[Node[T]](pList.a)
	pList.Node.Value = d
	pList.Node.next = temp
	return

}

func PushFront[T any, A memory.Allocator](a A, pList **Node[T], d T) {
	temp := *pList
	*pList = memory.New[Node[T]](a)
	(*pList).Value = d
	(*pList).next = temp
	return

}

func InsertNode[T any](pList **Node[T], index int, pNode *Node[T]) {
	if index < 0 {
		//index out of bounds
		return
	}

	if index == 0 {
		temp := *pList
		*pList = pNode
		(*pList).next = temp
		return
	}
	//pointer to node to add the new node after
	var toAdd *Node[T] = *pList
	if toAdd == nil {
		//index out of bounds
		return
	}
	for i := 1; i < index; i++ {
		toAdd = toAdd.next
		if toAdd == nil {
			//index out of bounds
			return
		}
	}
	pNode.next = toAdd.next
	toAdd.next = pNode

}

func (pList List[T, A]) Destruct() {
	for curr := pList.Node; curr != nil; {
		temp := curr.next
		memory.Delete(pList.a, curr)
		curr = temp

	}
}

func NodeDestruct[T any, A memory.Allocator](a A, pList *Node[T]) {
	for curr := pList; curr != nil; {
		temp := curr.next
		memory.Delete(a, curr)
		curr = temp

	}
}
