package tree_lab

import "labs/structures"

func preorderTraversalRecusive(root *structures.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversalRecusive(root.Left)...)
	res = append(res, preorderTraversalRecusive(root.Right)...)

	return res
}
