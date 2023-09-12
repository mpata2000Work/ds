package heap

import "testing"

func cmp(v1 int, v2 int) int {
	return v1 - v2
}

func TestCanCreateEmptyMinHeap(t *testing.T) {
	heap, err := NewHeap[int](cmp, MIN_HEAP)
	if err != nil {
		t.Error("Error creating empty min heap")
	}
	if heap.Size() != 0 {
		t.Error("Empty min heap should have size 0")
	}
	if !heap.IsEmpty() {
		t.Error("Empty min heap should be empty")
	}
	if heap.IsMaxHeap() {
		t.Error("Empty min heap should not be max heap")
	}
}

func TestCanCreateEmptyMaxHeap(t *testing.T) {
	heap, err := NewHeap[int](cmp, MAX_HEAP)
	if err != nil {
		t.Error("Error creating empty max heap")
	}
	if heap.Size() != 0 {
		t.Error("Empty max heap should have size 0")
	}
	if !heap.IsEmpty() {
		t.Error("Empty max heap should be empty")
	}
	if heap.IsMinHeap() {
		t.Error("Empty max heap should not be min heap")
	}
}

func TestNewMinHeapIsMinHeap(t *testing.T) {
	heap, _ := NewMinHeap[int](cmp)
	if !heap.IsMinHeap() {
		t.Error("New min heap should be min heap")
	}
}

func TestNewMaxHeapIsMaxHeap(t *testing.T) {
	heap, _ := NewMaxHeap[int](cmp)
	if !heap.IsMaxHeap() {
		t.Error("New max heap should be max heap")
	}
}

func TestCannotCreateHeapWithNilComparator(t *testing.T) {
	_, err := NewHeap[int](nil, MIN_HEAP)
	if err == nil {
		t.Error("Should not be able to create heap with nil comparator")
	}
}

func TestCannotCreateHeapWithInvalidHeapType(t *testing.T) {
	_, err := NewHeap[int](cmp, 2)
	if err == nil {
		t.Error("Should not be able to create heap with invalid heap type")
	}
}

func TestCanAddElementToEmptyMinHeap(t *testing.T) {
	heap := heap[int]{make([]int, 0), MIN_HEAP, cmp}
	heap.Push(1)
	if heap.Size() != 1 {
		t.Error("Min heap should have size 1")
	}
	if heap.IsEmpty() {
		t.Error("Min heap should not be empty")
	}
	if heap.arr[0] != 1 {
		t.Error("Min heap should have 1 as first element")
	}
}

func TestCanAddElementToEmptyMaxHeap(t *testing.T) {
	heap := heap[int]{make([]int, 0), MAX_HEAP, cmp}
	heap.Push(1)
	if heap.Size() != 1 {
		t.Error("Max heap should have size 1")
	}
	if heap.IsEmpty() {
		t.Error("Max heap should not be empty")
	}
	if heap.arr[0] != 1 {
		t.Error("Max heap should have 1 as first element")
	}
}

func TestPeekInEmptyHeap(t *testing.T) {
	heap, _ := NewHeap[int](cmp, MIN_HEAP)
	_, err := heap.Peek()
	if err == nil {
		t.Error("Should not be able to peek in empty heap")
	}
}

func TestPushMultipleValuesMinHeap(t *testing.T) {
	heap := heap[int]{make([]int, 0), MIN_HEAP, cmp}
	heap.Push(1)
	heap.Push(2)
	heap.Push(3)
	heap.Push(4)
	heap.Push(5)
	if heap.Size() != 5 {
		t.Error("Min heap should have size 5")
	}
	if heap.IsEmpty() {
		t.Error("Min heap should not be empty")
	}
	if heap.arr[0] != 1 {
		t.Error("Min heap should have 1 as first element")
	}
}

func TestPushMultipleValuesMaxHeap(t *testing.T) {
	heap := heap[int]{make([]int, 0), MAX_HEAP, cmp}
	heap.Push(1)
	heap.Push(2)
	heap.Push(3)
	heap.Push(4)
	heap.Push(5)
	if heap.Size() != 5 {
		t.Error("Max heap should have size 5")
	}
	if heap.IsEmpty() {
		t.Error("Max heap should not be empty")
	}
	if heap.arr[0] != 5 {
		t.Error("Max heap should have 5 as first element")
	}
}

func TestPeekInHeap(t *testing.T) {
	heap, _ := NewHeap[int](cmp, MIN_HEAP)
	heap.Push(1)
	heap.Push(2)
	heap.Push(3)
	heap.Push(4)
	heap.Push(5)
	value, err := heap.Peek()
	if err != nil {
		t.Error("Should be able to peek in heap")
	}
	if value != 1 {
		t.Error("Peek should return 1")
	}
}

func TestPopInEmptyHeap(t *testing.T) {
	heap, _ := NewHeap[int](cmp, MIN_HEAP)
	_, err := heap.Pop()
	if err == nil {
		t.Error("Should not be able to pop in empty heap")
	}
}

func TestPopInMinHeapValuesInOrder(t *testing.T) {
	heap, _ := NewHeap[int](cmp, MIN_HEAP)
	arrIn := []int{0, 1, 2, 3, 4, 5}
	arrExpected := []int{0, 1, 2, 3, 4, 5}
	for _, v := range arrIn {
		heap.Push(v)
	}
	for _, v := range arrExpected {
		value, err := heap.Pop()
		if err != nil {
			t.Error("Should be able to pop in heap")
		}
		if value != v {
			t.Errorf("Pop should return %d but returned %d", v, value)
		}
	}
}

func TestPopInMaxHeapValuesInOrder(t *testing.T) {
	heap, _ := NewHeap[int](cmp, MAX_HEAP)
	arrIn := []int{0, 1, 2, 3, 4, 5}
	arrExpected := []int{5, 4, 3, 2, 1, 0}
	for _, v := range arrIn {
		heap.Push(v)
	}
	for _, v := range arrExpected {
		value, err := heap.Pop()
		if err != nil {
			t.Error("Should be able to pop in heap")
		}
		if value != v {
			t.Errorf("Pop should return %d but returned %d", v, value)
		}
	}
}

func TestPopInMinHeapValuesInOrderWithDuplicates(t *testing.T) {
	heap, _ := NewHeap[int](cmp, MIN_HEAP)
	arrIn := []int{0, 1, 2, 3, 4, 5, 1, 2, 3, 4}
	arrExpected := []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5}
	for _, v := range arrIn {
		heap.Push(v)
	}
	for _, v := range arrExpected {
		value, err := heap.Pop()
		if err != nil {
			t.Error("Should be able to pop in heap")
		}
		if value != v {
			t.Errorf("Pop should return %d but returned %d", v, value)
		}
	}
}

func TestPopInMaxHeapValuesInOrderWithDuplicates(t *testing.T) {
	heap, _ := NewHeap[int](cmp, MAX_HEAP)
	arrIn := []int{0, 1, 2, 3, 4, 5, 1, 2, 3, 4}
	arrExpected := []int{5, 4, 4, 3, 3, 2, 2, 1, 1, 0}
	for _, v := range arrIn {
		heap.Push(v)
	}
	for _, v := range arrExpected {
		value, err := heap.Pop()
		if err != nil {
			t.Error("Should be able to pop in heap")
		}
		if value != v {
			t.Errorf("Pop should return %d but returned %d", v, value)
		}
	}
}

func TestPopInMinHeapValuesInOrderWithDuplicatesAndRandomOrder(t *testing.T) {
	heap, _ := NewHeap[int](cmp, MIN_HEAP)
	arrIn := []int{0, 1, 2, 3, 4, 5, 1, 2, 3, 4}
	arrExpected := []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5}
	for _, v := range arrIn {
		heap.Push(v)
	}
	for _, v := range arrExpected {
		value, err := heap.Pop()
		if err != nil {
			t.Error("Should be able to pop in heap")
		}
		if value != v {
			t.Errorf("Pop should return %d but returned %d", v, value)
		}
	}
}

func TestPopInMaxHeapValuesInOrderWithDuplicatesAndRandomOrder(t *testing.T) {
	heap, _ := NewHeap[int](cmp, MAX_HEAP)
	arrIn := []int{0, 1, 2, 3, 4, 5, 1, 2, 3, 4}
	arrExpected := []int{5, 4, 4, 3, 3, 2, 2, 1, 1, 0}
	for _, v := range arrIn {
		heap.Push(v)
	}
	for _, v := range arrExpected {
		value, err := heap.Pop()
		if err != nil {
			t.Error("Should be able to pop in heap")
		}
		if value != v {
			t.Errorf("Pop should return %d but returned %d", v, value)
		}
	}
}

func TestHeapifyNilComparator(t *testing.T) {
	_, err := Heapify[int](nil, MIN_HEAP, 1, 2, 3, 4, 5)
	if err == nil {
		t.Error("Should not be able to heapify with nil comparator")
	}
}

func TestHeapifyInvalidHeapType(t *testing.T) {
	_, err := Heapify[int](cmp, 2, 1, 2, 3, 4, 5)
	if err == nil {
		t.Error("Should not be able to heapify with invalid heap type")
	}
}

func TestHeapifyEmptyArray(t *testing.T) {
	heap, _ := Heapify[int](cmp, MIN_HEAP)
	if heap.Size() != 0 {
		t.Error("Heapify should return an empty heap")
	}
}

func TestHeapifyMinHeap(t *testing.T) {
	arrIn := []int{0, 1, 2, 3, 4, 5}
	arrExpected := []int{0, 1, 2, 3, 4, 5}
	heap, _ := Heapify[int](cmp, MIN_HEAP, arrIn...)
	for _, v := range arrExpected {
		value, err := heap.Pop()
		if err != nil {
			t.Error("Should be able to pop in heap")
		}
		if value != v {
			t.Errorf("Pop should return %d but returned %d", v, value)
		}
	}
}

func TestHeapifyMaxHeap(t *testing.T) {
	arrIn := []int{0, 1, 2, 3, 4, 5}
	arrExpected := []int{5, 4, 3, 2, 1, 0}
	heap, _ := Heapify[int](cmp, MAX_HEAP, arrIn...)
	for _, v := range arrExpected {
		value, err := heap.Pop()
		if err != nil {
			t.Error("Should be able to pop in heap")
		}
		if value != v {
			t.Errorf("Pop should return %d but returned %d", v, value)
		}
	}
}

func TestHeapSortInMinHeap(t *testing.T) {
	arrIn := []int{0, 1, 2, 3, 4, 5}
	arrExpected := []int{0, 1, 2, 3, 4, 5}
	arrSorted, _ := HeapSort[int](cmp, false, arrIn...)
	for i, v := range arrExpected {
		if arrSorted[i] != v {
			t.Errorf("Heap sort should return %d but returned %d", v, arrSorted[i])
		}
	}
}

func TestHeapSortInMaxHeap(t *testing.T) {
	arrIn := []int{0, 1, 2, 3, 4, 5}
	arrExpected := []int{5, 4, 3, 2, 1, 0}
	arrSorted, _ := HeapSort[int](cmp, true, arrIn...)
	for i, v := range arrExpected {
		if arrSorted[i] != v {
			t.Errorf("Heap sort should return %d but returned %d", v, arrSorted[i])
		}
	}
}

func TestHeapSOrtNilComparator(t *testing.T) {
	_, err := HeapSort[int](nil, false, 1, 2, 3, 4, 5)
	if err == nil {
		t.Error("Should not be able to heap sort with nil comparator")
	}
}
