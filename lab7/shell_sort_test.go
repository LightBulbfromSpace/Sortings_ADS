package shell_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	cases := []struct {
		unsortedSlice []int
		sortedSlice   []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{-1, 4, 23, 12}, []int{-1, 4, 12, 23}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 1, 2, 3, 4}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 5, 3, 4}, []int{1, 3, 4, 5}},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d for %v", i+1, tc.unsortedSlice), func(t *testing.T) {
			got := ShellSort(tc.unsortedSlice)
			assert.Equal(t, tc.sortedSlice, got)
		})
	}
}
