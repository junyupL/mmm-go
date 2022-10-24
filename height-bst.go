package mmm

import (
	"fmt"

	"github.com/junyupL/mmm-go/memory"
	"golang.org/x/exp/constraints"
)

type HeightNode[O constraints.Ordered] struct {
	height int
	data   O
	left   *HeightNode[O]
	right  *HeightNode[O]
}

type HeightTree[O constraints.Ordered, A memory.Allocator] struct {
	*HeightNode[O]
	a A
}

func (pNode *HeightNode[O]) PrintTree() {

	if pNode != nil {
		pNode.left.PrintTree()
		fmt.Print(pNode.height, ":", pNode.data, "    ")
		pNode.right.PrintTree()
	}
}
func (tree HeightTree[O, A]) Destruct() {
	if tree.HeightNode != nil {
		Destruct(tree.a, tree.HeightNode.left)
		Destruct(tree.a, tree.HeightNode.right)
		memory.Delete(tree.a, tree.HeightNode)
	}
}

func Destruct[O constraints.Ordered, A memory.Allocator](a A, pNode *HeightNode[O]) {
	if pNode != nil {
		Destruct(a, pNode.left)
		Destruct(a, pNode.right)
		memory.Delete(a, pNode)
	}
}

func (pNode *HeightNode[O]) Height() int {
	if pNode == nil {
		return 0
	}
	return pNode.height
}

func leftRotateHeight[O constraints.Ordered](pEdgeNode **HeightNode[O]) {
	right := (*pEdgeNode).right
	(*pEdgeNode).right = right.left
	(*pEdgeNode).height = 1 + Max(right.left.Height(), (*pEdgeNode).left.Height())

	right.left = *pEdgeNode
	right.height = 1 + Max(right.left.Height(), right.right.Height())
	//
	*pEdgeNode = right
}

func rightRotateHeight[O constraints.Ordered](pEdgeNode **HeightNode[O]) {
	left := (*pEdgeNode).left
	(*pEdgeNode).left = left.right
	(*pEdgeNode).height = 1 + Max(left.right.Height(), (*pEdgeNode).right.Height())

	left.right = *pEdgeNode
	left.height = 1 + Max(left.right.Height(), left.left.Height())
	//
	*pEdgeNode = left
}

func (pTree *HeightTree[O, A]) Insert(d O) {
	if pTree == nil {
		pTree.HeightNode = memory.New[HeightNode[O]](pTree.a)
		pTree.height = 1
		pTree.data = d
		pTree.left = nil
		pTree.right = nil
		return
	}
	rHeightInsert(pTree.a, pTree.HeightNode, d)
}

func rHeightInsert[O constraints.Ordered, A memory.Allocator](a A, pNode *HeightNode[O], d O) {
	if d < pNode.data {
		if pNode.left != nil {
			rHeightInsert(a, pNode.left, d)
			pNode.height = 1 + Max(pNode.left.Height(), pNode.right.Height())
			return
		}

		pNode.left = memory.New[HeightNode[O]](a)
		pNode.left.height = 1
		pNode.left.data = d
		pNode.left.left = nil
		pNode.left.right = nil
		if pNode.right == nil {
			pNode.height = 2
		}
		return
	}
	if d > pNode.data {
		if pNode.right != nil {
			rHeightInsert(a, pNode.right, d)
			pNode.height = 1 + Max(pNode.left.Height(), pNode.right.Height())
			return
		}
		pNode.right = memory.New[HeightNode[O]](a)
		pNode.right.height = 1
		pNode.right.data = d
		pNode.right.left = nil
		pNode.right.right = nil
		if pNode.left == nil {
			pNode.height = 2
		}
		return
	}
	//already exists
}
