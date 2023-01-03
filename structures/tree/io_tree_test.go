package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"labs/structures"
	labtest "labs/testing"
	"testing"
	"time"
)

func TestInputTreeCLI(t *testing.T) {

	cases := []struct {
		exp string
	}{
		{""},
		{"5"},
		{"5(4, 3)"},
		{"8(30(1, 6(4, 7)), 10(, 14(13, )))"},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d for %v", i+1, tc.exp), func(t *testing.T) {
			got := &structures.TreeNode{}
			var err error
			labtest.FailAfter(t, 100*time.Microsecond, func() {
				got, err = InputTreeCLI(tc.exp)
			})
			assert.NoError(t, err)

			gotExp := OutputTreeCLI(got)
			assert.Equal(t, tc.exp, gotExp)
			fmt.Println(gotExp)
		})
	}
}

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

//func TestOutputTreeCLI(t *testing.T) {
//	cases := []struct {
//		exp  string
//		tree *structures.TreeNode
//	}{
//		{"8(3(1, 6(4, 7)), 10(, 14(13, )))", Tree},
//	}
//	for i, tc := range cases {
//		t.Run(fmt.Sprintf("test #%d for %v", i+1, tc.exp), func(t *testing.T) {
//			var got string
//			failAfter(t, 20*time.Microsecond, func() {
//				got = OutputTreeCLI(tc.tree)
//			})
//			assert.Equal(t, tc.exp, got)
//		})
//	}
//}
