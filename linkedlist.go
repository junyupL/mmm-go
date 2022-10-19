package main

type Node[T any] struct {
	data T
	next *Node[T]
}

func insert[T any](p_list **Node[T], index int, d T) {
	if index < 0 {
		//index out of bounds
		return
	}

	if index == 0 {
		temp := *p_list
		*p_list = New[Node[T]]()
		(*p_list).data = d
		(*p_list).next = temp
		return
	}
	//pointer to node to add the new node after
	var to_add *Node[T] = *p_list
	if to_add == nil {
		//index out of bounds
		return
	}
	for i := 1; i < index; i++ {
		to_add = to_add.next
		if to_add == nil {
			//index out of bounds
			return
		}
	}
	p_node := New[Node[T]]()
	p_node.data = d
	p_node.next = to_add.next
	to_add.next = p_node

}

func insertNode[T any](p_list **Node[T], index int, p_node *Node[T]) {
	if index < 0 {
		//index out of bounds
		return
	}

	if index == 0 {
		temp := *p_list
		*p_list = p_node
		(*p_list).next = temp
		return
	}
	//pointer to node to add the new node after
	var to_add *Node[T] = *p_list
	if to_add == nil {
		//index out of bounds
		return
	}
	for i := 1; i < index; i++ {
		to_add = to_add.next
		if to_add == nil {
			//index out of bounds
			return
		}
	}
	p_node.next = to_add.next
	to_add.next = p_node

}

func (p_node *Node[T]) Destruct() {
	for curr := p_node; curr != nil; {
		temp := curr.next
		Delete(curr)
		curr = temp

	}
}
