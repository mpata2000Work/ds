package list

import (
	"errors"
	"fmt"
)

var (
	ErrorNoSuchElement     = errors.New("NoSuchElement")
	ErrorOutOfBounds       = errors.New("OutOfBounds")
	ErrorNoElementsInArray = errors.New("NoElementsInArray")
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

// IsEmpty Returns true if the list is empty
func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Size Get size of the list
func (l *LinkedList[T]) Size() int {
	return l.size
}

// GetFirst Returns the first element of the list
func (l *LinkedList[T]) GetFirst() (T, error) {
	if l.IsEmpty() {
		return *new(T), ErrorNoSuchElement
	}
	return l.head.value, nil
}

// GetLast Returns the last element of the list
func (l *LinkedList[T]) GetLast() (T, error) {
	if l.IsEmpty() {
		return *new(T), ErrorNoSuchElement
	}
	return l.tail.value, nil
}

// getNodeAt Returns the node at the given index from the closest end of the list
func (l *LinkedList[T]) getNodeAt(index int) (*Node[T], error) {
	if l.IsEmpty() {
		return nil, ErrorNoSuchElement
	}
	if index >= l.size || index < 0 {
		return nil, errors
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

// Get Returns the value at the given index from the closest end of the list
func (l *LinkedList[T]) Get(index int) (T, error) {
	node, err := l.getNodeAt(index)
	if err != nil {
		return *new(T), err
	}
	return node.value, nil
}

// AddFirst Adds value to the begining of the list
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

// AddLast Adds value to the end of the list
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

// Add Adds value to the given index of the list
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

// AddAll Adds all values to the end of the list. Returns error if array is empty
func (l *LinkedList[T]) AddAll(arr ...T) error {
	if len(arr) == 0 {
		return ErrorNoElementsInArray
	}

	for _, v := range arr {
		l.AddLast(v)
	}
	return nil
}

// AddArrayAt Adds all values to the given index of the list. Returns error if index is out of bounds or if array is empty
func (l *LinkedList[T]) AddArrayAt(index int, arr []T) error {
	if index > l.size || index < 0 {
		return errors
	}
	if len(arr) == 0 {
		return ErrorNoElementsInArray
	}
	if index == l.size {
		return l.AddAll(arr...)
	}
	tempList := LinkedListFromArray[T](arr, nil)
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

// Set Sets the value at the given index of the list. Returns error if index is out of bounds
func (l *LinkedList[T]) Set(index int, value T) error {
	node, err := l.getNodeAt(index)
	if err != nil {
		return err
	}
	node.value = value
	return nil
}

// RemoveLast Removes the last element of the list and returns its value or error if list is empty
func (l *LinkedList[T]) RemoveLast() (T, error) {
	if l.IsEmpty() {
		return *new(T), ErrorNoSuchElement
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

// RemoveFirst Removes the first element of the list and returns its value or error if list is empty
func (l *LinkedList[T]) RemoveFirst() (T, error) {
	if l.IsEmpty() {
		return *new(T), ErrorNoSuchElement
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

// Remove Removes the element at the given index of the list and returns its value or error if index is out of bounds or list is empty
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

// getNodeByValue Returns the index and node of the first occurence of value in the list or -1 and nil if it isnt in the list
func (l *LinkedList[T]) getNodeByValue(value T) (int, *Node[T]) {
	node := l.head
	index := 0
	for node != nil {
		if l.Comparator(node.value, value) == 0 {
			return index, node
		}
		node = node.next
		index++
	}
	return -1, nil
}

// RemoveElement Removes the first occurence of value in the list and returns true or false if it isnt in the list
func (l *LinkedList[T]) RemoveValue(value T) bool {
	if l.IsEmpty() {
		return false
	}
	i, node := l.getNodeByValue(value)
	switch i {
	case -1:
		return false
	case 0:
		l.RemoveFirst()
		return true
	case l.size - 1:
		l.RemoveLast()
		return true
	}
	node.previous.next = node.next
	node.next.previous = node.previous
	l.size--
	return true
}

// Clear Removes all elements from the list
func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// ToArray Returns an array of all the values in the list
func (l *LinkedList[T]) ToArray() []T {
	node := l.head
	arr := make([]T, 0, l.size)
	for node != nil {
		arr = append(arr, node.value)
		node = node.next
	}
	return arr
}

// PrettyPrint Prints the list in a pretty format
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

// IndexOf Returns the index of the first occurence of value in the list or -1 if it isnt in the list.
// Returns error if comparator is nil
func (l *LinkedList[T]) IndexOf(value T) (int, error) {
	if l.Comparator == nil {
		return -1, errors.New("comparator is nil")
	}
	i, _ := l.getNodeByValue(value)
	return i, nil
}

// Contains Returns true if the list contains the given value.
// Returns error if comparator is nil
func (l *LinkedList[T]) Contains(value T) (bool, error) {
	index, err := l.IndexOf(value)
	if err != nil {
		return false, err
	}
	return index != -1, nil
}

// ForEach Applies the function f to each element of the lists
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

// Map Applies the function f to each element of the lists and returns a new list with the results
func (l *LinkedList[T]) Map(f func(T) T) LinkedList[T] {
	mapList := NewLinkedList[T](l.Comparator)
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

// Filter Applies the function f to each element of the lists and returns a new list with the elements that returned true
func (l *LinkedList[T]) Filter(f func(T) bool) LinkedList[T] {
	filtList := NewLinkedList[T](l.Comparator)
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

// LinkedListFromArray Creates a new LinkedList from the given array
func LinkedListFromArray[T any](arr []T, comparator func(T, T) int) LinkedList[T] {
	l := NewLinkedList[T](comparator)
	if arr == nil {
		return l
	}

	l.AddAll(arr...)
	return l
}

//-----------------------------//
//       Stack Functions       //
//-----------------------------//

// NewStack Creates a new Stack
func NewStack[T any]() Stack[T] {
	return new(LinkedList[T])
}

// Pop Removes the value at the top of the stack and retrieves the value. Returns error if stack is empty.
// Equivalent to RemoveLast()
func (l *LinkedList[T]) Pop() (T, error) {
	return l.RemoveLast()
}

// Push Adds value to the top of the stack.
// Equivalent to AddLast()
func (l *LinkedList[T]) Push(value T) {
	l.AddLast(value)
}

// Top Returns the value at the top of the stack without removing it. Returns error if stack is empty.
// Equivalent to GetLast()
func (l *LinkedList[T]) Top() (T, error) {
	return l.GetLast()
}

//-----------------------------//
//       Queue Functions       //
//-----------------------------//

// NewQueue Creates a new Queue
func NewQueue[T any]() Queue[T] {
	return new(LinkedList[T])
}

// Enqueue Adds value to the end of the queue.
// Equivalent to AddLast()
func (l *LinkedList[T]) Enqueue(value T) {
	l.AddLast(value)
}

// Dequeue Removes the value at the front of the queue and retrieves the value. Returns error if queue is empty.
// Equivalent to RemoveFirst()
func (l *LinkedList[T]) Dequeue() (T, error) {
	return l.RemoveFirst()
}

// Peek Returns the value at the front of the queue without removing it. Returns error if queue is empty.
// Equivalent to GetFirst()
func (l *LinkedList[T]) Peek() (T, error) {
	return l.GetFirst()
}
