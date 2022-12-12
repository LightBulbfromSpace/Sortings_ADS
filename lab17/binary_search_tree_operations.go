package BST

import (
	"labs/structures"
)

func Search(node *structures.TreeNode, elem int) (*structures.TreeNode, bool) {
	if node.Val == nil {
		return nil, false
	}
	if *node.Val == elem {
		return node, true
	}

	if elem <= *node.Val {
		if node.Left == nil {
			node.Left = new(structures.TreeNode)
		}
		return Search(node.Left, elem)
	} else {
		if node.Right == nil {
			node.Right = new(structures.TreeNode)
		}
		return Search(node.Right, elem)
	}
}

func Add(node *structures.TreeNode, elem int) {
	if node.Val == nil {
		node.Val = &elem
		return
	}
	if elem <= *node.Val {
		if node.Left == nil {
			node.Left = new(structures.TreeNode)
		}
		Add(node.Left, elem)
	} else {
		if node.Right == nil {
			node.Right = new(structures.TreeNode)
		}
		Add(node.Right, elem)
	}
}

func Delete(node *structures.TreeNode, elem int) {

}