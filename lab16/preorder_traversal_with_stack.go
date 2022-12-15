package tree_lab

import (
	"fmt"
	"labs/structures"
	"strings"
)

// PreorderTraversalStack - нерекурсивный обход дерева с помощью стека.
// В стек складываются правые узлы. Алгоритм идет к самому левому узлу,
// обратывая все проходимые узлы. Как только достугнут нулевой узел,
// достается узел из стека. Все повторяется до тех пор, пока не будет достигнут
// самый правый узел, то есть пока стек не пуст и обрабатываемый узел ненулевой.
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
