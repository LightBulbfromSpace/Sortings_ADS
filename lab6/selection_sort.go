package selection_sort

func SelectionSort(arr []int) []int {
	current := 0
	lenght := len(arr)
	for i := current; i < lenght; i++ {
		min := arr[current]
		index := current
		for i := current + 1; i < lenght; i++ {
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
