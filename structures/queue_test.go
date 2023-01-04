package structures

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var q Queue[int]

func TestEnqueue(t *testing.T) {
	var values = []int{5, 2, 8, 6}
	for _, value := range values {
		q.Enqueue(value)
	}

	fmt.Println(q.data, q.head, q.tail)
	assert.Equal(t, 5, q.data[0])
	q.Dequele()
	fmt.Println(q.data, q.head, q.tail)
	assert.Equal(t, 0, q.data[0])
}

func TestDequeue(t *testing.T) {

}
