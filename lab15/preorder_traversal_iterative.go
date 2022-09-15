package tree_lab

import "labs/structures"

func preorderTraversalIteravite(root *structures.TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	initialNode := root

	for true {
		if initialNode.Left == nil && initialNode.Right == nil {
			res = append([]int{root.Val}, res...)
			break
		}

		var previosRoot *structures.TreeNode

		for root.Left != nil || root.Right != nil {
			previosRoot = root
			if root.Right != nil {
				root = root.Right
			} else {
				root = root.Left
			}

		}

		res = append([]int{root.Val}, res...)

		if previosRoot.Right != nil {
			previosRoot.Right = nil
		} else {
			previosRoot.Left = nil
		}

		root = initialNode
	}

	return res
}
