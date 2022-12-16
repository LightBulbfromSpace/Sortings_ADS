package hash_table

import (
	"errors"
)

// HashTable - интерфейс реализации хеш-таблицы с любым типом данных
type HashTable[T any] interface {
	Search(value T) bool
	Add(value T) error
	Delete(value T)
}

// HashTableStrings - хеш-таблица строкового типа данных,
// реализующая интерфейс хеш-таблицы.
type HashTableStrings struct {
	table            []*hastTableCell[string]
	reservedCellsNum int
	size             int
	implicitResizing bool    // true - table can be resized automatically, false - only manually.
	limitFillRatio   float32 // ratio of filling to call autoresizing, when implicitResizing is set to true
}

// hastTableCell - ячейка хеш-таблицы.
type hastTableCell[T any] struct {
	value    T
	reserved bool
}

// NewHashTableStrings - инициализация строковой хеш-таблицы.
func NewHashTableStrings(size int, implicitResizing bool, limitFillRatio float32) *HashTableStrings {
	HTS := new(HashTableStrings)
	HTS.table = make([]*hastTableCell[string], size)
	HTS.size = size
	HTS.implicitResizing = implicitResizing
	HTS.limitFillRatio = limitFillRatio
	return HTS
}

// Search возвращает true, если если элемент присутствует в хеш-таблице
func (HTS *HashTableStrings) Search(exp string) bool {
	_, found := HTS.getElementHash(exp)
	return found
}

// Delete удаляет элемент путем изменения значения в поле ячейки reserved
// ("мягкое удаление"). Это необходимо для корректной работы поиска.
func (HTS *HashTableStrings) Delete(exp string) {
	hash, found := HTS.getElementHash(exp)
	if !found {
		return
	}
	HTS.table[hash].unsetValue()
	HTS.reservedCellsNum--
}

// Add добавляет элементы в таблицу.
func (HTS *HashTableStrings) Add(exp string) error {
	if HTS.reservedCellsNum == HTS.size && !HTS.implicitResizing {
		return errors.New("hash table is full")
	}
	if HTS.implicitResizing {
		fillRatio := float32(HTS.reservedCellsNum) / float32(HTS.size)
		if HTS.limitFillRatio <= fillRatio {
			HTS.Resize(2)
		}
	}
	err := HTS.setValue(exp)
	if err != nil {
		return err
	}
	HTS.reservedCellsNum++
	return nil
}

// Resize изменяет размер таблицы. Может вызываться автоматически, если
// установлен флаг implicitResizing.
func (HTS *HashTableStrings) Resize(resizeRatio float32) error {
	if resizeRatio <= 1 {
		return errors.New("resizing ratio is less or equal to 1")
	}

	newLength := int(float32(HTS.size) * resizeRatio)
	newTable := NewHashTableStrings(newLength, HTS.implicitResizing, HTS.limitFillRatio)
	for i := 0; i < HTS.size; i++ {
		if HTS.table[i] == nil || !HTS.table[i].reserved {
			continue
		}
		err := newTable.setValue(HTS.table[i].value)
		if err != nil {
			return err
		}
	}
	HTS.table = newTable.table
	HTS.size = newTable.size
	return nil
}

// Rehash полностью удаляет элементы, помеченных в качестве удаленных флагом reserved=false.
// Позволяет увеличить скорость работы таблицы.
func (HTS *HashTableStrings) Rehash() {
	newTable := NewHashTableStrings(HTS.size, HTS.implicitResizing, HTS.limitFillRatio)
	for i := 0; i < HTS.size; i++ {
		if HTS.table[i] == nil || !HTS.table[i].reserved {
			continue
		}
		newTable.setValue(HTS.table[i].value)
	}
	HTS.table = newTable.table
}

// setValue устанавливает значение строки в таблицу, если есть доступные ячейки.
func (HTS *HashTableStrings) setValue(exp string) error {
	hash, err := HTS.getNewHashValue(exp)
	if err != nil {
		return err
	}
	if HTS.table[hash] == nil {
		HTS.table[hash] = new(hastTableCell[string])
	}
	HTS.table[hash].setCellValue(exp)
	return nil
}

// setCellValue устанавливает значение ячейки.
func (cell *hastTableCell[T]) setCellValue(val T) {
	cell.value = val
	cell.reserved = true
}

// unsetValue делает "мягкое удаление".
func (cell *hastTableCell[T]) unsetValue() {
	cell.reserved = false
}

// getNewHashValue высчитывает хеш элемента. Используется двойное хеширование.
// Поиск хеша происходит до тех пор, пока не будет найдена свободная ячейка.
// Если по каким-либо причинам хеш не будет найдет, функция вернет ошибку.
func (HTS *HashTableStrings) getNewHashValue(exp string) (int, error) {
	indexHash := HTS.hashFuncInitial(exp)
	stepHash := HTS.hashFuncStep(exp)
	for it := 0; it < HTS.size; it++ {
		if HTS.table[indexHash] == nil || !HTS.table[indexHash].reserved {
			return indexHash, nil
		}
		indexHash = (indexHash + stepHash) % HTS.size
	}
	if HTS.implicitResizing {
		HTS.Resize(1.5)
		return HTS.getNewHashValue(exp)
	}
	return 0, errors.New("can't create new hash")
}

// getElementHash считает хеши и проверяет наличие элемента в таблице.
// Если соответствия хеша и элемента не найдено, возращает false вторым аргументом.
func (HTS *HashTableStrings) getElementHash(exp string) (int, bool) {
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

// getHashedValue использует алгоритм Горнера для вычисления хеша.
// Ключ должен быть взаимно простым числом с размером таблицы для лучшей
// работы хеш-функции.
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
