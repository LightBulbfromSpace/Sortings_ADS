package structures

import "errors"

type BracketsStack struct {
	data []uint8
}

func (s *BracketsStack) SetData(exp string) {
	s.data = []uint8(exp)
}

func (s *BracketsStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *BracketsStack) Push(symb uint8) {
	s.data = append(s.data, symb)
}

func (s *BracketsStack) Pop() (uint8, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	lastIndex := len(s.data) - 1
	lastSymb := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return lastSymb, nil
}
