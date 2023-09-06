package list

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	next, previous *Node[T]
	value          T
}

type LinkedList[T any] struct {
	head, tail *Node[T]
	size       int
	comparator func(T, T) int
}

/*
Creates a new List
Comparator function should return
* 0 if ir is equal
* 0>(Greater than 0) if it v1 is greater than v2
* 0<(Less than 0) if it v1 is lower than v2
*/
func NewList[T any](comparator func(T, T) int) *LinkedList[T] {
	return &LinkedList[T]{nil, nil, 0, comparator}
}

func (l LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

/*
Get size of the list
*/
func (l LinkedList[T]) Size() int {
	return l.size
}

func (l LinkedList[T]) GetFirst() T {
	return l.head.value
}

func (l LinkedList[T]) GetLast() T {
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

func (l LinkedList[T]) GetAt(i int) (T, error) {
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

func (l *LinkedList[T]) Add(value T, index int) error {
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
		//Should never get here because outOfBounds check before
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

func (l *LinkedList[T]) RemoveLast() T {
	if l.size == 1 {
		l.head = nil
	}
	value := l.tail.value
	l.tail = l.tail.previous
	l.tail.next = nil
	return value
}

func (l *LinkedList[T]) RemoveFirst() T {
	if l.size == 1 {
		l.tail = nil
	}
	value := l.head.value
	l.head = l.head.next
	l.head.previous = nil
	return value
}

func (l *LinkedList[T]) Remove(index int) (T, error) {
	switch {
	case index >= l.size:
		return *new(T), errors.New("OutOfBounds")
	case index == 0:
		return l.RemoveFirst(), nil
	case index == l.size-1:
		return l.RemoveLast(), nil
	}

	nodeToRemove, err := getNode[T](l.head, index)
	if err != nil {
		return *new(T), err
	}
	nodeToRemove.previous.next = nodeToRemove.next
	nodeToRemove.next.previous = nodeToRemove.previous
	l.size--
	return nodeToRemove.value, nil
}

func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LinkedList[T]) ToArray() []T {
	node := l.head
	arr := make([]T, 0, l.size)
	for node != nil {
		arr = append(arr, node.value)
		node = node.next
	}
	return arr
}

func (l *LinkedList[T]) PreatyPrint() {
	node := l.head
	fmt.Print(node.value)
	for node.next != nil {
		fmt.Print(" <-> ")
		fmt.Print(node.next.value)
		node = node.next
	}
	fmt.Print("\n")
}

// Stack Functions

/*
Removes the value at the top of the stack and retrieves the value
Equivalent to RemoveLast()
*/
func (l *LinkedList[T]) Pop() T {
	return l.RemoveLast()
}

func (l *LinkedList[T]) Push(value T) {
	l.AddLast(value)
}

func (l LinkedList[T]) Top() T {
	return l.tail.value
}

//
// Queue Functions
//
/*
Add value to the end of the queue
Equivalent to AddLast()
*/
func (l *LinkedList[T]) Queue(value T) {
	l.AddLast(value)
}

/*
Removes the value at the begining of the queue and retrives its value
Equivalent to RemoveFirst()
*/
func (l *LinkedList[T]) Dequeue() T {
	return l.RemoveFirst()
}

/*
Retrieves the value of the next int the queue without deleting it
*/
func (l LinkedList[T]) Peek() T {
	return l.head.value
}

func getNodeByValue[T any](node *Node[T], value T, cmp func(T, T) int, index int) (int, *Node[T]) {
	if node == nil || cmp(value, node.value) == 0 {
		return index, node
	}
	return getNodeByValue[T](node.next, value, cmp, index+1)
}

/*
Returns index of value or -1 if it isnt in the List
*/
func (l LinkedList[T]) IndexOf(value T) int {
	index, node := getNodeByValue[T](l.head, value, l.comparator, 0)
	if node == nil {
		return -1 //TODO: Make const
	}
	return index
}

/*
Returns true if value is in List, false if it isnt
*/
func (l LinkedList[T]) Contains(value T) bool {
	return l.IndexOf(value) >= 0
}

func (l *LinkedList[T]) ForEach(f func(T) T) {
	node := l.head
	for node != nil {
		node.value = f(node.value)
		node = node.next
	}
}