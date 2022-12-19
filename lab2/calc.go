package calc_math_exp

import (
	"errors"
	"labs/structures"
	"strconv"
	"strings"
)

var opPriorities = map[string]int{
	"(": 0,
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"^": 3,
	"~": 4,
}

func ToPolishNotation(exp string) ([]string, error) {
	var result []string
	var operators structures.Stack[string]

	var numStr strings.Builder

	for i, ch := range strings.Split(exp, "") {
		// создается строка числа
		_, err := strconv.Atoi(ch)
		if err == nil || (ch == "." && numStr.Len() > 0) {
			numStr.WriteString(ch)
			continue
		}
		if numStr.Len() > 0 {
			result = append(result, numStr.String())
			numStr.Reset()
		}
		// записываем открывающую скобку в стек
		if ch == "(" {
			operators.Push(ch)
			continue
		}
		// если встретилась закрывающая скобка,
		// все операторы выталкиваюся из стека до открывающей скобки
		if ch == ")" {
			var elem string
			elem, err = operators.Pop()
			for elem != "(" && err == nil {
				result = append(result, elem)
				elem, err = operators.Pop()
			}
			continue
		}
		// если минус унарный, то он заменяется на тильду
		if ch == "-" && (i == 0 || ch == "(") {
			operators.Push("~")
			continue
		}
		// операторы из стека заносятся в результат, пока у полученного из стека оператора
		// более высокий приоритет
		for true {
			previosOp, err := operators.GetValueOfLastElement()
			if err != nil {
				break
			}
			if opPriorities[previosOp] >= opPriorities[ch] {
				previosOp, err = operators.Pop()
				if previosOp == "(" {
					return nil, errors.New("wrong brackets")
				}
				result = append(result, previosOp)
			} else {
				break
			}
		}
		// текущий (относительно цикла) оператор заносится в стек
		if _, ok := opPriorities[ch]; ok {
			operators.Push(ch)
		}
	}
	if numStr.Len() > 0 {
		result = append(result, numStr.String())
	}
	for !operators.IsEmpty() {
		elem, _ := operators.Pop()
		result = append(result, elem)
	}
	return result, nil
}
