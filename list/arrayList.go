package list

import (
	"errors"
)

type ArrayList[T any] struct {
	arr        []T
	Comparator func(T, T) int
}

func NewArrayList[T any](comp func(T, T) int) ArrayList[T] {
	return ArrayList[T]{make([]T, 0), comp}
}

func (l *ArrayList[T]) Size() int {
	return len(l.arr)
}

func (l *ArrayList[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *ArrayList[T]) AddFirst(value T) {
	l.arr = append([]T{value}, l.arr...)
}

func (l *ArrayList[T]) AddLast(value T) {
	l.arr = append(l.arr, value)
}

func (l *ArrayList[T]) Add(index int, value T) error {
	if index < 0 || index > l.Size() {
		return errors.New("IndexOutOfBounds")
	}
	if index == 0 {
		l.AddFirst(value)
		return nil
	}
	if index == l.Size() {
		l.AddLast(value)
		return nil
	}
	tempArr := []T{value}
	tempArr = append(tempArr, l.arr[index:]...)
	l.arr = append(l.arr[:index], tempArr...)
	return nil
}

func (l *ArrayList[T]) AddAll(arr ...T) error {
	if len(arr) == 0 {
		return errors.New("NoElementsInArray")
	}
	l.arr = append(l.arr, arr...)
	return nil
}

func (l *ArrayList[T]) Get(index int) (T, error) {
	if index < 0 || index >= l.Size() {
		return *new(T), ErrorOutOfBounds
	}
	return l.arr[index], nil
}

func (l *ArrayList[T]) Set(index int, value T) error {
	if index < 0 || index >= l.Size() {
		return ErrorOutOfBounds
	}
	l.arr[index] = value
	return nil
}

func (l *ArrayList[T]) RemoveFirst() (T, error) {
	if l.IsEmpty() {
		return *new(T), ErrorNoSuchElement
	}
	value := l.arr[0]
	l.arr = l.arr[1:]
	return value, nil
}

func (l *ArrayList[T]) RemoveLast() (T, error) {
	if l.IsEmpty() {
		return *new(T), ErrorNoSuchElement
	}
	value := l.arr[l.Size()-1]
	l.arr = l.arr[:l.Size()-1]
	return value, nil
}

func (l *ArrayList[T]) Remove(index int) (T, error) {
	if index < 0 || index >= l.Size() {
		return *new(T), ErrorOutOfBounds
	}
	if index == 0 {
		return l.RemoveFirst()
	}
	if index == l.Size()-1 {
		return l.RemoveLast()
	}
	value := l.arr[index]
	l.arr = append(l.arr[:index], l.arr[index+1:]...)
	return value, nil
}

func (l *ArrayList[T]) Contains(value T) bool {
	for _, v := range l.arr {
		if l.Comparator(v, value) == 0 {
			return true
		}
	}
	return false
}

func (l *ArrayList[T]) IndexOf(value T) int {
	for i, v := range l.arr {
		if l.Comparator(v, value) == 0 {
			return i
		}
	}
	return -1
}

func (l *ArrayList[T]) RemoveElement(value T) (T, error) {
	index := l.IndexOf(value)
	if index == -1 {
		return *new(T), ErrorNoSuchElement
	}
	return l.Remove(index)
}
