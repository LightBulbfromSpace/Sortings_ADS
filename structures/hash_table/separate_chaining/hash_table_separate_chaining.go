package separate_chaining

import (
	"errors"
	"fmt"
	"math"
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
	HTSC.table[hash].data = data
	HTSC.table[hash].key = key
	// if there's no nodes in chain for this hash,
	// we need to put node in array cell
	if HTSC.table[hash] == nil {
		HTSC.table[hash] = newNode
		return nil
	}
	// ...else we're searching for last node in chain
	node = HTSC.table[hash]
	for node.nextCell != nil {
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

func (HTSC *HashTable[T]) Delete(key string) bool {
	if key == "" {
		return false
	}
	node, _ := HTSC.getNodeByKey(key)
	if node == nil {
		return false
	}
	if node.nextCell == nil {
		node.key = ""
		node.data = getZero[T]()
		node.nextCell = nil
		return true
	}
	node.key = node.nextCell.key
	node.data = node.nextCell.data
	node.nextCell = node.nextCell.nextCell
	return true
}

// getNodeByKey searches for node and gives full information about it
func (HTSC *HashTable[T]) getNodeByKey(key string) (node *hashTableCell[T], hash int) {
	hash = getHash(key)
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

// getHash counts hash value for string key,
func getHash(key string) int {
	return 0
}

// MD5 realization
func MD5(key string) string {
	// 1. Length aligment: new length of bits' array should be L = 512*n + 488,
	//    n - any natural number. The first bit being appended is 1, next bits are 0.
	keyBytes := []byte(key)
	keyBytes = append(keyBytes, 0x80)
	// 	  512 / 8 = 64, 448 / 8 = 56
	zeroBytesToAppend := 64 - (len(key) % 64) + 55
	zeros := make([]byte, zeroBytesToAppend)
	keyBytes = append(keyBytes, zeros...)

	// 2. Count 64-bit (8 bytes) appearance of key's initial length.
	//    If length is larger than 2^64 - 1, only less significant
	//    bits should be written.
	//    At first 4 less significant bytes should be appended,
	//    next - 4 most sinificant bytes.
	//	  * integer returned by len() can hold only 32-bits,
	//    so it's unnessesary to count MS-bytes.
	length := len(key) % int(math.Pow(2.0, 64.0))
	LSB := intToNBytes(length, 4)
	keyBytes = append(keyBytes, LSB...)
	keyBytes = append(keyBytes, zeros[:4]...)

	// 3. Create initialising vector
	a0 := 0x67452301 // A
	b0 := 0xefcdab89 // B
	c0 := 0x98badcfe // C
	d0 := 0x10325476 // D

	// 4. Counts in cycle
	// 	  4.1 Create 4 functions
	func1 := func(x, y, z int) int {
		return (x & y) | (^x & z)
	}
	func2 := func(x, y, z int) int {
		return (x & z) | (^z & y)
	}
	func3 := func(x, y, z int) int {
		a := (x & ^y) | (^x & y)
		return (a & ^z) | (^a & z)
	}
	func4 := func(x, y, z int) int {
		a := ^z | x
		return (a & ^y) | (^y & z)
	}
	// 	  4.2 Create function to count "constants", n in 1..64
	t := func(n int) int {
		return int(math.Pow(2, 32) * math.Abs(math.Sin(float64(n))))
	}
	// 	  4.3 Calculation in cycle
	//    shift numbers
	shift := []int{
		7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
		5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
		4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
		6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
	}
	// break keyBytes in chucks of 64 bytes
	chunkNum := len(keyBytes) / 64
	fmt.Println(len(keyBytes))
	for j := 0; j < chunkNum; j++ {
		a := a0
		b := b0
		c := c0
		d := d0
		for i := 0; i < 63; i++ {
			var f, g int
			if i < 16 {
				f = func1(b, c, d)
				g = i
			} else if i < 32 {
				f = func2(d, c, b)
				g = (5*i + 1) % 16
			} else if i < 48 {
				f = func3(b, c, d)
				g = (3*i + 5) % 16
			} else {
				f = func4(d, c, b)
				g = (7 * i) % 16
			}
			f = f + a + t(i) + fourBytesToInt(keyBytes[g*4:g*4+4])
			a = d
			d = c
			c = b
			b = b + cycleShiftToLeft(f, shift[i])
		}
		a0 += a
		b0 += b
		c0 += c
		d0 += d
	}
	result := intToNBytes(a0, 4)
	result = append(result, intToNBytes(b0, 4)...)
	result = append(result, intToNBytes(c0, 4)...)
	result = append(result, intToNBytes(d0, 4)...)
	return string(result)
}

func intToNBytes(num, n int) (result []byte) {
	for i := 0; i < n; i++ {
		_byte := uint8(num & 0xFF)
		result = append(result, _byte)
		num = num >> 8
	}
	fmt.Printf("%d\n", result)
	return
}

func fourBytesToInt(bytes []byte) int {
	return int(bytes[0])<<24 + int(bytes[1])<<16 + int(bytes[2])<<8 + int(bytes[3])
}

// cycleShiftToLeft - the most significant shifted bits appear at the right
func cycleShiftToLeft(num, n int) int {
	mask := int(math.Pow(2, float64(n)) - 1)
	mask = mask << (32 - n)
	msb := num & mask >> (32 - n)
	return (num << n) | msb
}

func getZero[T any]() T {
	var elem T
	return elem
}
