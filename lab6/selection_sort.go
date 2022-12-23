package selection_sort

func SelectionSort(arr []int) []int {
	current := 0
	length := len(arr)
	for i := current; i < length; i++ {
		min := arr[current]
		index := current
		for i := current + 1; i < length; i++ {
			if arr[i] < min {
				min = arr[i]
				index = i
			}
		}
		arr[current], arr[index] = arr[index], arr[current]
		current++
	}
	return arr
}
