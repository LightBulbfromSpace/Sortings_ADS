package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnqueueAndDequeue(t *testing.T) {
	t.Run("set elements in queue with ", func(t *testing.T) {
		var q Queue[int]
		var values = []int{5, 2, 8, 6}
		q.data = make([]int, 4)
		for _, value := range values {
			q.Enqueue(value)
		}
		assert.Equal(t, 0, q.head)
		assert.Equal(t, 4, q.tail)
		assert.Equal(t, values, q.data)
	})
	t.Run("set elements in new queue (without capacity)", func(t *testing.T) {
		var q Queue[int]
		var values = []int{5, 2, 8, 6}
		for _, value := range values {
			q.Enqueue(value)
		}
		assert.Equal(t, 0, q.head)
		assert.Equal(t, 4, q.tail)
		assert.Equal(t, values, q.data)
	})

	t.Run("set elements after dequeue (in the beginning of array)", func(t *testing.T) {
		var q Queue[int]
		var values = []int{5, 2, 8, 6}
		for _, value := range values {
			q.Enqueue(value)
		}

		q.Dequele()

		cases := []struct {
			elem          int
			expectedArray []int
			expectedHead  int
			expectedTail  int
		}{
			{3, []int{3, 5, 2, 8, 6}, 2, 1},
			{4, []int{3, 4, 5, 2, 8, 6}, 3, 2},
			{7, []int{3, 4, 7, 5, 2, 8, 6}, 4, 3},
		}
		for _, tc := range cases {
			q.Enqueue(tc.elem)

			assert.Equal(t, tc.expectedHead, q.head)
			assert.Equal(t, tc.expectedTail, q.tail)
			assert.Equal(t, tc.expectedArray, q.data)
		}

		q.Dequele()
		q.Dequele()

		q.Enqueue(10)
		assert.Equal(t, 6, q.head)
		assert.Equal(t, 4, q.tail)
		assert.Equal(t, []int{3, 4, 7, 10, 2, 8, 6}, q.data)

	})

	var q Queue[int]

	t.Run("just initialized queue", func(t *testing.T) {
		_, err := q.Dequele()
		assert.Error(t, err)
	})

	var values = []int{5, 2, 8, 6}
	for _, value := range values {
		q.Enqueue(value)
	}
	t.Run("queue not empty", func(t *testing.T) {
		for _, elem := range values {
			value, err := q.Dequele()
			assert.NoError(t, err)
			assert.Equal(t, elem, value)
			//fmt.Println(q.data, q.head, q.tail)
		}
	})

	t.Run("underflow", func(t *testing.T) {
		_, err := q.Dequele()
		assert.Error(t, err)
	})
	t.Run("enqueue after underflow", func(t *testing.T) {
		q.Enqueue(10)
		assert.Equal(t, 0, q.head)
		assert.Equal(t, 1, q.tail)
		assert.Equal(t, []int{10, 2, 8, 6}, q.data)
		//fmt.Println(q.data, q.head, q.tail)
	})

	t.Run("enqueue after dequeue", func(t *testing.T) {
		q.data = []int{5, 2, 8, 6}
		q.head = 3
		q.tail = 4
		q.Enqueue(12)
		assert.Equal(t, 3, q.head)
		assert.Equal(t, 1, q.tail)
		assert.Equal(t, []int{12, 2, 8, 6}, q.data)
		//fmt.Println(q.data, q.head, q.tail)

		value, err := q.Dequele()
		assert.NoError(t, err)
		assert.Equal(t, 6, value)
		//fmt.Println(q.data, q.head, q.tail)

		value, err = q.Dequele()
		assert.NoError(t, err)
		assert.Equal(t, 12, value)
		//fmt.Println(q.data, q.head, q.tail)

		q.Enqueue(13)
		assert.Equal(t, 1, q.head)
		assert.Equal(t, 2, q.tail)
		assert.Equal(t, []int{12, 13, 8, 6}, q.data)
		//fmt.Println(q.data, q.head, q.tail)
	})
}
