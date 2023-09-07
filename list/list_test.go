package list

import (
	"testing"
)

func comp(v1 int, v2 int) int {
	return v1 - v2
}

func listOfFiveInts() *LinkedList[int] {
	l := NewList[int](comp)
	l.AddLast(0)
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)
	l.AddLast(4)
	return l
}

func TestCreateEmptyListWithoutComparator(t *testing.T) {
	l := NewList[int](nil)
	if l == nil || l.head != nil || l.tail != nil || l.size != 0 || l.Comparator != nil {
		t.Error("List wasnt correctly initialize ", l)
	}
}

func TestCreateEmptyListWithComparator(t *testing.T) {
	l := NewList[int](comp)
	if l == nil || l.head != nil || l.tail != nil || l.size != 0 || l.Comparator == nil {
		t.Error("List wasnt correctly initialize ", l)
	}
}

func TestAddFirstInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	l.AddFirst(0)
	if l.head == nil {
		t.Error("Head want initialize correctly")
	}
	if l.tail == nil {
		t.Error("Tail want initialize correctly")
	}
	if l.size != 1 {
		t.Error("Size wasnt updated Expected 1 got ", l.size)
	}
	if l.head.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.head.value)
	}
}

func TestAddLastInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	l.AddLast(0)
	if l.head == nil {
		t.Error("Head want initialize correctly")
	}
	if l.tail == nil {
		t.Error("Tail want initialize correctly")
	}
	if l.size != 1 {
		t.Error("Size wasnt updated Expected 1 got ", l.size)
	}
	if l.head.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.head.value)
	}
}

func TestAddFirstInList(t *testing.T) {
	l := NewList[int](comp)
	l.AddFirst(0)
	l.AddFirst(1)
	if l.head == nil {
		t.Error("Head want initialize correctly")
	}
	if l.tail == nil {
		t.Error("Tail want initialize correctly")
	}
	if l.size != 2 {
		t.Error("Size wasnt updated Expected 2 got ", l.size)
	}
	if l.head.value != 1 {
		t.Error("Value isnt correct Expected 1 got ", l.head.value)
	}
	if l.tail.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.tail.value)
	}
}

func TestAddLastInList(t *testing.T) {
	l := NewList[int](comp)
	l.AddLast(0)
	l.AddLast(1)
	if l.head == nil {
		t.Error("Head want initialize correctly")
	}
	if l.tail == nil {
		t.Error("Tail want initialize correctly")
	}
	if l.size != 2 {
		t.Error("Size wasnt updated Expected 2 got ", l.size)
	}
	if l.head.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.head.value)
	}
	if l.tail.value != 1 {
		t.Error("Value isnt correct Expected 1 got ", l.tail.value)
	}
}

func TestAddInListEmptyListAtIndexCero(t *testing.T) {
	l := NewList[int](comp)
	l.Add(0, 0)
	if l.head == nil {
		t.Error("Head want initialize correctly")
	}
	if l.tail == nil {
		t.Error("Tail want initialize correctly")
	}
	if l.size != 1 {
		t.Error("Size wasnt updated Expected 1 got ", l.size)
	}
	if l.head.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.head.value)
	}
	if l.tail.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.tail.value)
	}
}

func TestAddInListEmptyListAtIndexOut(t *testing.T) {
	l := NewList[int](comp)
	err := l.Add(1, 0)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestAddInListAtIndexCero(t *testing.T) {
	l := listOfFiveInts()
	l.Add(0, 5)
	if l.head == nil {
		t.Error("Head want initialize correctly")
	}
	if l.tail == nil {
		t.Error("Tail want initialize correctly")
	}
	if l.size != 6 {
		t.Error("Size wasnt updated Expected 6 got ", l.size)
	}
	if l.head.value != 5 {
		t.Error("Value isnt correct Expected 5 got ", l.head.value)
	}
	if l.tail.value != 4 {
		t.Error("Value isnt correct Expected 4 got ", l.tail.value)
	}
}

func TestAddInListAtLastIndex(t *testing.T) {
	l := listOfFiveInts()
	err := l.Add(5, 5)
	if err == nil {
		t.Error("Error wasnt raised")
	}
	if l.head == nil {
		t.Error("Head want initialize correctly")
	}
	if l.tail == nil {
		t.Error("Tail want initialize correctly")
	}
	if l.size != 6 {
		t.Error("Size wasnt updated Expected 6 got ", l.size)
	}
	if l.head.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.head.value)
	}
	if l.tail.value != 5 {
		t.Error("Value isnt correct Expected 5 got ", l.tail.value)
	}
}

func TestAddInListAtIndex(t *testing.T) {
	l := listOfFiveInts()
	l.Add(2, 5)
	if l.head == nil {
		t.Error("Head want initialize correctly")
	}
	if l.tail == nil {
		t.Error("Tail want initialize correctly")
	}
	if l.size != 6 {
		t.Error("Size wasnt updated Expected 6 got ", l.size)
	}
	if l.head.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.head.value)
	}
	if l.tail.value != 4 {
		t.Error("Value isnt correct Expected 4 got ", l.tail.value)
	}
	if l.head.next.value != 1 || l.head.next.next.value != 5 || l.head.next.next.next.value != 2 {
		t.Error("Value isnt correct Expected 1 5 2 got ", l.head.next.value, l.head.next.next.value, l.head.next.next.next.value)
	}
}

func TestGetFirstInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	_, err := l.GetFirst()
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestGetFirstInList(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.GetFirst()
	if err != nil {
		t.Error("Error was raised")
	}
	if v != 0 {
		t.Error("Value isnt correct Expected 0 got ", v)
	}
}

func TestGetLastInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	_, err := l.GetLast()
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestGetLastInList(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.GetLast()
	if err != nil {
		t.Error("Error was raised")
	}
	if v != 4 {
		t.Error("Value isnt correct Expected 4 got ", v)
	}
}

func TestGetAtInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	_, err := l.GetAt(0)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestGetAtInList(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.GetAt(2)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != 2 {
		t.Error("Value isnt correct Expected 2 got ", v)
	}
}

func TestGetAtInListOutOfBounds(t *testing.T) {
	l := listOfFiveInts()
	_, err := l.GetAt(5)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestGetAtInListOutOfBoundsNegative(t *testing.T) {
	l := listOfFiveInts()
	_, err := l.GetAt(-1)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}
