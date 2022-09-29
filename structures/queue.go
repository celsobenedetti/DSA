package structures

import (
	"errors"
)

type Queue struct {
	items []int
}

func (q *Queue) Push(data int) {
	q.items = append(q.items, data)
}

func (q *Queue) Dequeue() (data int, err error) {
	length := len(q.items)
	if length == 0 {
		return -1, errors.New("Queue is empty")
	}
	data = q.items[0]
	q.items = q.items[1:]

	return data, nil
}
