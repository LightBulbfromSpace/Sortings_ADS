package tree_lab

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"labs/structures"
	"testing"
)

var Tree = &structures.TreeNode{
	Val: 8,
	Left: &structures.TreeNode{
		Val:  3,
		Left: &structures.TreeNode{Val: 1},
		Right: &structures.TreeNode{
			Val:   6,
			Left:  &structures.TreeNode{Val: 4},
			Right: &structures.TreeNode{Val: 7},
		},
	},
	Right: &structures.TreeNode{
		Val: 10,
		Right: &structures.TreeNode{
			Val:  14,
			Left: &structures.TreeNode{Val: 13},
		},
	},
}
var Tree2 = &structures.TreeNode{
	Val: 5,
}
var Tree3 = &structures.TreeNode{
	Val:   5,
	Left:  &structures.TreeNode{Val: 4},
	Right: &structures.TreeNode{Val: 3},
}

func TestPreorderTraversalStack(t *testing.T) {
	cases := []struct {
		tree     *structures.TreeNode
		expected string
	}{
		{Tree2, "5"},
		{Tree3, "5, 4, 3"},
		{Tree, "8, 3, 1, 6, 4, 7, 10, 14, 13"},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d for %v", i+1, tc.expected), func(t *testing.T) {
			got := PreorderTraversalStack(tc.tree)
			assert.Equal(t, tc.expected, got)
		})
	}
}
