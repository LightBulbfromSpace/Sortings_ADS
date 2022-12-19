package calc_math_exp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestToPolishNotation(t *testing.T) {
	cases := []struct {
		exp, result string
	}{
		{"2+3=", "2 3 +"},
		{"33-(22+44)=", "33 22 44 + -"},
		{"(323/(96*4+8*7))=", "323 96 4 * 8 7 * + /"},
		{"3+4*2/(1-5)^2", "3 4 2 * 1 5 - 2 ^ / +"},
		{"3/9-5=", "3 9 / 5 -"},
		{"-2+7*(3/9)-5=", "2 ~ 7 3 9 / * + 5 -"},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("test for %s", tc.exp), func(t *testing.T) {
			got, err := ToPolishNotation(tc.exp)
			assert.NoError(t, err)
			assert.Equal(t, tc.result, strings.Join(got, " "))
		})
	}
}
