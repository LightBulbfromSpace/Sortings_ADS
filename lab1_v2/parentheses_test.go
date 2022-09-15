package lab1_v2

import (
	"fmt"
	labtest "labs"
	"testing"
)

func TestParenthesesCheck(t *testing.T) {
	cases := []struct {
		exp   string
		valid bool
	}{
		{"", true},
		{"(", false},
		{")", false},
		{"()", true},
		{"(}", false},
		{"((({{", false},
		{"()]}", false},
		{"(){}[]", true},
		{"()[({}[])]", true},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("test for expression %s", tc.exp), func(t *testing.T) {
			got := BracketsCheck(tc.exp)
			labtest.AssertEqual(t, tc.valid, got)
		})
	}
}
