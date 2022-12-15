package BST

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	tree_lab "labs/lab15"
	"labs/structures"
	labtest "labs/testing"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		initial string
		elem    int
		found   bool
	}{
		{"3(, 4)", 5, false},
		{"3(4, 5(6, 7))", 5, true},
		{"4(, 5)", 4, true},
		{"8(3(1, 6(4, 7)), 10(, 14(13, )))", 9, false},
		{"8(3(1, 6(4, 7)), 10(9, 14(13, )))", 13, true},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d for %v", i+1, tc.initial), func(t *testing.T) {
			got, err := tree_lab.InputTreeCLI(tc.initial)
			assert.NoError(t, err)
			resultNode := new(structures.TreeNode)
			var found bool
			labtest.FailAfter(t, 500*time.Millisecond, func() {
				resultNode, found = Search(got, tc.elem)
			})
			assert.Equal(t, tc.found, found)
			if found {
				assert.Equal(t, tc.elem, *resultNode.Val)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		initial  string
		elem     int
		expected string
	}{
		{"3(, 4)", 5, "3(, 4(, 5))"},
		{"3(, 5)", 4, "3(, 5(4, ))"},
		{"4(, 5)", 3, "4(3, 5)"},
		{"8(3(1, 6(4, 7)), 10(, 14(13, )))", 9, "8(3(1, 6(4, 7)), 10(9, 14(13, )))"},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d for %v", i+1, tc.expected), func(t *testing.T) {
			got, err := tree_lab.InputTreeCLI(tc.initial)
			assert.NoError(t, err)
			labtest.FailAfter(t, 500*time.Millisecond, func() {
				Add(got, tc.elem)
			})
			assert.Equal(t, tc.expected, tree_lab.OutputTreeCLI(got))
		})
	}
}

func TestDelete(t *testing.T) {
	cases := []struct {
		initial  string
		elem     int
		expected string
	}{
		{"3(, 4(, 5))", 5, "3(, 4(, ))"},
		{"3(, 5(4, ))", 5, "3(, 4)"},
		{"3(, 5(, 6))", 5, "3(, 6)"},
		{"4(3(1, 2), 5)", 3, "4(2(1, ), 5)"},
		{"8(4(1, 6(5, 7)), 8)", 4, "8(5(1, 6(, 7)), 8)"},
		{"8(3(1, 6(4, 7)), 10(9, 14(13, )))", 9, "8(3(1, 6(4, 7)), 10(, 14(13, )))"},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d for %v", i+1, tc.expected), func(t *testing.T) {
			got, err := tree_lab.InputTreeCLI(tc.initial)
			assert.NoError(t, err)
			labtest.FailAfter(t, 500*time.Millisecond, func() {
				Delete(got, tc.elem)
			})
			assert.Equal(t, tc.expected, tree_lab.OutputTreeCLI(got))
		})
	}
}
