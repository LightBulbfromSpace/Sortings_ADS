package calc_math_exp

import (
	"fmt"
	labtest "github.com/LightBulbfromSpace/Sortings_ADS/testing"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var cases = []struct {
	exp             string
	postfixNotation string
	result          float64
	valid           bool
}{
	{"2+3=", "2 3 +", 5, true},
	{"33-(22+44)=", "33 22 44 + -", -33, true},
	{"(323 / (96 * 4 + 8 * 7))=", "323 96 4 * 8 7 * + /", 0.7340909091, true},
	{"3+4*2/(1-5)^2", "3 4 2 * 1 5 - 2 ^ / +", 3.5, true},
	{"3/9-5=", "3 9 / 5 -", -4.6666666667, true},
	{"-2+7*(3/9)-5=", "2 ~ 7 3 9 / * + 5 -", -4.6666666667, true},
	{"-2+{7*(3/9)}-5=", "2 ~ 7 3 9 / * + 5 -", -4.6666666667, true},
	{"5.5 + 6.6", "5.5 6.6 +", 12.1, true},
	{"-2+7*[3/9)-5=", "2 ~ 7 3 9 / * + 5 -", 0, false},
	{"-2+{7*(3/9}}-5=", "2 ~ 7 3 9 / * + 5 -", 0, false},
	// valid for postfix notation convertion, invalid for calculation
	{"5.5 / 0", "5.5 0 /", 0, true},
}

func TestToPolishNotation(t *testing.T) {
	for _, tc := range cases {
		t.Run(fmt.Sprintf("test for %s", tc.exp), func(t *testing.T) {
			got, err := toPolishNotation(tc.exp)
			if tc.valid {
				assert.NoError(t, err)
				assert.Equal(t, tc.postfixNotation, strings.Join(got, " "))
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestCalcPostfixNotation(t *testing.T) {
	cases[len(cases)-1].valid = false
	for _, tc := range cases {
		t.Run(fmt.Sprintf("test for %s", tc.exp), func(t *testing.T) {
			got, err := CalcPostfixNotation(tc.exp)
			if tc.valid {
				assert.NoError(t, err)
				labtest.AssertEqualFloat64(t, got, tc.result)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
