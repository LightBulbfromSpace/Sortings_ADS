package structures

import (
	"errors"
)

type Queue[T any] struct {
	data []T
	head int // points to firts element in Queue to Out
	tail int // points to first free place after last element in Quele
}

func (q *Queue[T]) Enqueue(elem T) {
	if len(q.data) == q.tail && q.head != 0 {
		q.tail = 0
	}
	if len(q.data) == q.tail || q.head-1 == q.tail {
		q.data = append(q.data, getZero[T]())
		if q.head != 0 {
			for i := len(q.data) - 1; i >= q.head; i-- {
				q.data[i] = q.data[i-1]
			}
			q.tail %= len(q.data) - 1
			q.head++
		}
	}
	q.data[q.tail] = elem
	q.tail++
}

func (q *Queue[T]) Dequele() (T, error) {
	if q.head == q.tail {
		q.head, q.tail = 0, 0
		return getZero[T](), errors.New("underflow")
	}

	if q.head == len(q.data) {
		q.head = 0
	}
	elem := q.data[q.head]
	q.head++

	return elem, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.tail == q.head
}
