package heap

import "errors"

// Comparator function should return
// 0 if ir is equal;
// 0>(Greater than 0) if it v1 is greater than v2;
// 0<(Less than 0) if it v1 is lower than v2;
type Comparator[T any] func(T, T) int

type heap[T any] struct {
	arr      []T
	heapType int
	cmp      Comparator[T]
}

type Heap[T any] interface {
	// IsMaxHeap Returns true if the heap is a MaxHeap
	IsMaxHeap() bool
	// IsMinHeap Returns true if the heap is a MinHeap
	IsMinHeap() bool
	// IsEmpty Returns true if the heap is empty
	IsEmpty() bool
	// Size Returns the size of the heap
	Size() int
	// Peek Returns the value of the first element of the heap or error if heap is empty
	Peek() (T, error)
	// Push Pushes the given value to the heap
	Push(T)
	// Pop Removes the first element of the heap and returns its value or error if heap is empty
	Pop() (T, error)
}

const MIN_HEAP = 1
const MAX_HEAP = -1

var (
	ErrorComparatorIsNil     = errors.New("ErrorComparatorIsNil")
	ErrorNoSuchElement       = errors.New("ErrorNoSuchElement")
	ErrorHeapTypeDoesntExist = errors.New("ErrorHeapTypeDoesntExist")
)

// fatherPosition Returns the position of the father of the element at the given index
func fatherPosition(i int) int {
	return (i - 1) / 2
}

// leftSonPosition Returns the position of the left son of the element at the given index
func leftSonPosition(i int) int {
	return (i * 2) + 1
}

// rightSonPosition Returns the position of the right son of the element at the given index
func rightSonPosition(i int) int {
	return (i * 2) + 2
}

// NewHeap Creates a new Heap
// Comparator function should return
// 0 if ir is equal;
// 0>(Greater than 0) if it v1 is greater than v2;
// 0<(Less than 0) if it v1 is lower than v2;
// heapType should be MIN_HEAP or MAX_HEAP
func NewHeap[T any](cmp Comparator[T], heapType int) (Heap[T], error) {
	if cmp == nil {
		return nil, ErrorComparatorIsNil
	}
	if heapType != MIN_HEAP && heapType != MAX_HEAP {
		return nil, ErrorHeapTypeDoesntExist
	}
	return &heap[T]{make([]T, 0), heapType, cmp}, nil
}

// NewMinHeap Creates a new MinHeap
// Comparator function should return
// 0 if ir is equal;
// 0>(Greater than 0) if it v1 is greater than v2;
// 0<(Less than 0) if it v1 is lower than v2;
func NewMinHeap[T any](cmp Comparator[T]) (Heap[T], error) {
	return NewHeap[T](cmp, MIN_HEAP)
}

// NewMaxHeap Creates a new MaxHeap
// Comparator function should return
// 0 if ir is equal;
// 0>(Greater than 0) if it v1 is greater than v2;
// 0<(Less than 0) if it v1 is lower than v2;
func NewMaxHeap[T any](cmp Comparator[T]) (Heap[T], error) {
	return NewHeap[T](cmp, MAX_HEAP)
}

// Heapify Creates a new Heap from the given array
// Comparator function should return
// 0 if ir is equal;
// 0>(Greater than 0) if it v1 is greater than v2;
// 0<(Less than 0) if it v1 is lower than v2;
// heapType should be MIN_HEAP or MAX_HEAP
func Heapify[T any](cmp Comparator[T], heapType int, arr ...T) (Heap[T], error) {
	if cmp == nil {
		return nil, ErrorComparatorIsNil
	}
	if heapType != MIN_HEAP && heapType != MAX_HEAP {
		return nil, ErrorHeapTypeDoesntExist
	}
	h := &heap[T]{arr, heapType, cmp}
	for i := h.Size() - 1; i >= 0; i-- {
		h.siftDown(i)
	}
	return h, nil
}

// HeapifyMin Creates a new MinHeap from the given array
// Comparator function should return
// 0 if ir is equal;
// 0>(Greater than 0) if it v1 is greater than v2;
// 0<(Less than 0) if it v1 is lower than v2;
// reversed should be true if you want greater elements to be first
func HeapSort[T any](cmp Comparator[T], reversed bool, arr ...T) ([]T, error) {
	if cmp == nil {
		return nil, ErrorComparatorIsNil
	}
	hType := MIN_HEAP
	if reversed {
		hType = MAX_HEAP
	}
	h, err := Heapify[T](cmp, hType, arr...)
	if err != nil {
		return nil, err
	}
	sorted := make([]T, 0, h.Size())
	for !h.IsEmpty() {
		val, _ := h.Pop()
		sorted = append(sorted, val)
	}
	return sorted, nil
}

// HeapifyMin Creates a new MinHeap from the given array
func (h *heap[T]) IsMaxHeap() bool {
	return h.heapType == MAX_HEAP
}

// HeapifyMin Creates a new MinHeap from the given array
func (h *heap[T]) IsMinHeap() bool {
	return h.heapType == MIN_HEAP
}

// IsEmpty Returns true if the heap is empty
func (h *heap[T]) IsEmpty() bool {
	return len(h.arr) == 0
}

// Size Returns the size of the heap
func (h *heap[T]) Size() int {
	return len(h.arr)
}

// swap Swaps the elements at the given positions
func (h *heap[T]) swap(i int, j int) {
	if i >= h.Size() || j >= h.Size() {
		return
	}
	aux := h.arr[i]
	h.arr[i] = h.arr[j]
	h.arr[j] = aux
}

// compare returns 0 if ir is equal;
// 0>(Greater than 0) if it v1 is greater than v2;
// 0<(Less than 0) if it v1 is lower than v2;
// if heapType is MAX_HEAP, it returns the opposite
func (h *heap[T]) compare(v1 T, v2 T) int {
	return h.cmp(v1, v2) * h.heapType
}

// siftUp Sifts up the element at the given index
func (h *heap[T]) siftUp(i int) {
	if i == 0 || h.IsEmpty() {
		return
	}

	father := fatherPosition(i)

	if h.compare(h.arr[i], h.arr[father]) < 0 {
		h.swap(i, father)
		h.siftUp(father)
	}
}

// siftDown Sifts down the element at the given index
func (h *heap[T]) siftDown(i int) {
	if h.IsEmpty() {
		return
	}

	left := leftSonPosition(i)
	least := left
	if left >= h.Size() {
		return
	}

	if right := rightSonPosition(i); right < h.Size() {
		if h.compare(h.arr[left], h.arr[right]) > 0 {
			least = right
		}
	}

	if h.compare(h.arr[i], h.arr[least]) > 0 {
		h.swap(i, least)
		h.siftDown(least)
	}

}

// Peek Returns the value of the first element of the heap or error if heap is empty
func (h *heap[T]) Peek() (T, error) {
	if h.IsEmpty() {
		return *new(T), ErrorNoSuchElement
	}
	return h.arr[0], nil
}

// Push Pushes the given value to the heap
func (h *heap[T]) Push(value T) {
	h.arr = append(h.arr, value)
	h.siftUp(h.Size() - 1)
}

// Pop Removes the first element of the heap and returns its value or error if heap is empty
func (h *heap[T]) Pop() (T, error) {
	if h.IsEmpty() {
		return *new(T), ErrorNoSuchElement
	}
	val := h.arr[0]
	h.swap(0, h.Size()-1)
	h.arr = h.arr[:h.Size()-1]
	if !h.IsEmpty() {
		h.siftDown(0)
	}

	return val, nil
}
