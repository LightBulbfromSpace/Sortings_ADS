package radix_sort

import (
	"math"
)

func RadixSort(input []int, length, rang int) []int {
	for i := 1; i <= length; i++ {
		lists := make([][]int, rang)
		for _, num := range input {
			j := num % int(math.Pow10(i)) / int(math.Pow10(i-1))
			lists[j] = append(lists[j], num)
		}
		input = input[:0]
		for _, list := range lists {
			input = append(input, list...)
		}
	}
	return input
}
