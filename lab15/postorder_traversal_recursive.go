package tree_lab

import "labs/structures"

func postorderTraversalRecusive(root *structures.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(res, postorderTraversalRecusive(root.Left)...)
	res = append(res, postorderTraversalRecusive(root.Right)...)
	res = append(res, root.Val)

	return res
}
