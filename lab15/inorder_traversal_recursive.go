package tree_lab

import "github.com/LightBulbfromSpace/Sortings_ADS/structures"

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
