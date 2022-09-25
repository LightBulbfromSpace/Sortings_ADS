package comb_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombSort(t *testing.T) {
	cases := []struct {
		arr, expected []float64
	}{
		{
			[]float64{1, 3, 2},
			[]float64{1, 2, 3},
		},
		{
			[]float64{-5, -46, -7, 0, 10, 5, -72, -68, 30, 78, 21},
			[]float64{-72, -68, -46, -7, -5, 0, 5, 10, 21, 30, 78},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("test for %v", tc.arr), func(t *testing.T) {
			got := CombSort(tc.arr)
			assert.Equal(t, tc.expected, got)
		})
	}
}
