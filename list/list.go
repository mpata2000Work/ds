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
	Comparator func(T, T) int
}

type Stack[T any] interface {
	Pop() (T, error)
	Push(T)
	Top() (T, error)
	IsEmpty() bool
	Size() int
}

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Size() int
}

/*
NewList Creates a new List
Comparator function should return
* 0 if ir is equal
* 0>(Greater than 0) if it v1 is greater than v2
* 0<(Less than 0) if it v1 is lower than v2
*/
func NewList[T any](comparator func(T, T) int) LinkedList[T] {
	return LinkedList[T]{nil, nil, 0, comparator}
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

/*
Size Get size of the list
*/
func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) GetFirst() (T, error) {
	if l.IsEmpty() {
		return *new(T), errors.New("NoSuchElement")
	}
	return l.head.value, nil
}

func (l *LinkedList[T]) GetLast() (T, error) {
	if l.IsEmpty() {
		return *new(T), errors.New("NoSuchElement")
	}
	return l.tail.value, nil
}

func (l *LinkedList[T]) getNodeAt(index int) (*Node[T], error) {
	if l.IsEmpty() {
		return nil, errors.New("NoSuchElement")
	}
	if index >= l.size || index < 0 {
		return nil, errors.New("OutOfBounds")
	}
	var node *Node[T]
	var moveForward bool
	var i int
	if index > l.size/2 {
		node = l.tail
		moveForward = false
		i = l.size - 1
	} else {
		node = l.head
		moveForward = true
	}

	for i != index {
		if moveForward {
			node = node.next
			i++
		} else {
			node = node.previous
			i--
		}
	}
	return node, nil
}

func (l *LinkedList[T]) GetAt(index int) (T, error) {
	node, err := l.getNodeAt(index)
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
	if l.IsEmpty() {
		l.AddLast(value)
		return
	}
	new_node := Node[T]{nil, nil, value}
	new_node.next = l.head
	l.head.previous = &new_node
	l.head = &new_node
	l.size++
}

func (l *LinkedList[T]) Add(index int, value T) error {
	if index == 0 {
		l.AddFirst(value)
		return nil
	}
	if index == l.size {
		l.AddLast(value)
		return nil
	}

	nodeAtPosition, err := l.getNodeAt(index)
	if err != nil {
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
	l.size++

	return nil
}

func (l *LinkedList[T]) AddArray(arr []T) error {
	if len(arr) == 0 {
		return errors.New("NoElementsInArray")
	}

	for _, v := range arr {
		l.AddLast(v)
	}
	return nil
}

func (l *LinkedList[T]) AddArrayAt(index int, arr []T) error {
	if index > l.size || index < 0 {
		return errors.New("OutOfBounds")
	}
	if len(arr) == 0 {
		return errors.New("NoElementsInArray")
	}
	if index == l.size {
		return l.AddArray(arr)
	}
	tempList := ListFromArray[T](arr, nil)
	nodeAtPos, err := l.getNodeAt(index)
	if err != nil {
		return err
	}

	if nodeAtPos == l.head {
		l.head = tempList.head
	} else {
		prevNode := nodeAtPos.previous
		prevNode.next = tempList.head
		tempList.head.previous = prevNode
	}
	nodeAtPos.previous = tempList.tail
	tempList.tail.next = nodeAtPos
	l.size += tempList.size

	return nil
}

func (l *LinkedList[T]) Set(index int, value T) error {
	node, err := l.getNodeAt(index)
	if err != nil {
		return err
	}
	node.value = value
	return nil
}

func (l *LinkedList[T]) RemoveLast() (T, error) {
	if l.tail == nil {
		return *new(T), errors.New("NoSuchElement")
	}
	value := l.tail.value
	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.previous
		l.tail.next = nil
	}
	l.size--
	return value, nil
}

func (l *LinkedList[T]) RemoveFirst() (T, error) {
	if l.head == nil {
		return *new(T), errors.New("NoSuchElement")
	}
	if l.size == 1 {
		return l.RemoveLast()
	}
	value := l.head.value
	l.head = l.head.next
	l.head.previous = nil
	l.size--
	return value, nil
}

func (l *LinkedList[T]) Remove(index int) (T, error) {
	if index == l.size-1 {
		return l.RemoveLast()
	} else if index == 0 {
		return l.RemoveFirst()
	}

	nodeToRemove, err := l.getNodeAt(index)
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

func (l *LinkedList[T]) PrettyPrint() {
	node := l.head
	fmt.Print(node.value)
	for node.next != nil {
		fmt.Print(" <-> ")
		fmt.Print(node.next.value)
		node = node.next
	}
	fmt.Print("\n")
}

/*
IndexOf Returns index of value or -1 if it isnt in the List
*/
func (l *LinkedList[T]) IndexOf(value T) (int, error) {
	if l.Comparator == nil {
		return -1, errors.New("comparator is nil")
	}
	node := l.head
	var index int
	for node != nil {
		if l.Comparator(value, node.value) == 0 {
			return index, nil
		}
		node = node.next
		index++
	}
	return -1, nil
}

/*
Returns true if value is in List, false if it isnt
*/
func (l *LinkedList[T]) Contains(value T) (bool, error) {
	index, err := l.IndexOf(value)
	if err != nil {
		return false, err
	}
	return index != -1, nil
}

func (l *LinkedList[T]) ForEach(f func(T) T) {
	if f == nil {
		return
	}
	node := l.head
	for node != nil {
		node.value = f(node.value)
		node = node.next
	}
}

func (l *LinkedList[T]) Map(f func(T) T) LinkedList[T] {
	mapList := NewList[T](l.Comparator)
	if f == nil {
		return mapList
	}
	node := l.head
	for node != nil {
		mapList.AddLast(f(node.value))
		node = node.next
	}
	return mapList
}

func (l *LinkedList[T]) Filter(f func(T) bool) LinkedList[T] {
	filtList := NewList[T](l.Comparator)
	if f == nil {
		return filtList
	}
	node := l.head
	for node != nil {
		if f(node.value) {
			filtList.AddLast(node.value)
		}
		node = node.next
	}
	return filtList
}

func ListFromArray[T any](arr []T, comparator func(T, T) int) LinkedList[T] {
	l := NewList[T](comparator)
	if arr == nil {
		return l
	}

	l.AddArray(arr)
	return l
}

// Stack Functions

func NewStack[T any]() Stack[T] {
	return new(LinkedList[T])
}

/*
Removes the value at the top of the stack and retrieves the value
Equivalent to RemoveLast()
*/
func (l *LinkedList[T]) Pop() (T, error) {
	return l.RemoveLast()
}

func (l *LinkedList[T]) Push(value T) {
	l.AddLast(value)
}

func (l *LinkedList[T]) Top() (T, error) {
	return l.GetLast()
}

//
// Queue Functions
//

func NewQueue[T any]() Queue[T] {
	return new(LinkedList[T])
}

/*
Add value to the end of the queue
Equivalent to AddLast()
*/
func (l *LinkedList[T]) Enqueue(value T) {
	l.AddLast(value)
}

/*
Removes the value at the begining of the queue and retrives its value
Equivalent to RemoveFirst()
*/
func (l *LinkedList[T]) Dequeue() (T, error) {
	return l.RemoveFirst()
}

/*
Peek Retrieves the value of the next int the queue without deleting it
*/
func (l *LinkedList[T]) Peek() (T, error) {
	return l.GetFirst()
}
