package tree_lab

import "github.com/LightBulbfromSpace/Sortings_ADS/structures"

func postorderTraversalIterative(root *structures.TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	initialNode := root

	for {
		if initialNode.Left == nil && initialNode.Right == nil {
			res = append(res, *root.Val)
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

		res = append(res, *root.Val)

		if previosRoot != nil {
			if previosRoot.Left != nil {
				previosRoot.Left = nil
			} else {
				previosRoot.Right = nil
			}
		}

		root = initialNode
	}

	return res
}
