package separate_chaining

import (
	"errors"
)

// HashTable with separate chaining
type HashTable[T any] struct {
	table []*hashTableCell[T]
	size  int
}

// hashTableCell is a cell of hash table with separate chaining
type hashTableCell[T any] struct {
	key      string
	data     T
	nextCell *hashTableCell[T]
}

// NewTableSepChaining creates new table with separate chaining
func NewTableSepChaining[T any](size int) *HashTable[T] {
	HTSC := new(HashTable[T])
	HTSC.table = make([]*hashTableCell[T], size)
	HTSC.size = size
	return HTSC
}

// Search data by key, incapsulates getNodeByKey()
func (HTSC *HashTable[T]) Search(key string) (T, bool) {
	if key == "" {
		return getZero[T](), false
	}
	node, _ := HTSC.getNodeByKey(key)
	if node == nil {
		return getZero[T](), false
	}
	return node.data, true
}

// Add new value with new key
func (HTSC *HashTable[T]) Add(key string, data T) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	node, hash := HTSC.getNodeByKey(key)
	if node != nil {
		return errors.New("element already exists")
	}
	// create new node for chain
	newNode := new(hashTableCell[T])
	newNode.data = data
	newNode.key = key
	// if there's no nodes in chain for this hash,
	// we need to put node in array cell
	if HTSC.table[hash] == nil {
		HTSC.table[hash] = newNode
		return nil
	} else if HTSC.table[hash].key == "" {
		HTSC.table[hash] = newNode
		return nil
	}
	// ...else we're searching for last node in chain
	// (nil or initialized empty node)
	node = HTSC.table[hash]
	for node.nextCell != nil {
		if node.nextCell.key == "" {
			break
		}
		node = node.nextCell
	}
	// ...and append new node
	node.nextCell = newNode
	return nil
}

// Update value of existing key
func (HTSC *HashTable[T]) Update(key string, data T) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	node, _ := HTSC.getNodeByKey(key)
	if node == nil {
		return errors.New("element not found")
	}
	node.data = data
	return nil
}

// Delete element from hash table. Delete returns true if element was actually deleted,
// false if table doesn't contain such key.
func (HTSC *HashTable[T]) Delete(key string) bool {
	if key == "" {
		return false
	}
	node, _ := HTSC.getNodeByKey(key)
	if node == nil {
		return false
	}
	// the last element in chain can't be deleted fully (without effeciency loss),
	// so it becomes "empty initialized node".
	if node.nextCell == nil {
		node.key = ""
		node.data = getZero[T]()
		node.nextCell = nil
		return true
	}
	// deleting of the not last element in chain
	node.key = node.nextCell.key
	node.data = node.nextCell.data
	node.nextCell = node.nextCell.nextCell
	return true
}

// getNodeByKey searches for node and gives full information about it
func (HTSC *HashTable[T]) getNodeByKey(key string) (node *hashTableCell[T], hash int) {
	hash = getHash(key, HTSC.size)
	node = HTSC.table[hash]
	// iterates through chain while searched element not found
	// or the last element isn't reached
	for node != nil {
		if node.key == key {
			return
		}
		node = node.nextCell
	}
	return
}

// getHash counts hash value for string key
func getHash(key string, numberOfElem int) (result int) {
	num := numberOfElem - 1
	for _, ch := range key {
		result = (num*result + int(ch)) % numberOfElem
	}
	result = (result*2 + 1) % numberOfElem
	return result
}

func getZero[T any]() T {
	var elem T
	return elem
}
