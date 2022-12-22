package external_polyphase_sort

import (
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//func ExternalPolyphaseSort(firstFile, secondFile, thirdFile io.ReadWriteSeeker) (io.ReadWriteSeeker, error) {
//	var err error
//	// true - end of file belongs to first file,
//	// false - end of file belongs to second file
//	var EOFofFirstFile bool
//	for err != io.EOF {
//		var seq1 []byte
//		seq1, err = readSequence(firstFile)
//		if err != nil {
//			if err == io.EOF {
//				EOFofFirstFile = true
//				continue
//			}
//			return nil, err
//		}
//
//		var seq2 []byte
//		seq2, err = readSequence(secondFile)
//		if err != nil {
//			if err == io.EOF {
//				EOFofFirstFile = false
//				continue
//			}
//			return nil, err
//		}
//		seq := mergeSequences(seq1, seq2)
//		_, err = thirdFile.Write(seq)
//		if err != nil {
//			return nil, err
//		}
//	}
//	if EOFofFirstFile {
//		for err != io.EOF {
//
//		}
//	}
//}

//func Round(firstFile, secondFile, thirdFile io.ReadWriteSeeker) error {
//	var nthFileisEmply int
//	seq1, err := readSequence(firstFile)
//	if err != nil {
//		if err == io.EOF {
//			nthFileisEmply = 1
//		} else {
//			return err
//		}
//	}
//
//	seq2, err := readSequence(secondFile)
//	if err != nil {
//		if err == io.EOF {
//			nthFileisEmply = 2
//		} else {
//			return err
//		}
//	}
//
//	seq := mergeSequences(seq1, seq2)
//	_, err = thirdFile.Write(seq)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func ExternalPolyphaseSort(initialFile *os.File, maxSeqInRAM int) (*os.File, error) {
	info, err := initialFile.Stat()
	if err != nil {
		return nil, err
	}
	initialFileSize := info.Size()
	numberOfTmpFiles := int(math.Ceil(float64(initialFileSize) / float64(maxSeqInRAM)))
	numberOfTmpFiles += numberOfTmpFiles % 2
	tmpFileSize := int(math.Ceil(float64(initialFileSize) / float64(numberOfTmpFiles)))

	//fmt.Printf("numberOfTmpFiles = %d, tmpFileSize = %d\n", numberOfTmpFiles, tmpFileSize)

	// divide information from initial file to several files
	var tmpFiles []*os.File
	for i := 0; i < numberOfTmpFiles; i++ {
		//buffer := make([]byte, tmpFileSize)
		//
		//n, err := initialFile.Read(buffer)
		//if err != nil && err != io.EOF {
		//	return nil, err
		//}
		//
		//buffer = buffer[:n]
		//fmt.Println(buffer)
		//nums := byteSlicetoIntSlice(buffer)
		//fmt.Println(*nums)
		nums, err := readMaxAmountOfNums(initialFile, tmpFileSize)
		if err != nil {
			return nil, err
		}
		QuickSort(nums)

		fmt.Printf("nums = %d\n", *nums)

		tmp, err := os.Create(createTmpFileName())
		if err != nil {
			return nil, err
		}

		buffer := intsToBytesSlice(*nums...)

		n, err := tmp.Write(buffer)
		fmt.Printf("n = %d\n", n)
		if err != nil {
			return nil, err
		}

		tmpFiles = append(tmpFiles, tmp)
	}

	// merge files into one decreasing number of files in two times with each interation
	for numberOfTmpFiles > 1 {
		var tmpFiles2 []*os.File
		for i := 0; i < numberOfTmpFiles-1; i++ {
			tmp, err := os.Create(createTmpFileName())
			if err != nil {
				return nil, err
			}
			tmp, err = mergeSortedSequencesInFiles(tmpFiles[i], tmpFiles[i+1], tmp)
			if err != nil {
				return nil, err
			}
			tmpFiles2 = append(tmpFiles2, tmp)
		}
		//for _, file := range tmpFiles {
		//	err := os.Remove(file.Name())
		//	if err != nil {
		//		return nil, err
		//	}
		//}
		tmpFiles = tmpFiles2
		numberOfTmpFiles >>= 1
	}
	initialFile.Seek(0, 0)
	//mergeSortedSequencesInFiles(tmpFiles[0], tmpFiles[1], initialFile)
	return tmpFiles[0], nil
}

func readMaxAmountOfNums(file *os.File, maxLength int) (*[]int, error) {
	var (
		err   error
		nums  []int
		n, nc int
	)
	for n < maxLength {
		var num int
		num, nc, err = getNum(file)
		if nc > maxLength {
			return nil, errors.New("not enough memory: numbers in file are too big")
		}
		if err != nil {
			if err == io.EOF {
				return &nums, err
			}
			return nil, err
		}
		if n > maxLength {
			file.Seek(-int64(nc), 1)
			break
		}
		nums = append(nums, num)
		n += nc
	}

	return &nums, nil
}

func createTmpFileName() string {
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("tmp_%d.dat", rand.Int63())
}

func byteSlicetoIntSlice(bytes []byte) *[]int {
	var nums []int
	var num int
	for _, ch := range bytes {
		if ch > 47 || ch < 58 {
			num = num*10 + (int(ch) - 48)
		} else {
			nums = append(nums, num)
			num = 0
		}
	}
	fmt.Printf("nums = %q", nums)
	return &nums
}

func mergeSortedSequencesInFiles(inputFile1, inputFile2, outputFile *os.File) (*os.File, error) {
	for true {
		n1, read1, err := getNum(inputFile1)
		if err != nil {
			if err == io.EOF {
				inputFile1.Seek(-int64(read1), 1)
				break
			}
			return nil, err
		}
		n2, read2, err := getNum(inputFile2)
		if err != nil {
			if err == io.EOF {
				inputFile2.Seek(-int64(read2), 1)
				break
			}
			return nil, err
		}
		var num int
		if n1 < n2 {
			num = n1
			inputFile2.Seek(-int64(read2), 1)
		} else {
			num = n2
			inputFile1.Seek(-int64(read1), 1)
		}

		_, err = outputFile.Write(intsToBytesSlice(num))
		if err != nil {
			return nil, err
		}
	}
	var err error
	for err != io.EOF {
		var n1 int
		n1, _, err = getNum(inputFile1)
		if err != nil && err != io.EOF {
			return nil, err
		}
		_, e := outputFile.Write(intsToBytesSlice(n1))
		if e != nil {
			return nil, err
		}
	}

	for err != io.EOF {
		var n2 int
		n2, _, err = getNum(inputFile2)
		if err != nil && err != io.EOF {
			return nil, err
		}
		_, e := outputFile.Write(intsToBytesSlice(n2))
		if e != nil {
			return nil, err
		}
	}
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
		if len(seq) == 0 {
			seq = append(seq, num)
			continue
		}
		if nc > 0 {
			if seq[len(seq)-1] > num {
				nInt64 := int64(nc)
				nInt64 = -nInt64
				file.Seek(nInt64, 1)
				n -= nc
				break
			}
			seq = append(seq, num)
		}
		if err == io.EOF {
			break
		}
	}

	return seq, n, nil
}

func getNum(file *os.File) (num int, n int, err error) {
	var char byte
	var numStr []byte
	for true {
		char, err = getChar(file)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, 0, err
		}
		n++
		if char < 48 || char > 57 {
			break
		}
		numStr = append(numStr, char)
	}
	if n > 0 {
		var e error
		num, e = strconv.Atoi(string(numStr))
		if e != nil {
			return num, n, e
		}
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
