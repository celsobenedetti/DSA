package linkedList

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	Head   *node
	length int
}

func (l *LinkedList) Prepend(data int) {
	newHead := &node{data: data}
	newHead.next = l.Head
	l.Head = newHead
	l.length++
}

func (l *LinkedList) Append(data int) {
	newNode := &node{data: data}
	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}
	l.length++
}

func (l *LinkedList) RemoveByValue(data int) (err error) {
	if l.length == 0 {
		return errors.New("LinkedList is Empty")
	}

	if l.Head.data == data {
		l.Head = l.Head.next
		l.length--
		return
	}

	previousToDelete := l.Head
	for previousToDelete.next != nil && previousToDelete.next.data != data {
		previousToDelete = previousToDelete.next
	}

	if previousToDelete.next == nil {
		return errors.New("Value not found in LinkedList")
	}

	previousToDelete.next = previousToDelete.next.next

	l.length--
	return
}

func (l LinkedList) PrintList() {
	toPrint := l.Head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
	}
	fmt.Println("")
}
