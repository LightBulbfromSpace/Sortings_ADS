package lab1_v2

import "strings"

var bracketsType = []string{"()", "[]", "{}"}

func BracketsCheck(exp string) bool {
	for containsAny(exp, bracketsType) {
		for _, br := range bracketsType {
			exp = strings.Replace(exp, br, "", -1)
		}
	}
	return exp == ""
}

func containsAny(exp string, samples []string) bool {
	for _, sample := range samples {
		if strings.Contains(exp, sample) {
			return true
		}
	}
	return false
}
