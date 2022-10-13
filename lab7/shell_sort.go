package shell_sort

func ShellSort(arr []int) []int {
	lenght := len(arr)
	for gap := lenght / 2; gap > 0; gap /= 2 {
		for i := gap; i < lenght; i++ {
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
