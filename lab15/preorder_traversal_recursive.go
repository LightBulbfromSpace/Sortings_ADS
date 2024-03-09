package tree_lab

import "github.com/LightBulbfromSpace/Sortings_ADS/structures"

func preorderTraversalRecursive(root *structures.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, *root.Val)
	res = append(res, preorderTraversalRecursive(root.Left)...)
	res = append(res, preorderTraversalRecursive(root.Right)...)

	return res
}
