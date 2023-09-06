package list

import (
	"errors"
)

type Node[T any] struct {
	next, previous *Node[T]
	value          T
}

type LinkedList[T any] struct {
	head, tail *Node[T]
	size       int
}

func NewList[T any]() *LinkedList[T] {
	return &LinkedList[T]{nil, nil, 0}
}

func (l LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) Add(value T) {
	var new_node = Node[T]{nil, nil, value}

	if l.head == nil {
		l.head = &new_node
	} else {
		l.tail.next = &new_node
		new_node.previous = l.tail
	}
	l.tail = &new_node
	l.size++
}

func (l LinkedList[T]) Peek() T {
	return l.head.value
}

func (l LinkedList[T]) Top() T {
	return l.tail.value
}

func getNode[T any](node *Node[T], i int) (*Node[T], error) {
	if node == nil {
		return nil, errors.New("OutOfBounds")
	}
	if i == 0 {
		return node, nil
	}
	return getNode[T](node.next, i-1)
}

func (l LinkedList[T]) GeatAt(i int) (T, error) {
	node, err := getNode[T](l.head, i)
	if err != nil {
		return *new(T), err
	}
	return node.value, nil
}
