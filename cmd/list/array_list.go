package list

import (
	"errors"
)

type ArrayList[T any] struct {
	arr []T
}

//------------------------------//
//         List Methods         //
//------------------------------//

// Size Get size of the list
func (l *ArrayList[T]) Size() int {
	return len(l.arr)
}

// IsEmpty Returns true if the list is empty
func (l *ArrayList[T]) IsEmpty() bool {
	return l.Size() == 0
}

// AddFirst Adds the given value to the beginning of the list
func (l *ArrayList[T]) AddFirst(value T) {
	l.arr = append([]T{value}, l.arr...)
}

// Append Adds the given value to the end of the list
func (l *ArrayList[T]) Append(value T) {
	l.arr = append(l.arr, value)
}

// Add Adds the given value to the given index of the list. Returns error if index is out of bounds
func (l *ArrayList[T]) Add(index int, value T) error {
	if index < 0 || index > l.Size() {
		return ErrorOutOfBounds
	}
	if index == 0 {
		l.AddFirst(value)
		return nil
	}
	if index == l.Size() {
		l.Append(value)
		return nil
	}
	tempArr := []T{value}
	tempArr = append(tempArr, l.arr[index:]...)
	l.arr = append(l.arr[:index], tempArr...)
	return nil
}

// AddAll Adds the given values to the end of the list. Returns error if array is empty
func (l *ArrayList[T]) AddAll(arr ...T) error {
	if len(arr) == 0 {
		return errors.New("NoElementsInArray")
	}
	l.arr = append(l.arr, arr...)
	return nil
}

// AddAllAt Adds the given values to the given index of the list
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

// GetFirst Returns the first element of the list
func (l *ArrayList[T]) GetFirst() (T, error) {
	if l.IsEmpty() {
		return *new(T), ErrorNoSuchElement
	}
	return l.arr[0], nil
}

// GetLast Returns the last element of the list
func (l *ArrayList[T]) GetLast() (T, error) {
	if l.IsEmpty() {
		return *new(T), ErrorNoSuchElement
	}

	return l.arr[l.Size()-1], nil
}

// Get Returns the element at the given index of the list. Returns error if index is out of bounds
func (l *ArrayList[T]) Get(index int) (T, error) {
	if index < 0 || index >= l.Size() {
		return *new(T), ErrorOutOfBounds
	}
	return l.arr[index], nil
}

// RemoveFirst Removes the first element of the list and returns its value or error if list is empty
func (l *ArrayList[T]) Set(index int, value T) error {
	if index < 0 || index >= l.Size() {
		return ErrorOutOfBounds
	}
	l.arr[index] = value
	return nil
}

// RemoveFirst Removes the first element of the list and returns its value or error if list is empty
func (l *ArrayList[T]) RemoveFirst() (T, error) {
	if l.IsEmpty() {
		return *new(T), ErrorNoSuchElement
	}
	value := l.arr[0]
	l.arr = l.arr[1:]
	return value, nil
}

// RemoveLast Removes the last element of the list and returns its value or error if list is empty
func (l *ArrayList[T]) RemoveLast() (T, error) {
	if l.IsEmpty() {
		return *new(T), ErrorNoSuchElement
	}
	value := l.arr[l.Size()-1]
	l.arr = l.arr[:l.Size()-1]
	return value, nil
}

// Remove Removes the element at the given index of the list and returns its value or error if index is out of bounds or list is empty
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

// RemoveElement Removes the first occurence of value in the list and returns true or false if it isnt in the list
func (l *ArrayList[T]) RemoveElement(value T, cmp Comparator[T]) bool {
	if cmp == nil {
		return false
	}
	index, _ := l.IndexOf(value, cmp)
	_, err := l.Remove(index)
	return err == nil
}

// Contains Returns true if the list contains the given value
func (l *ArrayList[T]) Contains(value T, cmp Comparator[T]) (bool, error) {
	if cmp == nil {
		return false, ErrorNoComparator
	}
	for _, v := range l.arr {
		if cmp(v, value) == 0 {
			return true, nil
		}
	}
	return false, nil
}

// IndexOf Returns the index of the given value in the list or -1 if the list does not contain the value
func (l *ArrayList[T]) IndexOf(value T, cmp Comparator[T]) (int, error) {
	if cmp == nil {
		return -1, ErrorNoComparator
	}
	for i, v := range l.arr {
		if cmp(v, value) == 0 {
			return i, nil
		}
	}
	return -1, nil
}

// ForEach Applies the given function to each element of the list
func (l *ArrayList[T]) ForEach(f func(T) T) {
	if f == nil {
		return
	}
	for i, v := range l.arr {
		l.arr[i] = f(v)
	}
}

// Map Applies the given function to each element of the list and returns a new list with the results.
// Returns an empty list if f is nil
func (l *ArrayList[T]) Map(f func(T) T) List[T] {
	if f == nil {
		return NewArrayList[T]()
	}
	arrMap := make([]T, l.Size())
	for i, v := range l.arr {
		arrMap[i] = f(v)
	}
	return &ArrayList[T]{arrMap}
}

// Filter Applies the given function to each element of the list and returns a new list with the elements that returned true.
func (l *ArrayList[T]) Filter(f func(T) bool) List[T] {
	if f == nil {
		return NewArrayList[T]()
	}
	arrFilter := make([]T, 0, l.Size())
	for _, v := range l.arr {
		if f(v) {
			arrFilter = append(arrFilter, v)
		}
	}
	return &ArrayList[T]{arrFilter}
}

// ToSlice Returns a slice with the elements of the list
func (l *ArrayList[T]) ToSlice() []T {
	arr := make([]T, l.Size())
	copy(arr, l.arr)
	return arr
}

// Clear Removes all elements from the list
func (l *ArrayList[T]) Clear() {
	l.arr = make([]T, 0)
}

// Copy Returns a new list with the same elements as the list
func (l *ArrayList[T]) Copy() List[T] {
	return &ArrayList[T]{l.ToSlice()}
}

//-----------------------------//
//       Stack Functions       //
//-----------------------------//

// Pop Removes the value at the top of the stack and retrieves the value. Returns error if stack is empty.
// Equivalent to RemoveLast()
func (l *ArrayList[T]) Pop() (T, error) {
	return l.RemoveLast()
}

// Push Adds value to the top of the stack.
// Equivalent to AddLast()
func (l *ArrayList[T]) Push(value T) {
	l.Append(value)
}

// Top Returns the value at the top of the stack without removing it. Returns error if stack is empty.
// Equivalent to GetLast()
func (l *ArrayList[T]) Top() (T, error) {
	return l.GetLast()
}
