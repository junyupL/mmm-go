package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type TNode[O constraints.Ordered] struct {
	data  O
	left  *TNode[O]
	right *TNode[O]
}

func TInsert[O constraints.Ordered](pTree **TNode[O], d O) {
	if pTree == nil {
		*pTree = New[TNode[O]]()
		(*pTree).data = d
		(*pTree).left = nil
		(*pTree).right = nil
		return
	}
	rTInsert(*pTree, d)
}

func rTInsert[O constraints.Ordered](pNode *TNode[O], d O) {
	if d < pNode.data {
		if pNode.left != nil {
			rTInsert(pNode.left, d)
		} else {
			pNode.left = New[TNode[O]]()
			pNode.left.data = d
			pNode.left.left = nil
			pNode.left.right = nil
		}
	} else if d > pNode.data {
		if pNode.right != nil {
			rTInsert(pNode.right, d)
		} else {
			pNode.right = New[TNode[O]]()
			pNode.right.data = d
			pNode.right.left = nil
			pNode.right.right = nil
		}
	} else {
		//already exists
	}
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

type HeightNode[O constraints.Ordered] struct {
	height int
	data   O
	left   *HeightNode[O]
	right  *HeightNode[O]
}

func (pNode *HeightNode[O]) PrintTree() {

	if pNode != nil {
		pNode.left.PrintTree()
		fmt.Print(pNode.height, ":", pNode.data, "    ")
		pNode.right.PrintTree()
	}
}
func (pNode *HeightNode[O]) Destruct() {
	if pNode != nil {
		pNode.left.Destruct()
		pNode.right.Destruct()
		Delete(pNode)
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

func HeightInsert[O constraints.Ordered](pTree **HeightNode[O], d O) {
	if pTree == nil {
		*pTree = New[HeightNode[O]]()
		(*pTree).height = 1
		(*pTree).data = d
		(*pTree).left = nil
		(*pTree).right = nil
		return
	}
	rHeightInsert(*pTree, d)
}

func rHeightInsert[O constraints.Ordered](pNode *HeightNode[O], d O) {
	if d < pNode.data {
		if pNode.left != nil {
			rHeightInsert(pNode.left, d)
			pNode.height = 1 + Max(pNode.left.Height(), pNode.right.Height())
			return
		}

		pNode.left = New[HeightNode[O]]()
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
			rHeightInsert(pNode.right, d)
			pNode.height = 1 + Max(pNode.left.Height(), pNode.right.Height())
			return
		}
		pNode.right = New[HeightNode[O]]()
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

func AVLInsert[O constraints.Ordered](pTree **HeightNode[O], d O) {
	if *pTree == nil {
		*pTree = New[HeightNode[O]]()
		(*pTree).height = 1
		(*pTree).data = d
		(*pTree).left = nil
		(*pTree).right = nil
		return
	}
	rAVLInsert(*pTree, d)
	AVLRotate(pTree)
}

func rAVLInsert[O constraints.Ordered](pNode *HeightNode[O], d O) {
	if d < pNode.data {
		if pNode.left != nil {
			rAVLInsert(pNode.left, d)
			AVLRotate(&pNode.left)
			pNode.height = 1 + Max(pNode.left.Height(), pNode.right.Height())
			return
		}

		pNode.left = New[HeightNode[O]]()
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
			rAVLInsert(pNode.right, d)
			AVLRotate(&pNode.right)
			pNode.height = 1 + Max(pNode.left.Height(), pNode.right.Height())
			return
		}
		pNode.right = New[HeightNode[O]]()
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
