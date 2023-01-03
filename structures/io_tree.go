package structures

import (
	"regexp"
	"strconv"
	"strings"
)

func InputTreeCLI(exp string) (*TreeNode, error) {
	stack := Stack[*TreeNode]{}
	currentNode := new(TreeNode)
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
			currentNode.Right = new(TreeNode)
			stack.Push(currentNode.Right)
			currentNode.Left = new(TreeNode)
			currentNode = currentNode.Left

		}
		if exp[0] == ',' {
			currentNode, _ = stack.Pop()
		}
		exp = exp[1:]
	}
	return root, nil
}

func OutputTreeCLI(node *TreeNode) string {
	return strings.ReplaceAll(strings.Join(outputTreeRecursion(node, []string{}), ""), "(, )", "")
}

func outputTreeRecursion(node *TreeNode, exp []string) []string {
	if node.Val != nil {
		exp = append(exp, strconv.Itoa(*node.Val))
	}
	exp = append(exp, "(")
	if node.Left != nil {
		exp = outputTreeRecursion(node.Left, exp)
	}

	exp = append(exp, ", ")
	if node.Right != nil {
		exp = outputTreeRecursion(node.Right, exp)
	}
	exp = append(exp, ")")
	return exp
}
