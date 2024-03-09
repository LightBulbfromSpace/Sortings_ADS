package external_polyphase_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

type file struct {
	name string
	data string
}
type testFilesystem struct {
	files  []*file
	result *file
}

func TestExternalPolyPhaseSort(t *testing.T) {

	cases := []struct {
		initialFile file
		maxMemory   int
		expected    string
	}{
		{
			file{"a.dat", "4 1 85 2 43"},
			3,
			"1 2 4 43 85",
		},
		{
			file{"b.txt", "7 64 53454 545 97 124 827 3342 2341 34 1"},
			12,
			"1 7 34 64 97 124 545 827 2341 3342 53454",
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("test #%d", i+1), func(t *testing.T) {
			tmp, err, rm := createTmpFile(&tc.initialFile)

			defer rm()

			got, err := ExternalPolyPhaseSort(tmp, tc.maxMemory)
			if err != nil {
				t.Errorf("unexpected error: %q", err)
			}

			buffer := make([]byte, len(tc.expected))
			_, err = got.Seek(0, 0)
			if err != nil {
				t.Errorf("unexpected error: %q", err)
			}
			_, err = got.Read(buffer)
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, string(buffer))

			err = os.Remove(got.Name())
			if err != nil {
				t.Errorf("unexpected error: %q", err)
			}
		})
	}
}

func TestCreateAndDeleteNFiles(t *testing.T) {
	tmpFiles, err := createNFiles(5)
	assert.NoError(t, err)

	err = deleteFiles(tmpFiles)
	assert.NoError(t, err)
}

func TestCreateTmpFileName(t *testing.T) {
	var names []string
	nNames := 10
	for i := 0; i < nNames; i++ {
		names = append(names, createTmpFileName())
	}
	for i := 0; i < nNames-1; i++ {
		for j := i + 1; j < nNames; j++ {
			assert.NotEqual(t, names[i], names[j])
		}
	}
}

func TestReadMaxAmountOfNums(t *testing.T) {
	t.Run("too large number", func(t *testing.T) {
		_case := struct {
			initialFile file
			maxMemory   int
			expected    *[]int
		}{
			file{"b.dat", "764534124 827 3342 2341 34 1"},
			8,
			&[]int{},
		}

		tmp, err, rm := createTmpFile(&_case.initialFile)
		assert.NoError(t, err)
		defer rm()

		_, err = readMaxAmountOfNums(tmp, _case.maxMemory)
		assert.Error(t, err)

	})
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
			9,
			&[]int{7, 64},
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
	t.Run("the second reading from file", func(t *testing.T) {
		tmp, err, rm := createTmpFile(&cases[0].initialFile)
		assert.NoError(t, err)
		defer rm()

		_, err = readMaxAmountOfNums(tmp, cases[0].maxMemory)
		if err != nil {
			t.Errorf("unexpected error: %q", err)
		}
		nums, _ := readMaxAmountOfNums(tmp, 18)

		assert.Equal(t, []int{85, 2, 43}, *nums)
	})
}

func TestMergeSortedSequencesInFiles(t *testing.T) {
	cases := []*testFilesystem{
		{
			[]*file{
				{"a.dat", "1 2 4 43 85 87"},
				{"b.dat", "1 7 34 2347 3342"},
				{"c.dat", ""},
			},
			&file{"c.dat", "1 1 2 4 7 34 43 85 87 2347 3342"},
		},
		{
			[]*file{
				{"a.dat", "1 2 4 43 85 87 234219 699999 "},
				{"b.dat", "1 7 34 2347 3342 "},
				{"c.dat", ""},
			},
			&file{"c.dat", "1 1 2 4 7 34 43 85 87 2347 3342 234219 699999"},
		},
		{
			[]*file{
				{"a.dat", "1 2 4 43 87 699999 "},
				{"b.dat", "1 7 34 2347 3342 943843"},
				{"c.dat", ""},
			},
			&file{"c.dat", "1 1 2 4 7 34 43 87 2347 3342 699999 943843"},
		},
	}
	for i, fs := range cases {
		t.Run(fmt.Sprintf("test for #%d", i+1), func(t *testing.T) {
			tmps, _, delFuncs, err := createAllTmpFiles(fs)
			assert.NoError(t, err)

			got, err := mergeSortedSequencesInFiles(tmps[0], tmps[1])
			assert.NoError(t, err)

			buffer := make([]byte, len(fs.result.data))

			_, err = got.Seek(0, 0)
			if err != nil {
				t.Errorf("unexpected error: %q", err)
			}

			_, err = got.Read(buffer)
			if err != nil {
				t.Errorf("unexpected error: %q", err)
			}

			assert.Equal(t, fs.result.data, string(buffer))

			for _, rm := range delFuncs {
				rm()
			}

			err = os.Remove(got.Name())
			if err != nil {
				t.Errorf("unexpected error: %q", err)
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
		_, _, err = readSequence(tmp)
		if err != nil {
			t.Errorf("unexpected error: %q", err)
		}

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

			assert.Equal(t, tc.result, *got)
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
		err = os.Remove(fl.name)
		if err != nil {
			panic(err)
		}
	}

	_, err = tmp.Write([]byte(fl.data))
	if err != nil {
		return nil, err, nil
	}
	_, err = tmp.Seek(0, 0)
	if err != nil {
		return nil, err, nil
	}

	return tmp, nil, delete
}
