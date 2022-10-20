package main

type Node[T any] struct {
	data T
	next *Node[T]
}

func insert[T any](pList **Node[T], index int, d T) {
	if index < 0 {
		//index out of bounds
		return
	}

	if index == 0 {
		temp := *pList
		*pList = New[Node[T]]()
		(*pList).data = d
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
	pNode := New[Node[T]]()
	pNode.data = d
	pNode.next = toAdd.next
	toAdd.next = pNode

}

func insertNode[T any](pList **Node[T], index int, pNode *Node[T]) {
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

func (pNode *Node[T]) Destruct() {
	for curr := pNode; curr != nil; {
		temp := curr.next
		Delete(curr)
		curr = temp

	}
}
