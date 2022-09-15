package tree_lab

import (
	"labs/structures"
)

func inorderTraversalIterative(root *structures.TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	initialNode := root
	previosRoot := root

	for true {

		if initialNode.Left == nil && initialNode.Right == nil {
			res = append(res, root.Val)
			break
		}

		for root.Left != nil {
			previosRoot = root
			root = root.Left
		}

		res = append(res, root.Val)

		if root.Right != nil {
			if root == initialNode {
				initialNode = root.Right
			} else {
				previosRoot.Left = root.Right
			}
		} else {
			previosRoot.Left = nil
		}
		root = initialNode
	}

	return res
}
