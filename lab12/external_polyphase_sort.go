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

func cycle() {

}

func ExternalPolyphaseSort(initialFile *os.File, maxSeqInRAM int) (*os.File, error) {
	info, err := initialFile.Stat()
	if err != nil {
		return nil, err
	}
	initialFileSize := info.Size()
	numberOfTmpFiles := int(math.Ceil(float64(initialFileSize) / float64(maxSeqInRAM)))
	tmpFileSize := int(math.Ceil(float64(initialFileSize) / float64(numberOfTmpFiles)))

	// create temporary files
	tmpFiles, err := createNFiles(numberOfTmpFiles)
	if err != nil {
		return nil, err
	}
	// divide information from initial file to several files
	for _, f := range tmpFiles {
		nums, err := readMaxAmountOfNums(initialFile, tmpFileSize)
		if err != nil {
			return nil, err
		}
		QuickSort(nums)

		_, err = f.Write(intsToBytesSlice(*nums...))
		f.Seek(0, 0)
	}

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

	//var tmpFiles []*os.File
	//var deleteFuncs []func() error
	//for i := 0; i < numberOfTmpFiles; i++ {
	//	nums, err := readMaxAmountOfNums(initialFile, tmpFileSize)
	//	if err != nil {
	//		return nil, err
	//	}
	//	QuickSort(nums)
	//
	//	tmp, del, err := createTemporaryFile()
	//	if err != nil {
	//		return nil, err
	//	}
	//	deleteFuncs = append(deleteFuncs, del)
	//
	//	bytes := intsToBytesSlice(*nums...)
	//	_, err = tmp.Write(bytes)
	//	tmp.Seek(0, 0)
	//	//fmt.Printf("n = %d bytes = %q\n", n, bytes)
	//	if err != nil {
	//		return nil, err
	//	}
	//	tmp.Read(bytes)
	//	fmt.Printf("bytes = %q f = %p\n", bytes, tmp)
	//	tmp.Seek(0, 0)
	//	tmpFiles = append(tmpFiles, tmp)
	//	if i > 0 {
	//		tmpFiles[i-1].Read(bytes)
	//		fmt.Printf("bytes = %q f = %p\n", bytes, tmpFiles[i-1])
	//	}
	//}
	//for _, f := range tmpFiles {
	//	buffer := make([]byte, 15)
	//	n, err := f.Read(buffer)
	//	fmt.Printf("f = %p\n", f)
	//	fmt.Println(n, err, buffer)
	//}

	// merge files into one decreasing number of files in two times with each interation
	//for numberOfTmpFiles > 1 {
	//	var deleteFuncs2 []func() error
	//	var tmpFiles2 []*os.File
	//	for i := 0; i < numberOfTmpFiles-1; i++ {
	//		tmp, del, err := createTemporaryFile()
	//		if err != nil {
	//			return nil, err
	//		}
	//		deleteFuncs2 = append(deleteFuncs2, del)
	//
	//		tmp, err = mergeSortedSequencesInFiles(tmpFiles[i], tmpFiles[i+1], tmp)
	//		if err != nil {
	//			return nil, err
	//		}
	//		tmpFiles2 = append(tmpFiles2, tmp)
	//	}
	//	tmpFiles = tmpFiles2
	//	//for _, del := range deleteFuncs {
	//	//	err := del()
	//	//	if err != nil {
	//	//		return nil, err
	//	//	}
	//	//}
	//	deleteFuncs = deleteFuncs2
	//	fmt.Printf("numberOfTmpFiles = %d\n", numberOfTmpFiles)
	//	numberOfTmpFiles /= 2
	//	fmt.Printf("numberOfTmpFiles = %d\n", numberOfTmpFiles)
	//}

	//for numberOfTmpFiles > 1 {
	//	var deleteFuncs2 []func() error
	//	var tmpFiles2 []*os.File
	//	for i := 0; i < numberOfTmpFiles-1; i++ {
	//		tmp, del, err := createTemporaryFile()
	//		if err != nil {
	//			return nil, err
	//		}
	//		deleteFuncs2 = append(deleteFuncs2, del)
	//
	//		tmp, err = mergeSortedSequencesInFiles(tmpFiles[i], tmpFiles[i+1], tmp)
	//		if err != nil {
	//			return nil, err
	//		}
	//		tmpFiles2 = append(tmpFiles2, tmp)
	//	}
	//	fmt.Println(tmpFiles)
	//	tmpFiles = tmpFiles2
	//	fmt.Println(tmpFiles)
	//	for _, del := range deleteFuncs {
	//		err := del()
	//		if err != nil {
	//			return nil, err
	//		}
	//	}
	//	deleteFuncs = deleteFuncs2
	//	numberOfTmpFiles /= 2
	//}
	//initialFile.Seek(0, 0)
	//mergeSortedSequencesInFiles(tmpFiles[0], tmpFiles[1], initialFile)
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
			if err == io.EOF {
				return &nums, err
			}
			return nil, err
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

//func byteSlicetoIntSlice(bytes []byte) *[]int {
//	var nums []int
//	var num int
//	for _, ch := range bytes {
//		if ch > 47 || ch < 58 {
//			num = num*10 + (int(ch) - 48)
//		} else {
//			nums = append(nums, num)
//			num = 0
//		}
//	}
//	fmt.Printf("nums = %q", nums)
//	return &nums
//}

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
			//read1Common += read1
			inputFile2.Seek(-int64(read2), 1)
			read2Common -= read2
		} else {
			num = *num2
			//read2Common += read2
			inputFile1.Seek(-int64(read1), 1)
			read1Common -= read1
		}
		outputFile.Write(intsToBytesSlice(num))
	}
	outputFile.Seek(0, 0)
	return outputFile, nil
}

//func mergeSortedSequencesInFiles(inputFile1, inputFile2 *os.File) (*os.File, error) {
//	inputFile1.Seek(0, 0)
//	inputFile2.Seek(0, 0)
//	outputFile, err := createTemporaryFile()
//	if err != nil {
//		return nil, err
//	}
//
//	var (
//		err1, err2               error
//		read1Common, read2Common int
//	)
//	for err1 != io.EOF && err2 != io.EOF {
//
//		var (
//			read1, read2 int
//			n1, n2       *int
//		)
//
//		n1, read1, err1 = getNum(inputFile1)
//		if err1 != nil && err1 != io.EOF {
//			return nil, err1
//		}
//		n2, read2, err2 = getNum(inputFile2)
//		if err2 != nil && err2 != io.EOF {
//			return nil, err2
//		}
//
//		if n1 == nil {
//			inputFile2.Seek(-int64(read2), 1)
//			read2Common -= read2
//			break
//		}
//
//		if n2 == nil {
//			inputFile1.Seek(-int64(read1), 1)
//			read1Common -= read1
//			break
//		}
//
//		var num int
//		if *n1 < *n2 {
//			num = *n1
//			read1Common += read1
//			inputFile2.Seek(-int64(read2), 1)
//			read2Common -= read2
//		} else {
//			num = *n2
//			read2Common += read2
//			inputFile1.Seek(-int64(read1), 1)
//			read1Common -= read1
//		}
//
//		written, e := outputFile.Write(intsToBytesSlice(num))
//		if e != nil {
//			return nil, e
//		}
//		fmt.Printf("written = %d\n", written)
//	}
//
//	if err1 != io.EOF {
//		inputFile1Info, _ := inputFile1.Stat()
//		buffer := make([]byte, inputFile1Info.Size()-int64(read1Common))
//		_, e := inputFile1.Read(buffer)
//		if e != nil {
//			return nil, err
//		}
//
//		_, e = outputFile.Write(buffer)
//		if e != nil {
//			return nil, err
//		}
//	}
//
//	if err2 != io.EOF {
//		inputFile2Info, _ := inputFile2.Stat()
//		buffer := make([]byte, inputFile2Info.Size()-int64(read2Common))
//		_, e := inputFile2.Read(buffer)
//		if e != nil {
//			return nil, err
//		}
//
//		_, e = outputFile.Write(buffer)
//		if e != nil {
//			return nil, err
//		}
//	}
//	outputFile.Seek(0, 0)
//	return outputFile, nil
//}

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
