package lab1_v1

import "labs/structures"

func ParenthesesCheck(exp string) bool {
	stack := &structures.BracketsStack{}

	for _, char := range exp {
		if isCloseBracket(char) {
			elem, err := stack.Pop()
			if err != nil {
				return false
			}
			if !isRightBrackets(elem, char) {
				return false
			}
		} else {
			stack.Push(char)
		}
	}
	return stack.IsEmpty()
}

func isCloseBracket(char int32) bool {
	return char == ')' || char == ']' || char == '}'
}

func isRightBrackets(openBracket, closeBracket int32) bool {
	return openBracket == '(' && closeBracket == ')' ||
		openBracket == '[' && closeBracket == ']' ||
		openBracket == '{' && closeBracket == '}'
}
