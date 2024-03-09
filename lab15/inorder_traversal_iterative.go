package tree_lab

import (
	"github.com/LightBulbfromSpace/Sortings_ADS/structures"
)

func inorderTraversalIterative(root *structures.TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	initialNode := root
	previousRoot := root

	for {

		if initialNode.Left == nil && initialNode.Right == nil {
			res = append(res, *root.Val)
			break
		}

		for root.Left != nil {
			previousRoot = root
			root = root.Left
		}

		res = append(res, *root.Val)

		if root.Right != nil {
			if root == initialNode {
				initialNode = root.Right
			} else {
				previousRoot.Left = root.Right
			}
		} else {
			previousRoot.Left = nil
		}
		root = initialNode
	}

	return res
}
