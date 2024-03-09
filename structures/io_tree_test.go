package structures

import (
	"fmt"
	labtest "github.com/LightBulbfromSpace/Sortings_ADS/testing"
	"github.com/stretchr/testify/assert"
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
			got := &TreeNode{}
			var err error
			labtest.FailAfter(t, 100*time.Microsecond, func() {
				got, err = InputTreeCLI(tc.exp)
			})
			assert.NoError(t, err)

			gotExp := OutputTreeCLI(got)
			assert.Equal(t, tc.exp, gotExp)
		})
	}
}

func TestOutputTreeCLI(t *testing.T) {
	values := []int{8, 3, 1, 6, 4, 7, 10, 14, 13, 5}
	var Tree = &TreeNode{
		Val: &values[0],
		Left: &TreeNode{
			Val:  &values[1],
			Left: &TreeNode{Val: &values[2]},
			Right: &TreeNode{
				Val:   &values[3],
				Left:  &TreeNode{Val: &values[4]},
				Right: &TreeNode{Val: &values[5]},
			},
		},
		Right: &TreeNode{
			Val: &values[6],
			Right: &TreeNode{
				Val:  &values[7],
				Left: &TreeNode{Val: &values[8]},
			},
		},
	}

	Tree2 := &TreeNode{
		Val: &values[9],
	}

	var Tree3 = &TreeNode{
		Val:   &values[9],
		Left:  &TreeNode{Val: &values[1]},
		Right: &TreeNode{Val: &values[4]},
	}

	cases := []struct {
		exp  string
		tree *TreeNode
	}{
		{"8(3(1, 6(4, 7)), 10(, 14(13, )))", Tree},
		{"5", Tree2},
		{"5(3, 4)", Tree3},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d for %v", i+1, tc.exp), func(t *testing.T) {
			var got string
			labtest.FailAfter(t, 20*time.Microsecond, func() {
				got = OutputTreeCLI(tc.tree)
			})
			assert.Equal(t, tc.exp, got)
		})
	}
}
