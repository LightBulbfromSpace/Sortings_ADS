package lab1_v1

import "labs/structures"

func ParenthesesCheck(exp string) bool {
	stack := &structures.BracketsStack{}

	for _, char := range exp {
		if isCloseBracket(uint8(char)) {
			elem, err := stack.Pop()
			if err != nil {
				return false
			}
			if !isRightBrackets(elem, uint8(char)) {
				return false
			}
		} else {
			stack.Push(uint8(char))
		}
	}
	return stack.IsEmpty()
}

func isCloseBracket(char uint8) bool {
	return char == ')' || char == ']' || char == '}'
}

func isRightBrackets(openBracket, closeBracket uint8) bool {
	return openBracket == '(' && closeBracket == ')' ||
		openBracket == '[' && closeBracket == ']' ||
		openBracket == '{' && closeBracket == '}'
}
