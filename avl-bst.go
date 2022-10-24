package mmm

import (
	"github.com/junyupL/mmm-go/memory"
	"golang.org/x/exp/constraints"
)

func AVLRotate[O constraints.Ordered](pEdgeNode **HeightNode[O]) {
	if (*pEdgeNode).right.Height()-(*pEdgeNode).left.Height() == 2 {
		if (*pEdgeNode).right.right.Height() < (*pEdgeNode).right.left.Height() {
			rightRotateHeight(&(*pEdgeNode).right)
		}
		leftRotateHeight(pEdgeNode)
	}
	if (*pEdgeNode).right.Height()-(*pEdgeNode).left.Height() == -2 {
		if (*pEdgeNode).left.left.Height() < (*pEdgeNode).left.right.Height() {
			leftRotateHeight(&(*pEdgeNode).left)
		}
		rightRotateHeight(pEdgeNode)

	}
}

func (pTree *HeightTree[O, A]) AVLInsert(d O) {
	if pTree.HeightNode == nil {
		pTree.HeightNode = memory.New[HeightNode[O]](pTree.a)
		pTree.height = 1
		pTree.data = d
		pTree.left = nil
		pTree.right = nil
		return
	}
	rAVLInsert(pTree.a, pTree.HeightNode, d)
	AVLRotate(&pTree.HeightNode)
}

func rAVLInsert[O constraints.Ordered, A memory.Allocator](a A, pNode *HeightNode[O], d O) {
	if d < pNode.data {
		if pNode.left != nil {
			rAVLInsert(a, pNode.left, d)
			AVLRotate(&pNode.left)
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
			rAVLInsert(a, pNode.right, d)
			AVLRotate(&pNode.right)
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
