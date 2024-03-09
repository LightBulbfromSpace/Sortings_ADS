package tree_lab

import (
	"fmt"
	"github.com/LightBulbfromSpace/Sortings_ADS/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

//var Tree = &structures.TreeNode{
//	Val: 8,
//	Left: &structures.TreeNode{
//		Val:  3,
//		Left: &structures.TreeNode{Val: 1},
//		Right: &structures.TreeNode{
//			Val:   6,
//			Left:  &structures.TreeNode{Val: 4},
//			Right: &structures.TreeNode{Val: 7},
//		},
//	},
//	Right: &structures.TreeNode{
//		Val: 10,
//		Right: &structures.TreeNode{
//			Val:  14,
//			Left: &structures.TreeNode{Val: 13},
//		},
//	},
//}
//var Tree2 = &structures.TreeNode{
//	Val: 5,
//}
//var Tree3 = &structures.TreeNode{
//	Val:   5,
//	Left:  &structures.TreeNode{Val: 4},
//	Right: &structures.TreeNode{Val: 3},
//}

func TestPreorderTraversalStack(t *testing.T) {
	cases := []struct {
		tree     string
		expected string
	}{
		{"5", "5"},
		{"5(4, 3)", "5, 4, 3"},
		{"8(3(1, 6(4, 7)), 10(, 14(13, )))", "8, 3, 1, 6, 4, 7, 10, 14, 13"},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d for %v", i+1, tc.expected), func(t *testing.T) {
			tree, _ := structures.InputTreeCLI(tc.tree)
			got := PreorderTraversalStack(tree)
			assert.Equal(t, tc.expected, got)
		})
	}
}
