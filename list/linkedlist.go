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

func (l *LinkedList[T]) AddLast(value T) {
	newNode := Node[T]{nil, nil, value}

	if l.IsEmpty() {
		l.head = &newNode
	} else {
		l.tail.next = &newNode
		newNode.previous = l.tail
	}
	l.tail = &newNode
	l.size++
}

func (l *LinkedList[T]) AddFirst(value T) {
	new_node := Node[T]{nil, nil, value}

	if l.IsEmpty() {
		l.AddLast(value)
	} else {
		new_node.next = l.head
		l.head = &new_node
		l.size++
	}
}

func (l *LinkedList[T]) AddAt(value T, index int) error {
	switch {
	case index >= l.size:
		return errors.New("OutOfBounds")
	case index == 0:
		l.AddFirst(value)
		return nil
	case index == l.size-1:
		l.AddLast(value)
		return nil
	}

	nodeAtPosition, err := getNode[T](l.head, index)
	if err != nil {
		//SHould never get here because outOfBounds check before
		return err
	}
	previousNode := nodeAtPosition.previous
	newNode := Node[T]{
		previous: previousNode,
		next:     nodeAtPosition,
		value:    value,
	}

	previousNode.next = &newNode
	nodeAtPosition.previous = &newNode

	return nil
}
