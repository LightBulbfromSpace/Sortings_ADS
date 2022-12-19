package hash_table

import (
	"fmt"
	"math"
)

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