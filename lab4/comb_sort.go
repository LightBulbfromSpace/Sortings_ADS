package comb_sort

func CombSort(arr []float64) []float64 {
	gap := len(arr)
	length := gap
	var swapped bool
	swapped = true
	for gap >= 1 || swapped == true {
		swapped = false
		gap = int(float64(gap) / 1.3)
		for i := 0; i < length-gap; i++ {
			if arr[i] > arr[i+gap] {
				swapped = true
				swap(&arr[i], &arr[i+gap])
			}
		}
	}
	return arr
}

func swap(n1, n2 *float64) {
	*n1 = *n1 + *n2
	*n2 = *n1 - *n2
	*n1 = *n1 - *n2
}
