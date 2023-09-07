package list

import (
	"errors"
)

type ArrayList[T any] struct {
	arr        []T
	Comparator func(T, T) int
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
		return ErrorOutOfBounds
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

func (l *ArrayList[T]) AddAllAt(index int, arr ...T) error {
	if index < 0 || index > l.Size() {
		return ErrorOutOfBounds
	}
	if len(arr) == 0 {
		return errors.New("NoElementsInArray")
	}
	tempArr := []T{}
	tempArr = append(tempArr, l.arr[:index]...)
	tempArr = append(tempArr, arr...)
	tempArr = append(tempArr, l.arr[index:]...)
	l.arr = tempArr
	return nil
}

func (l *ArrayList[T]) GetFirst() (T, error) {
	if l.IsEmpty() {
		return *new(T), ErrorNoSuchElement
	}
	return l.arr[0], nil
}

func (l *ArrayList[T]) GetLast() (T, error) {
	if l.IsEmpty() {
		return *new(T), ErrorNoSuchElement
	}

	return l.arr[l.Size()-1], nil
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

func (l *ArrayList[T]) Contains(value T) (bool, error) {
	if l.Comparator == nil {
		return false, ErrorNoComparator
	}
	for _, v := range l.arr {
		if l.Comparator(v, value) == 0 {
			return true, nil
		}
	}
	return false, nil
}

func (l *ArrayList[T]) IndexOf(value T) (int, error) {
	if l.Comparator == nil {
		return -1, ErrorNoComparator
	}
	for i, v := range l.arr {
		if l.Comparator(v, value) == 0 {
			return i, nil
		}
	}
	return -1, nil
}

func (l *ArrayList[T]) RemoveElement(value T) bool {
	index, err := l.IndexOf(value)
	if err != nil {
		return false
	}
	_, err = l.Remove(index)
	return err == nil
}

func (l *ArrayList[T]) ForEach(f func(T) T) {
	if f == nil {
		return
	}
	for i, v := range l.arr {
		l.arr[i] = f(v)
	}
}

func (l *ArrayList[T]) Map(f func(T) T) List[T] {
	if f == nil {
		return NewArrayList[T](l.Comparator)
	}
	arrMap := make([]T, l.Size())
	for i, v := range l.arr {
		arrMap[i] = f(v)
	}
	return &ArrayList[T]{arrMap, l.Comparator}
}

func (l *ArrayList[T]) Filter(f func(T) bool) List[T] {
	if f == nil {
		return NewArrayList[T](l.Comparator)
	}
	arrFilter := make([]T, 0, l.Size())
	for _, v := range l.arr {
		if f(v) {
			arrFilter = append(arrFilter, v)
		}
	}
	return &ArrayList[T]{arrFilter, l.Comparator}
}

func (l *ArrayList[T]) ToSlice() []T {
	arr := make([]T, l.Size())
	copy(arr, l.arr)
	return arr
}

func (l *ArrayList[T]) Clear() {
	l.arr = make([]T, 0)
}
