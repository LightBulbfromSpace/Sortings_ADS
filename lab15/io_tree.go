package tree_lab

import (
	"labs/structures"
	"regexp"
	"strconv"
	"strings"
)

func InputTreeCLI(exp string) (*structures.TreeNode, error) {
	stack := structures.Stack[*structures.TreeNode]{}
	currentNode := new(structures.TreeNode)
	root := currentNode
	regex, err := regexp.Compile(`^\d+`)
	if err != nil {
		return nil, err
	}
	for len(exp) > 0 {
		if regex.Match([]byte(exp)) {
			res := regex.Find([]byte(exp))
			val, err := strconv.Atoi(string(res))
			if err != nil {
				return nil, err
			}
			currentNode.Val = &val
			exp = exp[len(res):]
			continue
		}
		if exp[0] == '(' {
			currentNode.Right = new(structures.TreeNode)
			stack.Push(currentNode.Right)
			currentNode.Left = new(structures.TreeNode)
			currentNode = currentNode.Left

		}
		if exp[0] == ',' {
			currentNode = stack.Pop()
		}
		exp = exp[1:]
	}
	return root, nil
}

func OutputTreeCLI(node *structures.TreeNode) string {
	return strings.ReplaceAll(strings.Join(OutputTreeRecursion(node, []string{}), ""), "(, )", "")
}

func OutputTreeRecursion(node *structures.TreeNode, exp []string) []string {
	if node.Val != nil {
		exp = append(exp, strconv.Itoa(*node.Val))
	}
	exp = append(exp, "(")
	if node.Left != nil {
		exp = OutputTreeRecursion(node.Left, exp)
	}

	exp = append(exp, ", ")
	if node.Right != nil {
		exp = OutputTreeRecursion(node.Right, exp)
	}
	exp = append(exp, ")")
	return exp
}
