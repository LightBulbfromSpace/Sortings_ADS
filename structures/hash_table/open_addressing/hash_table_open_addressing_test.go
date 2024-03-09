package open_addressing

import (
	labtest "github.com/LightBulbfromSpace/Sortings_ADS/testing"
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
	labtest.WalkSlice[string](cases, func(elem string) {
		found := hashTable.Search(elem)
		assert.False(t, found)
	})
	labtest.WalkSlice[string](cases, func(elem string) {
		err := hashTable.Add(elem)
		//fmt.Printf("%v\n", hashTable)
		//for i := 0; i < hashTable.size; i++ {
		//	fmt.Printf("%v\n", hashTable.table[i])
		//
		//}
		assert.NoError(t, err, elem)
	})
	labtest.WalkSlice[string](cases, func(elem string) {
		found := hashTable.Search(elem)
		assert.True(t, found)
	})
	labtest.WalkSlice[string](cases, func(elem string) {
		hashTable.Delete(elem)
	})
	labtest.WalkSlice[string](cases, func(elem string) {
		found := hashTable.Search(elem)
		assert.False(t, found)
	})
}

func TestResize(t *testing.T) {
	oldSize := hashTable.size
	err := hashTable.Resize(2.5)
	if err != nil {
		t.Errorf("unexpected error: %q", err)
	}

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
	labtest.WalkSlice(cases, func(elem string) {
		err := tbl.Add(elem)
		assert.NoError(t, err, elem)
	})
	assert.Equal(t, 8, tbl.size, tbl.reservedCellsNum)
}
