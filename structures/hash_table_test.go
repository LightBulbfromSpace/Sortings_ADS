package structures_test

import (
	"github.com/stretchr/testify/assert"
	s "labs/structures"
	"testing"
)

var hashTable = s.NewHashTableStrings(20)

func TestHashTable(t *testing.T) {
	cases := []string{
		"qfsdf",
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
		assert.NoError(t, err)
	})
	walkSlice[string](cases, func(elem string) {
		found := hashTable.Search(elem)
		assert.True(t, found)
	})
	walkSlice[string](cases, func(elem string) {
		hashTable.Delete(elem)
		found := hashTable.Search(elem)
		assert.False(t, found)
	})
}

func walkSlice[T any](s []T, f func(elem T)) {
	length := len(s)
	for i := 0; i < length; i++ {
		f(s[i])
	}
}
