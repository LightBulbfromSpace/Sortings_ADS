package insertion_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInserctionSort(t *testing.T) {
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
			got := InsertionSort(tc.unsortedSlice)
			assert.Equal(t, tc.sortedSlice, got)
		})
	}
}

func TestInsert(t *testing.T) {
	cases := []struct {
		elem, index int
		slice       []int
		expected    []int
	}{
		{24, 0, []int{1, 2, 3, 4}, []int{24, 1, 2, 3, 4}},
		{24, 1, []int{1, 2, 3, 4}, []int{1, 24, 2, 3, 4}},
		{24, 3, []int{1, 2, 3, 4}, []int{1, 2, 3, 24, 4}},
		{24, 4, []int{1, 2, 3, 4}, []int{1, 2, 3, 4, 24}},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d for %v", i+1, tc.slice), func(t *testing.T) {
			got := insert(tc.elem, tc.index, tc.slice)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestPop(t *testing.T) {
	cases := []struct {
		index    int
		slice    []int
		expected []int
	}{
		{0, []int{1, 2, 3, 4}, []int{2, 3, 4}},
		{1, []int{1, 2, 3, 4}, []int{1, 3, 4}},
		{3, []int{1, 2, 3, 4}, []int{1, 2, 3}},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d for %v", i+1, tc.slice), func(t *testing.T) {
			got := pop(tc.index, tc.slice)
			assert.Equal(t, tc.expected, got)
		})
	}
}
