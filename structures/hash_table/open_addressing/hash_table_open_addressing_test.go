package open_addressing

import (
	"github.com/stretchr/testify/assert"
	test "labs/testing"
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
	test.WalkSlice[string](cases, func(elem string) {
		found := hashTable.Search(elem)
		assert.False(t, found)
	})
	test.WalkSlice[string](cases, func(elem string) {
		err := hashTable.Add(elem)
		//fmt.Printf("%v\n", hashTable)
		//for i := 0; i < hashTable.size; i++ {
		//	fmt.Printf("%v\n", hashTable.table[i])
		//
		//}
		assert.NoError(t, err, elem)
	})
	test.WalkSlice[string](cases, func(elem string) {
		found := hashTable.Search(elem)
		assert.True(t, found)
	})
	test.WalkSlice[string](cases, func(elem string) {
		hashTable.Delete(elem)
	})
	test.WalkSlice[string](cases, func(elem string) {
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
	test.WalkSlice(cases, func(elem string) {
		err := tbl.Add(elem)
		assert.NoError(t, err, elem)
	})
	assert.Equal(t, 8, tbl.size, tbl.reservedCellsNum)
}
