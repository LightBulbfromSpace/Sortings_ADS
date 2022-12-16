package hash_table

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var hashTable = NewHashTableStrings(5, true, 0.75)

func TestHashTable(t *testing.T) {
	cases := []string{
		"qfsdf",
		"qfsdf",
		"kdfkskf",
		"kdfkskf",
		"nkl;hcx",
		"ml;ngh,;dydb",
		"kwjehfowehf^*H%^GU",
		"илывлды",
	}
	walkSlice[string](cases, func(elem string) {
		found := hashTable.Search(elem)
		assert.False(t, found)
	})
	walkSlice[string](cases, func(elem string) {
		err := hashTable.Add(elem)
		//fmt.Printf("%v\n", hashTable)
		//for i := 0; i < hashTable.size; i++ {
		//	fmt.Printf("%v\n", hashTable.table[i])
		//
		//}
		assert.NoError(t, err, elem)
	})
	walkSlice[string](cases, func(elem string) {
		found := hashTable.Search(elem)
		assert.True(t, found)
	})
	walkSlice[string](cases, func(elem string) {
		hashTable.Delete(elem)
	})
	walkSlice[string](cases, func(elem string) {
		found := hashTable.Search(elem)
		assert.False(t, found)
	})
}

func TestResize(t *testing.T) {
	oldSize := hashTable.size
	hashTable.Resize(2.5)
	assert.Equal(t, int(float32(oldSize)*2.5), hashTable.size)
}

func TestImplicitResizing(t *testing.T) {
	tbl := NewHashTableStrings(4, true, 1)
	cases := []string{
		"qfsdf",
		"kdfkskf",
		"nkl;hcx",
		"ml;ngh,;dydb",
		"kwjehfowehf^*H%^GU",
		"илывлды",
	}
	walkSlice(cases, func(elem string) {
		err := tbl.Add(elem)
		assert.NoError(t, err, elem)
	})
	assert.Equal(t, 8, tbl.size, tbl.reservedCellsNum)
}

func walkSlice[T any](s []T, f func(elem T)) {
	length := len(s)
	for i := 0; i < length; i++ {
		f(s[i])
	}
}
