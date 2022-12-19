package calc_math_exp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	test "labs/testing"
	"strings"
	"testing"
)

var cases = []struct {
	exp, postfixNotation string
	result               float64
}{
	{"2+3=", "2 3 +", 5},
	{"33-(22+44)=", "33 22 44 + -", -33},
	{"(323 / (96 * 4 + 8 * 7))=", "323 96 4 * 8 7 * + /", 0.7340909091},
	{"3+4*2/(1-5)^2", "3 4 2 * 1 5 - 2 ^ / +", 3.5},
	{"3/9-5=", "3 9 / 5 -", -4.6666666667},
	{"-2+7*(3/9)-5=", "2 ~ 7 3 9 / * + 5 -", -4.6666666667},
	{"5.5 + 6.6", "5.5 6.6 +", 12.1},
}

func TestToPolishNotation(t *testing.T) {
	for _, tc := range cases {
		t.Run(fmt.Sprintf("test for %s", tc.exp), func(t *testing.T) {
			got, err := toPolishNotation(tc.exp)
			assert.NoError(t, err)
			assert.Equal(t, tc.postfixNotation, strings.Join(got, " "))
		})
	}
}

func TestCalcPostfixNotation(t *testing.T) {
	for _, tc := range cases {
		t.Run(fmt.Sprintf("test for %s", tc.exp), func(t *testing.T) {
			got, err := CalcPostfixNotation(tc.exp)
			assert.NoError(t, err)
			test.AssertEqualFloat64(t, got, tc.result)
		})
	}
}
