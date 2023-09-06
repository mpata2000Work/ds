package list

import (
	"testing"
)

func comp(v1 int, v2 int) int {
	return v1 - v2
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

func TestAddLasttInEmptyList(t *testing.T) {
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
