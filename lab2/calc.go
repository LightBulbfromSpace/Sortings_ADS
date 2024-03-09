package calc_math_exp

import (
	"errors"
	"github.com/LightBulbfromSpace/Sortings_ADS/structures"
	"math"
	"strconv"
	"strings"
)

var opPriorities = map[string]int{
	"(": 0,
	"[": 0,
	"{": 0,
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"^": 3,
	"~": 4,
}

var brackets = map[string]int{
	"(": 0,
	"[": 1,
	"{": 2,
	")": 3,
	"]": 4,
	"}": 5,
}

func toPolishNotation(exp string) ([]string, error) {
	var result []string
	var operators structures.Stack[string]
	var numStr strings.Builder

	for i, ch := range strings.Split(exp, "") {
		// пропуск пробела
		if ch == " " {
			continue
		}
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
		if ind, ok := brackets[ch]; ok && ind < 3 {
			operators.Push(ch)
			continue
		}
		// если встретилась закрывающая скобка,
		// все операторы выталкиваюся из стека до открывающей скобки
		if ind, ok := brackets[ch]; ok && ind > 2 {

			var elem string
			elem, err = operators.Pop()

			_, ok := brackets[elem]
			for !ok && err == nil {
				result = append(result, elem)
				elem, err = operators.Pop()
				_, ok = brackets[elem]
			}
			if brackets[ch]-brackets[elem] != 3 {
				return nil, errors.New("invalid brackets")
			}
			continue
		}
		// если минус унарный, то он заменяется на тильду
		index, ok := brackets[ch]
		if ch == "-" && (i == 0 || ok && index < 3) {
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
				if ind, ok := brackets[previosOp]; ok && ind < 3 {
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

func CalcPostfixNotation(exp string) (result float64, err error) {
	var nums structures.Stack[float64]
	postfixNotation, err := toPolishNotation(exp)
	if err != nil {
		return 0.0, err
	}
	for _, item := range postfixNotation {
		// если елемент является числом, заносим его в стек
		num, err := strconv.ParseFloat(item, 64)
		if err == nil {
			nums.Push(num)
			continue
		}
		if _, ok := opPriorities[item]; ok {
			if item == "~" {
				num, err = nums.Pop()
				if err != nil {
					return 0, err
				}
				num = -num
			} else {
				secondNum, e := nums.Pop()
				if e != nil {
					return 0, e
				}
				firstNum, e := nums.Pop()
				if e != nil {
					return 0, e
				}
				num, e = executeOperation(firstNum, secondNum, item)
				if e != nil {
					return 0, e
				}
			}
			nums.Push(num)
		}
	}
	return nums.Pop()
}

func executeOperation(n1, n2 float64, op string) (float64, error) {
	switch op {
	case "+":
		return n1 + n2, nil
	case "-":
		return n1 - n2, nil
	case "*":
		return n1 * n2, nil
	case "/":
		if n2 == 0 {
			return 0, errors.New("division by zero")
		}
		return n1 / n2, nil
	case "^":
		return math.Pow(n1, n2), nil
	default:
		return 0, errors.New("operator not found")
	}
}
