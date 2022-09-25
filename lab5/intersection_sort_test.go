package insertion_sort

import (
	"fmt"
	labtest "labs"
	"reflect"
	"testing"
)

/*func TestInserctionSort(t *testing.T) {
	cases := []struct {
		unsortedSlice *[]int
		sortedSlice   *[]int
	}{
		{&[]int{1, 2, 3, 4}, &[]int{1, 2, 3, 4}},
		{&[]int{-1, 4, 23, 12}, &[]int{-1, 4, 12, 23}},
		{&[]int{5, 4, 3, 2, 1}, &[]int{1, 2, 3, 4, 5}},
		{&[]int{1, 5, 3, 4}, &[]int{1, 3, 4, 5}},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {
			got, err := InsertionSort(tc.unsortedSlice)
			labtest.AssertNoError(t, err)
			if !reflect.DeepEqual(tc.sortedSlice, got) {
				t.Errorf("got %v, but want %v", got, tc.sortedSlice)
			}
		})
	}
}*/

func TestShiftElements(t *testing.T) {
	cases := []struct {
		initialSlice  *[]int
		offset        int
		index         int
		expectedSlice *[]int
	}{
		{&[]int{1, 2, 3, 4, 5}, 1, 2, &[]int{1, 2, 0, 3, 4}},
		{&[]int{-1, 4, 23, 12, 15, 17, 18}, 2, 1, &[]int{-1, 0, 0, 4, 23, 12, 15}},
		{&[]int{5, 4, 3, 2, 1}, 4, 3, &[]int{5, 4, 3, 0, 0}},
		{&[]int{1, 5, 3, 4}, 0, 2, &[]int{1, 5, 3, 4}},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {
			err := shiftElements(tc.initialSlice, tc.index, tc.offset)
			labtest.AssertNoError(t, err)
			if !reflect.DeepEqual(tc.expectedSlice, tc.initialSlice) {
				t.Errorf("got %v, but want %v", tc.initialSlice, tc.expectedSlice)
			}
		})
	}
}
