package list

import (
	"errors"
)

var (
	ErrorNoSuchElement     = errors.New("NoSuchElement")
	ErrorOutOfBounds       = errors.New("OutOfBounds")
	ErrorNoElementsInArray = errors.New("NoElementsInArray")
	ErrorNoComparator      = errors.New("NoComparator")
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
	Append(T)
	Add(int, T) error
	AddAll(...T) error
	AddAllAt(int, ...T) error
	RemoveFirst() (T, error)
	RemoveLast() (T, error)
	Remove(int) (T, error)
	RemoveElement(T, Comparator[T]) bool
	Contains(T, Comparator[T]) (bool, error)
	IndexOf(T, Comparator[T]) (int, error)
	ForEach(func(T) T)
	Map(func(T) T) List[T]
	Filter(func(T) bool) List[T]
	ToSlice() []T
	Clear()
}

// Comparator function should return
// 0 if ir is equal;
// 0>(Greater than 0) if it v1 is greater than v2;
// 0<(Less than 0) if it v1 is lower than v2;
type Comparator[T any] func(T, T) int

// NewList Creates a new LinkedList
// Comparator function should return
// 0 if ir is equal
// 0>(Greater than 0) if it v1 is greater than v2
// 0<(Less than 0) if it v1 is lower than v2
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{nil, nil, 0}
}

// LinkedListFromArray Creates a new LinkedList from the given array
func LinkedListFromArray[T any](arr ...T) *LinkedList[T] {
	l := NewLinkedList[T]()
	if len(arr) == 0 {
		return l
	}

	l.AddAll(arr...)
	return l
}

// NewStack Creates a new Stack
func NewStack[T any]() Stack[T] {
	return new(LinkedList[T])
}

// NewQueue Creates a new Queue
func NewQueue[T any]() Queue[T] {
	return new(LinkedList[T])
}

func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{make([]T, 0)}
}

func ArrayListFromArray[T any](arr ...T) *ArrayList[T] {
	l := NewArrayList[T]()
	if len(arr) == 0 {
		return l
	}

	l.AddAll(arr...)
	return l
}

func NewList[T any](t string) List[T] {
	if t == "LinkedList" {
		return NewLinkedList[T]()
	}
	return NewArrayList[T]()
}

func intComp(v1 int, v2 int) int {
	return v1 - v2
}
