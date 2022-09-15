package tree_lab

import "labs/structures"

func postorderTraversalIteravite(root *structures.TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	initialNode := root

	for true {
		if initialNode.Left == nil && initialNode.Right == nil {
			res = append(res, root.Val)
			break
		}

		var previosRoot *structures.TreeNode

		for root.Left != nil || root.Right != nil {
			previosRoot = root
			if root.Left != nil {
				root = root.Left
			} else {
				root = root.Right
			}

		}

		res = append(res, root.Val)

		if previosRoot.Left != nil {
			previosRoot.Left = nil
		} else {
			previosRoot.Right = nil
		}

		root = initialNode
	}

	return res
}
