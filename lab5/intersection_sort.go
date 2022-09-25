package insertion_sort

import "fmt"

func InsertionSort(arr *[]int) (*[]int, error) {
	lth := len(*arr)

	for i := 1; i < lth; i++ {
		if (*arr)[i-1] < (*arr)[i] {
			continue
		}
		err := findPlace(arr, i)
		if err != nil {
			return nil, err
		}
	}
	return arr, nil
}

func findPlace(arr *[]int, index int) error {
	for i := 0; i < index; i++ {
		if (*arr)[i] > (*arr)[index] {
			tmp := (*arr)[index]
			err := shiftElements(arr, i, 1)
			if err != nil {
				return err
			}
			(*arr)[i] = tmp
		}
	}
	return nil
}

func shiftElements(arr *[]int, index, offset int) error { // NEED TO BE FINISHED!!!
	if (index+offset) > len(*arr) || index > len(*arr) {
		return fmt.Errorf("index out of range")
	}
	buff := 0
	for i := index; i < len(*arr); i++ {
		buff = (*arr)[i]
		if i+offset < len(*arr) {
			(*arr)[i+offset] = buff
		}
		if i+offset+1 < len(*arr) {
			buff = (*arr)[i+offset+1]
		}
	}
	return nil
}
