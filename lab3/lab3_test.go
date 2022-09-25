package solution_search

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFSimpleMultipliers(t *testing.T) {
	cases := []struct {
		x        int
		expected string
	}{
		{3, "1\n3\n"},
		{6, "1\n3\n5\n"},
		{7, "1\n3\n5\n7\n"},
		{21, "1\n3\n5\n7\n9\n15\n21\n"},
		{25, "1\n3\n5\n7\n9\n15\n21\n25\n"},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("test for %s", tc.expected), func(t *testing.T) {
			buff := &bytes.Buffer{}
			FSimpleMultipliers(buff, tc.x)
			assert.Equal(t, tc.expected, buff.String())
		})
	}
}
