package external_polyphase_sort

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func ExternalPolyphaseSort(initialFile *os.File, maxSeqInRAM int) (*os.File, error) {

	var tmpFiles []*os.File

	// divide information from initial file to several files
	var err error
	for err != io.EOF {
		var nums *[]int
		nums, err = readMaxAmountOfNums(initialFile, maxSeqInRAM)
		if err != nil && err != io.EOF {
			return nil, err
		}
		QuickSort(nums)

		f, err := createTemporaryFile()
		if err != nil {
			return nil, err
		}

		_, err = f.Write(intsToBytesSlice(*nums...))
		if err != nil {
			return nil, err
		}
		f.Seek(0, 0)
		tmpFiles = append(tmpFiles, f)
	}

	numberOfTmpFiles := len(tmpFiles)
	for numberOfTmpFiles > 1 {
		var newFiles []*os.File
		for i := 0; i < (numberOfTmpFiles >> 1 << 1); i += 2 {
			newFile, err := mergeSortedSequencesInFiles(tmpFiles[i], tmpFiles[i+1])
			if err != nil {
				return nil, err
			}

			newFiles = append(newFiles, newFile)
		}

		if (numberOfTmpFiles & 1) == 1 {
			newFile, err := createTemporaryFile()
			if err != nil {
				return nil, err
			}

			lastFile := tmpFiles[len(tmpFiles)-1]
			fileInfo, err := lastFile.Stat()
			if err != nil {
				return nil, err
			}

			buffer := make([]byte, fileInfo.Size())
			_, err = lastFile.Read(buffer)
			if err != nil {
				return nil, err
			}

			fmt.Println(buffer)

			_, err = newFile.Write(buffer)
			if err != nil {
				return nil, err
			}
			newFile.Seek(0, 0)
		}

		err = deleteFiles(tmpFiles)
		if err != nil {
			return nil, err
		}
		tmpFiles = newFiles
		numberOfTmpFiles = (numberOfTmpFiles & 1) + (numberOfTmpFiles >> 1)
	}
	return tmpFiles[0], nil
}

func deleteFiles(files []*os.File) error {
	for _, f := range files {
		err := os.Remove(f.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

func createNFiles(n int) ([]*os.File, error) {
	var files []*os.File
	for i := 0; i < n; i++ {
		tmp, err := os.Create(createTmpFileName())
		if err != nil {
			return nil, err
		}
		files = append(files, tmp)
	}
	return files, nil
}

func createTemporaryFile() (*os.File, error) {
	fileName := createTmpFileName()
	tmp, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

// spaces should be taken in count in maxLength
func readMaxAmountOfNums(file *os.File, maxLength int) (*[]int, error) {
	var (
		err   error
		nums  []int
		n, nc int
	)
	for n < maxLength {
		var num *int
		num, nc, err = getNum(file)
		n += nc
		if nc > maxLength {
			return nil, errors.New("not enough memory: numbers in file are too big")
		}
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			if nc > 0 {
				nums = append(nums, *num)
			}
			return &nums, err
		}
		if n > maxLength {
			file.Seek(-int64(nc), 1)
			break
		}
		if num != nil {
			nums = append(nums, *num)
		}
	}

	return &nums, nil
}

func createTmpFileName() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("tmp_%d.dat", rand.Int63())
}

func mergeSortedSequencesInFiles(inputFile1, inputFile2 *os.File) (*os.File, error) {
	inputFile1.Seek(0, 0)
	inputFile2.Seek(0, 0)
	outputFile, err := createTemporaryFile()
	if err != nil {
		return nil, err
	}
	var read1Common, read2Common int
	for true {
		num1, read1, e := getNum(inputFile1)
		read1Common += read1

		if e != nil && e != io.EOF {
			return nil, e
		}
		if num1 == nil {
			inputFile2Info, _ := inputFile2.Stat()
			buffer := make([]byte, inputFile2Info.Size()-int64(read2Common))
			inputFile2.Read(buffer)
			outputFile.Write(buffer)
			break
		}

		num2, read2, e := getNum(inputFile2)
		read2Common += read2

		if e != nil && e != io.EOF {
			return nil, e
		}
		if num2 == nil {
			read1Common -= read1
			inputFile1.Seek(-int64(read1), 1)
			inputFile1Info, _ := inputFile1.Stat()
			buffer := make([]byte, inputFile1Info.Size()-int64(read1Common))
			inputFile1.Read(buffer)
			outputFile.Write(buffer)
			break
		}

		var num int
		if *num1 < *num2 {
			num = *num1
			inputFile2.Seek(-int64(read2), 1)
			read2Common -= read2
		} else {
			num = *num2
			inputFile1.Seek(-int64(read1), 1)
			read1Common -= read1
		}
		outputFile.Write(intsToBytesSlice(num))
	}
	outputFile.Seek(0, 0)
	return outputFile, nil
}

func intsToBytesSlice(nums ...int) (result []byte) {
	for _, n := range nums {
		result = append(result, []byte(strconv.Itoa(n))...)
		result = append(result, byte(0x20))
	}
	return result
}

func QuickSort(seq *[]int) {
	quickSortRecursion(seq, 0, len(*seq)-1)
}

func quickSortRecursion(seq *[]int, firstIndex, lastIndex int) {
	if firstIndex < lastIndex {
		p := partition(seq, firstIndex, lastIndex)
		quickSortRecursion(seq, firstIndex, p)
		quickSortRecursion(seq, p+1, lastIndex)
	}
}

func partition(seq *[]int, firstIndex, lastIndex int) int {
	pivot := (*seq)[(firstIndex+lastIndex)/2]
	i := firstIndex
	j := lastIndex
	for i < j {
		for (*seq)[i] < pivot {
			i++
		}
		for (*seq)[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		(*seq)[i], (*seq)[j] = (*seq)[j], (*seq)[i]
		i++
		j--
	}
	return j
}

func mergeSequences(seq1, seq2 []int) (seq []int) {
	it1, it2 := 0, 0
	for it1 < len(seq1) && it2 < len(seq2) {
		if seq1[it1] < seq2[it2] {
			seq = append(seq, seq1[it1])
			it1++
		} else {
			seq = append(seq, seq2[it2])
			it2++
		}
	}
	for ; it1 < len(seq1); it1++ {
		seq = append(seq, seq1[it1])
	}
	for ; it2 < len(seq2); it2++ {
		seq = append(seq, seq2[it2])
	}

	return
}

func readSequence(file *os.File) (seq []int, n int, err error) {
	for true {
		num, nc, err := getNum(file)
		n += nc
		if err != nil && err != io.EOF {
			return nil, 0, err
		}
		if len(seq) == 0 && num != nil {
			seq = append(seq, *num)
			continue
		}
		if nc > 0 && num != nil {
			if seq[len(seq)-1] > *num {
				file.Seek(-int64(nc), 1)
				n -= nc
				break
			}
			seq = append(seq, *num)
		}
		if err == io.EOF {
			break
		}
	}

	return seq, n, nil
}

func getNum(file *os.File) (*int, int, error) {
	var (
		n    int
		num  *int
		err  error
		char byte
	)
	for true {
		char, err = getChar(file)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, 0, err
		}
		n++
		if char < 48 || char > 57 {
			break
		}
		if num == nil {
			num = new(int)
		}
		*num = *num*10 + (int(char) - 48)
	}
	return num, n, err
}

func getChar(file *os.File) (byte, error) {
	charBuff := make([]byte, 1)

	n, err := file.Read(charBuff)
	if n == 0 || err != nil {
		if err == io.EOF {
			return charBuff[0], err
		}
		return 0x00, err
	}

	return charBuff[0], nil
}
