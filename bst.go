package mmm

import (
	"github.com/junyupL/mmm-go/memory"
	"golang.org/x/exp/constraints"
)

type TNode[O constraints.Ordered] struct {
	data  O
	left  *TNode[O]
	right *TNode[O]
}

type Tree[O constraints.Ordered, A memory.Allocator] struct {
	*TNode[O]
	a A
}

func (pTree *Tree[O, A]) TInsert(d O) {
	if pTree == nil {
		pTree.TNode = memory.New[TNode[O]](pTree.a)
		pTree.TNode.data = d
		pTree.TNode.left = nil
		pTree.TNode.right = nil
		return
	}
	rTInsert(pTree.a, pTree.TNode, d)
}

func rTInsert[O constraints.Ordered, A memory.Allocator](a A, pNode *TNode[O], d O) {
	if d < pNode.data {
		if pNode.left != nil {
			rTInsert(a, pNode.left, d)
		} else {
			pNode.left = memory.New[TNode[O]](a)
			pNode.left.data = d
			pNode.left.left = nil
			pNode.left.right = nil
		}
	} else if d > pNode.data {
		if pNode.right != nil {
			rTInsert(a, pNode.right, d)
		} else {
			pNode.right = memory.New[TNode[O]](a)
			pNode.right.data = d
			pNode.right.left = nil
			pNode.right.right = nil
		}
	} else {
		//already exists
	}
}

// not finished
func (pTree *Tree[O, A]) TDelete(d O) {
	if pTree == nil {
		//nothing to delete
		return
	}
	if pTree.TNode.data == d {
		if pTree.TNode.left == nil {
			right := pTree.TNode.right
			memory.Delete(pTree.a, pTree.TNode)
			pTree.TNode = right
			return
		}
		if pTree.TNode.right == nil {
			left := pTree.TNode.left
			memory.Delete(pTree.a, pTree.TNode)
			pTree.TNode = left
			return
		}

		if pTree.TNode.right.left == nil {
			parent := pTree.TNode
			pNode := pTree.TNode.right

			parent.data = pNode.data
			parent.right = pNode.right
			memory.Delete(pTree.a, pNode)
			return
		}
		parent := pTree.TNode.right
		pNode := pTree.TNode.right.left
		for ; pNode.left != nil; pNode = pNode.left {
			parent = pNode
		}

		pTree.TNode.data = pNode.data
		parent.left = pNode.right
		memory.Delete(pTree.a, pNode)

		return
	}

	if d > pTree.TNode.data {
		rTDelete(pTree.a, pTree.TNode.right, d)
		return
	}
	rTDelete(pTree.a, pTree.TNode.left, d)
}

func rTDelete[O constraints.Ordered, A memory.Allocator](a A, pNode *TNode[O], d O) {

}

func leftRotate[O constraints.Ordered](pEdgeNode **TNode[O]) {
	right := (*pEdgeNode).right
	(*pEdgeNode).right = right.left
	right.left = *pEdgeNode
	*pEdgeNode = right
}

func rightRotate[O constraints.Ordered](pEdgeNode **TNode[O]) {
	left := (*pEdgeNode).left
	(*pEdgeNode).left = left.right
	left.right = *pEdgeNode
	*pEdgeNode = left
}
