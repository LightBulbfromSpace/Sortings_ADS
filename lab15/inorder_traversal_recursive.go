package tree_lab

import "labs/structures"

func inorderTraversalRecursive(root *structures.TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}
	res = append(res, inorderTraversalRecursive(root.Left)...)
	res = append(res, *root.Val)
	res = append(res, inorderTraversalRecursive(root.Right)...)
	return res
}
