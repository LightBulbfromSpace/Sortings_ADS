package shell_sort

func ShellSort(arr []int) []int {
	length := len(arr)
	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			tmp := arr[i]
			var j int
			for j = i; j >= gap && tmp < arr[j-gap]; j -= gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = tmp
		}
	}
	return arr
}
