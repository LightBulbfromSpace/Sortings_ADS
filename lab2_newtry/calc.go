package calc

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var bracketsType = []string{"()", "[]", "{}"}
var operators = []string{"/", "*", "-", "+"}

func Calc(exp string) (float64, error) {

	return result, nil
}

func matchRegExpCheck(exp, regExp string) (newExp string, num float64, op uint8, err error) {
	brRegExp, err := regexp.Compile(regExp)
	if err != nil {
		return exp, 0, 0, err
	}
	if brRegExp.Match([]byte(exp)) {
		exp, num, err = extractNum(exp)
		if err != nil {
			return exp, 0, 0, err
		}
		exp, op, err = extractOperator(exp)
		if err != nil {
			return exp, num, 0, err
		}
	}
	return exp, num, op, nil
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
