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
		*p_list = mm_new[Node[T]]()
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
	p_node := mm_new[Node[T]]()
	p_node.data = d
	p_node.next = to_add.next
	to_add.next = p_node
	

}
func (p_node *Node[T]) destruct() {
	for curr := p_node; curr != nil; {
		temp := curr.next
		delete(curr)
		curr = temp
		
	}
}