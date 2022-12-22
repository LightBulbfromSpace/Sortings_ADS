package external_polyphase_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

//func TestExternalPolyphaseSort(t *testing.T) {
//	cases := []struct{
//		testFileSystems *fstest.MapFS
//		result *fstest.MapFile
//	}{
//		{
//			&fstest.MapFS{
//				"a.txt" : &fstest.MapFile{Data: []byte("4 1 85 23 45 56 78 2 3 2 1 43")},
//				"b.txt" : &fstest.MapFile{Data: []byte("7 3342 2341 34 1 234 67 3423 4556 76553")},
//				"c.txt" : &fstest.MapFile{},
//			},
//			&fstest.MapFile{Data: []byte("1 1 1 2 2 3 4 7 23 34 45 56 67 85 ...")},
//		},
//
//		{
//			&fstest.MapFS{
//			"a.txt" : &fstest.MapFile{Data: []byte("4 1 85  2  43")},
//			"b.txt" : &fstest.MapFile{Data: []byte("7 3342 2341 34 1")},
//			"c.txt" : &fstest.MapFile{},
//			},
//			&fstest.MapFile{Data: []byte("1 1 2 4 7 34 43 85 2341 3342")},
//		},
//	}
//	for i, tc := range cases {
//		t.Run(fmt.Sprintf("test #%d", i), func(t *testing.T) {
//			f1, _ := os.Open("a.txt")
//			f2, _ := os.Open("b.txt")
//			f3, _ := os.Open("c.txt")
//			got, err := ExternalPolyphaseSort(f1, f2, f3)
//			assert.NoError(t, err)
//			assert.Equal(t, *tc.result, got)
//		})
//	}
//}

type file struct {
	name string
	data string
}
type testFilesystem struct {
	files  []*file
	result *file
}

func TestExternalPolyphaseSort(t *testing.T) {

	cases := []struct {
		initialFile file
		maxMemory   int
		expected    string
	}{
		//{
		//	file{"a.dat", "4 1 85 2 43"},
		//	3,
		//	"1 2 4 85 1 43",
		//},
		{
			file{"b.dat", "7 64 53454 545 97 124 827 3342 2341 34 1"},
			12,
			"1 7 34 64 97 124 545 827 2341 3342 53454",
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {
			tmp, err, rm := createTmpFile(&tc.initialFile)

			defer rm()

			got, err := ExternalPolyphaseSort(tmp, tc.maxMemory)
			if err != nil && err != io.EOF {
				t.Errorf("unexpected error: %q", err)
			}

			//fmt.Println(len(got))

			buffer := make([]byte, len(tc.expected))
			got.Seek(0, 0)
			n, _ := got.Read(buffer)
			fmt.Println(n)

			assert.Equal(t, tc.expected, string(buffer))
		})
	}
}

func TestReadMaxAmountOfNums(t *testing.T) {
	cases := []struct {
		initialFile file
		maxMemory   int
		expected    *[]int
	}{
		{
			file{"a.dat", "4 1 85 2 43"},
			4,
			&[]int{4, 1},
		},
		{
			file{"b.dat", "7 64 53454 545 97 124 827 3342 2341 34 1"},
			8,
			&[]int{7, 64, 53454},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test for #%d", i+1), func(t *testing.T) {
			tmp, err, rm := createTmpFile(&tc.initialFile)
			assert.NoError(t, err)
			defer rm()

			nums, err := readMaxAmountOfNums(tmp, tc.maxMemory)
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, nums)
		})
	}
}

func TestMergeSortedsequencesInFiles(t *testing.T) {
	cases := []*testFilesystem{
		{
			[]*file{
				&file{"a.dat", "1 2 4 43 85 87"},
				&file{"b.dat", "1 7 34 2347 3342"},
				&file{"c.dat", ""},
			},
			&file{"c.dat", "1 1 2 4 7 34 43 85 87 2347 3342"},
		},
	}
	for i, fs := range cases {
		t.Run(fmt.Sprintf("test for #%d", i+1), func(t *testing.T) {
			tmps, result, delFuncs, err := createAllTmpFiles(fs)
			assert.NoError(t, err)

			got, err := mergeSortedSequencesInFiles(tmps[0], tmps[1], result)
			assert.NoError(t, err)

			buffer := make([]byte, len(fs.result.data))
			got.Seek(0, 0)
			got.Read(buffer)
			//assert.NoError(t, err)

			assert.Equal(t, fs.result.data, string(buffer))

			for _, rm := range delFuncs {
				rm()
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	cases := []struct {
		initial, expected []int
	}{
		{[]int{2, 20, 7, 2341, 3342, 43, 3, 4}, []int{2, 3, 4, 7, 20, 43, 2341, 3342}},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test for #%d", i+1), func(t *testing.T) {
			QuickSort(&tc.initial)
			assert.Equal(t, tc.expected, tc.initial)
		})
	}
}

func TestMergeSequences(t *testing.T) {
	cases := []struct {
		seq1, seq2, expected []int
	}{
		{[]int{2, 20, 43}, []int{3, 4, 7, 2341, 3342}, []int{2, 3, 4, 7, 20, 43, 2341, 3342}},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test for #%d", i+1), func(t *testing.T) {
			got := mergeSequences(tc.seq1, tc.seq2)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestReadSequence(t *testing.T) {
	cases := []struct {
		fl       *file
		expected []int
	}{
		{&file{"a.dat", "1 4 85 2 43"}, []int{1, 4, 85}},
		{&file{"b.dat", "3 4 7 2341 3342"}, []int{3, 4, 7, 2341, 3342}},
		{&file{"c.dat", "4 1 85 2 43"}, []int{4}},
	}
	var deleteFuncs []func()
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test for #%d", i+1), func(t *testing.T) {
			tmp, err, del := createTmpFile(tc.fl)
			assert.NoError(t, err)
			deleteFuncs = append(deleteFuncs, del)

			seq, _, err := readSequence(tmp)
			//tmp.Seek(int64(n), 1)
			if err != nil && err != io.EOF {
				t.Errorf("unexpected error: %q", err)
			}
			assert.Equal(t, tc.expected, seq)
		})
	}

	t.Run("second read from file", func(t *testing.T) {
		tmp, err := os.Open("a.dat")
		assert.NoError(t, err)
		readSequence(tmp)

		seq, _, err := readSequence(tmp)
		if err != nil && err != io.EOF {
			t.Errorf("unexpected error: %q", err)
		}
		assert.Equal(t, []int{2, 43}, seq)
	})

	for _, rm := range deleteFuncs {
		rm()
	}
}

func TestGetNum(t *testing.T) {
	cases := []struct {
		input  *file
		result int
	}{
		{&file{"a.txt", "678"}, 678},
		{&file{"a.txt", "342 234"}, 342},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {

			tmp, err, rm := createTmpFile(tc.input)
			assert.NoError(t, err)
			defer rm()

			got, _, err := getNum(tmp)
			if err != nil && err != io.EOF {
				t.Errorf("unexpected error: %q", err)
			}

			assert.Equal(t, tc.result, got)
		})
	}
}

func createAllTmpFiles(fs *testFilesystem) (initialFiles []*os.File, resultFile *os.File, deleteFuncs []func(), err error) {
	for _, fl := range fs.files {
		tmp, err, rm := createTmpFile(fl)
		if err != nil {
			return nil, nil, nil, err
		}
		deleteFuncs = append(deleteFuncs, rm)
		initialFiles = append(initialFiles, tmp)
	}
	resultFile, err, rm := createTmpFile(fs.result)
	if err != nil {
		return nil, nil, nil, err
	}
	deleteFuncs = append(deleteFuncs, rm)
	return
}

func createTmpFile(fl *file) (tmp *os.File, err error, delete func()) {
	tmp, err = os.Create(fl.name)
	if err != nil {
		return nil, err, nil
	}
	delete = func() {
		os.Remove(fl.name)
	}

	_, err = tmp.Write([]byte(fl.data))
	tmp.Seek(0, 0)
	if err != nil {
		return nil, err, nil
	}

	return tmp, nil, delete
}
