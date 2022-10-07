package structures

import (
	"errors"
)

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Push(data T) {
	q.items = append(q.items, data)
}

func (q Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) Dequeue() (data T, err error) {
	length := len(q.items)
	if length == 0 {
		return data, errors.New("Queue is empty")
	}
	data = q.items[0]
	q.items = q.items[1:]

	return data, nil
}
