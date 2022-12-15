package structures

import "errors"

type HashTable[T any] interface {
	Search(value T) (string, bool)
	Add(value T) error
	Delete(value T)
}

type HashTableStrings struct {
	table            []*hastTableCell[string]
	reservedCellsNum int
	size             int
}

type hastTableCell[T any] struct {
	value    T
	reserved bool
}

func NewHashTableStrings(size int) *HashTableStrings {
	HTS := new(HashTableStrings)
	HTS.table = make([]*hastTableCell[string], size)
	HTS.size = size
	return HTS
}

func (HTS *HashTableStrings) Search(exp string) bool {
	_, found := HTS.getElementID(exp)
	return found
}

func (HTS *HashTableStrings) Delete(exp string) {
	id, found := HTS.getElementID(exp)
	if !found {
		return
	}
	HTS.table[id].unsetValue()
	HTS.reservedCellsNum--
}

func (HTS *HashTableStrings) Add(exp string) error {
	if HTS.reservedCellsNum == HTS.size {
		return errors.New("hashtable is full")
	}
	indexHash := HTS.hashFuncInitial(exp)
	stepHash := HTS.hashFuncStep(exp)
	for it := 0; it < HTS.size; it++ {
		if HTS.table[indexHash] == nil || !HTS.table[indexHash].reserved {
			HTS.table[indexHash] = new(hastTableCell[string])
			HTS.table[indexHash].setValue(exp)
			HTS.reservedCellsNum++
			return nil
		}
		indexHash = (indexHash + stepHash) % HTS.size
	}
	return errors.New("place for item wasn't found")
}

func (cell *hastTableCell[T]) setValue(val T) {
	cell.value = val
	cell.reserved = true
}

func (cell *hastTableCell[T]) unsetValue() {
	cell.reserved = false
}

func (HTS *HashTableStrings) getElementID(exp string) (int, bool) {
	indexHash := HTS.hashFuncInitial(exp)
	stepHash := HTS.hashFuncStep(exp)
	it := 0
	for it < HTS.reservedCellsNum && HTS.table[indexHash] != nil {
		if HTS.table[indexHash].value == exp && HTS.table[indexHash].reserved {
			return indexHash, true
		}
		indexHash = (indexHash + stepHash) % HTS.size
		it++
	}
	return 0, false
}

func (HTS *HashTableStrings) walk() {

}

func (HTS *HashTableStrings) getHashedValue(exp string, key int) (result int) {
	for _, ch := range exp {
		result = (key*result + int(ch)) % HTS.size
	}
	result = (result*2 + 1) % HTS.size
	return result
}

// HashFuncInitial делает начальный расчет хеша для данного элемента
func (HTS *HashTableStrings) hashFuncInitial(exp string) (result int) {
	return HTS.getHashedValue(exp, HTS.size-1)
}

// HashFuncStep рассчитывает шаг, с которым будет совершаться проход по хеш-таблице для данного элемента
func (HTS *HashTableStrings) hashFuncStep(exp string) (result int) {
	return HTS.getHashedValue(exp, HTS.size+1)
}
