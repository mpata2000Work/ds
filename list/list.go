package list

import (
	"errors"
)

var (
	ErrorNoSuchElement     = errors.New("NoSuchElement")
	ErrorOutOfBounds       = errors.New("OutOfBounds")
	ErrorNoElementsInArray = errors.New("NoElementsInArray")
)

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

type List[T any] interface {
	IsEmpty() bool
	Size() int
	Get(int) (T, error)
	GetFirst() (T, error)
	GetLast() (T, error)
	AddFirst(T)
	AddLast(T)
	Add(int, T) error
	AddAll(...T) error
	AddAllAt(int, ...T) error
	RemoveFirst() (T, error)
	RemoveLast() (T, error)
	Remove(int) (T, error)
	RemoveElement(T) bool
	Contains(T) bool
	IndexOf(T) int
	ForEach(func(T) T)
	Map(func(T) T) List[T]
	PrettyPrint()
}

// NewList Creates a new LinkedList
// Comparator function should return
// 0 if ir is equal
// 0>(Greater than 0) if it v1 is greater than v2
// 0<(Less than 0) if it v1 is lower than v2
func NewLinkedList[T any](comparator func(T, T) int) LinkedList[T] {
	return LinkedList[T]{nil, nil, 0, comparator}
}
