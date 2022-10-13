package radix_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	cases := []struct {
		unsortedSlice []int
		length        int
		rang          int
		sortedSlice   []int
	}{
		{[]int{111, 34, 23, 2, 57}, 3, 10, []int{2, 23, 34, 57, 111}},
		{[]int{5, 4, 3, 2, 1}, 1, 10, []int{1, 2, 3, 4, 5}},
		{[]int{15, 1, 12, 3, 4}, 2, 10, []int{1, 3, 4, 12, 15}},
		{[]int{1, 5, 3, 4}, 1, 10, []int{1, 3, 4, 5}},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d for %v", i+1, tc.unsortedSlice), func(t *testing.T) {
			got := RadixSort(tc.unsortedSlice, tc.length, tc.rang)
			assert.Equal(t, tc.sortedSlice, got)
		})
	}
}
