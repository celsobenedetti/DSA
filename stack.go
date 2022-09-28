package main

import "errors"

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() (data int, err error) {
	if len(s.items) == 0 {
		return -1, errors.New("Stack is empty")
	}
	length := len(s.items) - 1
	data = s.items[length]
	s.items = s.items[:length]

	return data, nil
}
