package structures

import "errors"

type Queue[T any] struct {
	data []T
	head int // points to firts element in Queue to Out
	tail int // points to first free place after last element in Quele
}

func (q *Queue[T]) Enqueue(elem T) {
	if len(q.data) == 0 {
		q.data = append(q.data, elem, getZero[T]())
		q.tail++
		return
	}
	if q.tail == len(q.data)-1 {
		if q.head == 0 {
			q.data[q.tail] = elem
			q.data = append(q.data, getZero[T]())
			q.tail++
		} else {
			q.data[q.tail] = elem
			q.tail = 1
		}
	} else {
		q.data[q.tail] = elem
		q.tail++
	}
}

func (q *Queue[T]) Dequele() (T, error) {
	if q.head == q.tail {
		return getZero[T](), errors.New("underflow")
	}
	elem := q.data[q.head]
	if q.head == len(q.data)-1 {
		q.head = 0
	} else {
		q.head++
	}
	return elem, nil
}
