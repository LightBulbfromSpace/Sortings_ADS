package separate_chaining

import (
	"fmt"
	labtest "github.com/LightBulbfromSpace/Sortings_ADS/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

type _case struct {
	key  string
	data int
}

var tbl = NewTableSepChaining[int](5)
var cases = []_case{
	{"Sandra Lee", 291041},
	{"Tim Berton", 304823},
	{"Eugen Chernov", 495834},
	{"Tomas Bern", 304903},
	{"Artem Verbitsky", 445834},
	{"Andrey Listov", 495834},
	{"Bill Sapolsky", 304903},
	{"John Nowak", 445834},
}

func TestHashTableAdd(t *testing.T) {
	t.Run("call Add() on non-existing in table elements", func(t *testing.T) {
		labtest.WalkSlice(cases, func(elem _case) {
			err := tbl.Add(elem.key, elem.data)
			assert.NoError(t, err)
		})
	})
	//printAllChains[int](tbl)
	t.Run("call Add() on existing element", func(t *testing.T) {
		err := tbl.Add("Andrey Listov", 495834)
		assert.Error(t, err)
	})
	t.Run("add element after deletion of the last element in chain", func(t *testing.T) {
		found := tbl.Delete("Andrey Listov")
		assert.True(t, found)
		found = tbl.Delete("Tim Berton")
		assert.True(t, found)
		//printAllChains[int](tbl)
		err := tbl.Add("Andrey Listov", 495834)
		assert.NoError(t, err)
		err = tbl.Add("Tim Berton", 304823)
		assert.NoError(t, err)
		//printAllChains[int](tbl)
		walkAllNodesInTable[int](t, "", tbl, func(t *testing.T, node *hashTableCell[int], value interface{}) {
			assert.NotEqual(t, value, node.key)
		})
	})
}

func TestHashTableSearch(t *testing.T) {
	//printAllChains[int](tbl)
	labtest.WalkSlice(cases, func(elem _case) {
		data, found := tbl.Search(elem.key)
		assert.True(t, found, elem.key)
		rightData := false
		for _, tc := range cases {
			if tc.data == data && elem.key == tc.key {
				rightData = true
			}
		}
		assert.True(t, rightData, elem.key, elem.data)
	})
}

func TestHashTableUpdate(t *testing.T) {
	tbl := tbl
	labtest.WalkSlice(cases, func(elem _case) {
		err := tbl.Update(elem.key, 0)
		assert.NoError(t, err)
	})
	walkAllNodesInTable[int](t, 0, tbl, func(t *testing.T, node *hashTableCell[int], value interface{}) {
		assert.Equal(t, 0, node.data)
	})
	err := tbl.Update("Ray Sonne", 279827)
	assert.Error(t, err)
}

func TestHashTableDelete(t *testing.T) {
	labtest.WalkSlice(cases, func(elem _case) {
		found := tbl.Delete(elem.key)
		assert.True(t, found)
	})
}

func printAllChains[T any](tbl *HashTable[T]) {
	for i := 0; i < len(tbl.table); i++ {
		if tbl.table[i] != nil {
			fmt.Printf("%q", tbl.table[i].key)
			node := tbl.table[i].nextCell
			for node != nil {
				fmt.Printf("%q", node.key)
				node = node.nextCell
			}
			fmt.Println()
		} else {
			fmt.Println(nil)
		}
	}
	fmt.Println()
}

func walkAllNodesInTable[T any](t *testing.T, value interface{}, tbl *HashTable[T],
	f func(t *testing.T, node *hashTableCell[T], value interface{})) {
	for i := 0; i < tbl.size; i++ {
		if tbl.table[i] != nil {
			node := tbl.table[i]
			for node != nil {
				f(t, node, value)
				node = node.nextCell
			}
		}
	}
}
