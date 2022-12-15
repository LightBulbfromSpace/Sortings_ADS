package insertion_sort

func InsertionSort(arr []int) []int {
	lenght := len(arr)
	for i := 1; i < lenght; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
			j--
		}
	}
	return arr
}

func insert(elem, index int, arr []int) []int {
	arr = append(arr, 0)
	for i := len(arr) - 1; i > index; i-- {
		arr[i] = arr[i-1]
	}
	arr[index] = elem
	return arr
}

func pop(index int, arr []int) []int {
	newLength := len(arr) - 1
	for i := index; i < newLength; i++ {
		arr[i] = arr[i+1]
	}
	return arr[:newLength]
}
