package tree_lab

import "github.com/LightBulbfromSpace/Sortings_ADS/structures"

func postorderTraversalRecursive(root *structures.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(res, postorderTraversalRecursive(root.Left)...)
	res = append(res, postorderTraversalRecursive(root.Right)...)
	res = append(res, *root.Val)

	return res
}
