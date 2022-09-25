package calc

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	ADD_SUB_PRIORITY     = 0
	MULT_DIV_PRIORITY    = 1
	PARENTHESES_PRIORITY = 2
)

var bracketsType = []string{"()", "[]", "{}"}
var operators = []string{"/", "*", "-", "+"}

func Calc(exp string) (float64, error) {
	brRegExp, err := regexp.Compile(`^\d+.\d+=`)
	if err != nil {
		return 0, err
	}

	found := brRegExp.Match([]byte(exp))
	if !found {
		brRegExp, err := regexp.Compile(`\d+.\d+=`)
		if err != nil {
			return 0, err
		}

		found := brRegExp.Match([]byte(exp))
		if !found {
			return 0, errors.New("no match with regexp")
		}

	}

	result, err := simpleCalc(exp)
	if err != nil {
		return 0, err
	}
	/*if ok, brackets := containsAny(exp, bracketsType); ok {
		pattern := fmt.Sprintf("%c.+%c", brackets[0], brackets[1])
		brRegExp, err := regexp.Compile(pattern)
		if err != nil {
			return 0, err
		}
		found := brRegExp.Find([]byte(exp))
		Calc(string(found))
	}*/
	/*var result float64
	priority := priorityCheck(str)
	str, num, err := extractNum(str)
	if err != nil {
		return 0, err
	}*/

	return result, nil
}

func containsAny(exp string, samples []string) (bool, string) {
	for _, sample := range samples {
		if strings.Contains(exp, sample) {
			return true, sample
		}
	}
	return false, ""
}

func priorityCheck(exp string) (priority int) {
	for _, char := range exp {
		if (char == '*' || char == '/') && priority < MULT_DIV_PRIORITY {
			priority = MULT_DIV_PRIORITY
			continue
		}
		if char == '(' && priority < PARENTHESES_PRIORITY {
			priority = PARENTHESES_PRIORITY
			break
		}
	}
	return
}

func extractNum(exp string) (string, float64, error) {
	var (
		numStr strings.Builder
		num    float64
	)
	for exp[0] > 47 && exp[0] < 58 || exp[0] == 46 {
		numStr.WriteString(string(exp[0]))
		exp = exp[1:]
	}
	num, err := strconv.ParseFloat(numStr.String(), 64)
	if err != nil {
		return exp, 0, err
	}
	return exp, num, nil
}

func extractOperator(exp string) (string, uint8, error) {
	if exp[0] == '+' || exp[0] == '-' || exp[0] == '*' || exp[0] == '/' {
		return exp[1:], exp[0], nil
	}
	return exp, 0, errors.New("unable to parse operator")
}

func simpleCalc(exp string) (float64, error) {
	exp, num1, err := extractNum(exp)
	if err != nil {
		return 0, err
	}

	exp, op, err := extractOperator(exp)
	if err != nil {
		return 0, err
	}

	exp, num2, err := extractNum(exp)
	if err != nil {
		return 0, err
	}

	switch op {
	case '+':
		return num1 + num2, nil
	case '-':
		return num1 - num2, nil
	case '*':
		return num1 * num2, nil
	case '/':
		return num1 / num2, nil
	default:
		return 0, errors.New("operator not found")
	}
}
func regExpSearch() {

}
