package tree_lab

import "github.com/LightBulbfromSpace/Sortings_ADS/structures"

func preorderTraversalIterative(root *structures.TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	initialNode := root

	for {
		if initialNode.Left == nil && initialNode.Right == nil {
			res = append([]int{*root.Val}, res...)
			break
		}

		var previousRoot *structures.TreeNode

		for root.Left != nil || root.Right != nil {
			previousRoot = root
			if root.Right != nil {
				root = root.Right
			} else {
				root = root.Left
			}

		}

		res = append([]int{*root.Val}, res...)

		if previousRoot != nil {
			if previousRoot.Right != nil {
				previousRoot.Right = nil
			} else {
				previousRoot.Left = nil
			}
		}

		root = initialNode
	}

	return res
}
