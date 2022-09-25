package calc

import (
	"fmt"
	"labs"
	"testing"
)

func TestCalc(t *testing.T) {
	cases := []struct {
		exp    string
		result float64
	}{
		{"2+3=", 5},
		{"33-22=", 11},
		{"3/9=", 0.333333333},
		//{"3.3+2.2=", 5.5},
		//{"3/9-5=", -4.666666667},
		//{"2+7*(3/9)-5=", -0.666666667},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("test for %s", tc.exp), func(t *testing.T) {
			got, err := Calc(tc.exp)
			labtest.AssertNoError(t, err)
			labtest.AssertEqualFloat64(t, got, tc.result)
		})
	}
}
