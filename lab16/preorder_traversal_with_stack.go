package tree_lab

import (
	"fmt"
	"labs/structures"
	"strings"
)

func PreorderTraversalStack(root *structures.TreeNode) string {
	var res strings.Builder
	stack := structures.Stack[*structures.TreeNode]{}
	for root != nil || !stack.IsEmpty() {
		if !stack.IsEmpty() {
			root = stack.Pop()
		}
		for root != nil {
			if root.Right != nil {
				stack.Push(root.Right)
			}
			if root.Val != nil {
				res.WriteString(fmt.Sprintf("%d, ", *root.Val))
			}
			root = root.Left
		}
	}
	return strings.TrimSuffix(res.String(), ", ")
}
