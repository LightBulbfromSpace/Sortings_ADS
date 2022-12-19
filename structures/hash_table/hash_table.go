package hash_table

// HashTable - интерфейс реализации хеш-таблицы с любым типом данных
type HashTable[T any] interface {
	Search(value T) bool
	Add(value T) error
	Delete(value T)
}
