package separate_chaining

import (
	"fmt"
	"testing"
)

type _case struct {
	key  string
	data int
}

var tbl = NewTableSepChaining[int](10)
var cases = []_case{
	{"Sandra Lee", 291041},
	{"Tim Berton", 304823},
	{"Eugen Chernov", 495834},
}

//func TestHashTableAdd(t *testing.T) {
//
//	test.WalkSlice(cases, func(elem _case) {
//		err := tbl.Add(elem.key, elem.data)
//		assert.NoError(t, err)
//	})
//}
//
//func TestHashTableSearch(t *testing.T) {
//	fmt.Printf("here")
//	test.WalkSlice(cases, func(elem _case) {
//		_, found := tbl.Search(elem.key)
//		assert.True(t, found)
//	})
//}

func TestMD5(t *testing.T) {
	fmt.Println(MD5("Leewefibjbocgkjgclgjcjgcvljgcjlcjgvjhvlhvlihvivlivivivihilv"))
}
