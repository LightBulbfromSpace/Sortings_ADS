package structures

import (
	"errors"
)

type Stack[T any] struct {
	data       []T
	numOfItems int
}

func (s *Stack[T]) IsEmpty() bool {
	return s.numOfItems == 0
}

func (s *Stack[T]) Push(elem T) {
	s.data = append(s.data, elem)
	s.numOfItems++
}

func (s *Stack[T]) Pop() (T, error) {
	if s.numOfItems == 0 {
		return getZero[T](), errors.New("stack is empty")
	}
	s.numOfItems--
	elem := s.data[s.numOfItems]
	s.data = s.data[:s.numOfItems]
	return elem, nil
}

func (s *Stack[T]) GetValueOfLastElement() (T, error) {
	if s.numOfItems != 0 {
		return s.data[s.numOfItems-1], nil
	}
	return getZero[T](), errors.New("stack is empty")
}

type OperatorsStack struct {
	BracketsStack
}

type BracketsStack struct {
	data []int32
}

func (s *BracketsStack) SetData(exp string) {
	s.data = []int32(exp)
}

func (s *BracketsStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *BracketsStack) Push(symb int32) {
	s.data = append(s.data, symb)
}

func (s *BracketsStack) Pop() (int32, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	lastIndex := len(s.data) - 1
	lastSymb := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return lastSymb, nil
}

func (s *OperatorsStack) GetValueOfLastElement() (int32, error) {
	if !s.IsEmpty() {
		return s.data[len(s.data)-1], nil
	}
	return 0, errors.New("stack is empty")
}

func getZero[T any]() T {
	var zero T
	return zero
}
