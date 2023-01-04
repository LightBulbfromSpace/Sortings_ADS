package structures

import (
	labtest "github.com/LightBulbfromSpace/Labs_AlgorithmsAndDataStructure/testing"
	"testing"
)

func TestStack(t *testing.T) {
	elems := []uint8{'(', ')'}
	stack := &BracketsStack{}
	for _, char := range elems {
		stack.Push(int32(char))
	}
	elem1, err := stack.Pop()
	labtest.AssertNoError(t, err)
	labtest.AssertEqual(t, int32(elems[1]), elem1)

	epm := stack.IsEmpty()
	labtest.AssertEqual(t, false, epm)

	elem2, err := stack.Pop()
	labtest.AssertNoError(t, err)
	labtest.AssertEqual(t, int32(elems[0]), elem2)

	epm = stack.IsEmpty()
	labtest.AssertEqual(t, true, epm)
}
