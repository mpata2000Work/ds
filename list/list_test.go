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
	if l.head.value != 1 || l.head.next.value != 0 {
		t.Error("Value isnt correct Expected 1 0 got ", l.head.value, l.head.next.value)
	}
	if l.tail.value != 0 || l.tail.previous.value != 1 {
		t.Error("Value isnt correct Expected 0 1 got ", l.tail.value, l.tail.previous.value)
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
	if err != nil {
		t.Error("Error was raised")
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
	err := l.Add(4, 2)

	if err != nil {
		t.Error("Error was raised")
	}
	if l.head == nil || l.tail == nil {
		t.Error("Head or tail want initialize correctly")
	}
	if l.size != 6 {
		t.Error("Size wasnt updated Expected 6 got ", l.size)
	}
	if l.head.value != 0 || l.tail.value != 4 {
		t.Error("Value isnt correct for Head or Tail Expected 0 4 got ", l.head.value, l.tail.value)
	}
	nodeFromTail := l.tail.previous            // Index:5->4
	nodeFromHead := l.head.next.next.next.next // Index:0->1->2->3->4
	if nodeFromTail != nodeFromHead {
		t.Error("Node from tail and head arent the same")
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

func TestRemoveFirstInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	_, err := l.RemoveFirst()
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestRemoveFirstInList(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.RemoveFirst()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if v != 0 {
		t.Error("Value isnt correct Expected 0 got ", v)
	}
	if l.head.value != 1 {
		t.Error("Value isnt correct Expected 1 got ", l.head.value)
	}
	if l.tail.value != 4 {
		t.Error("Value isnt correct Expected 4 got ", l.tail.value)
	}
	if l.size != 4 {
		t.Error("Size wasnt updated Expected 4 got ", l.size)
	}
}

func TestRemoveLastInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	_, err := l.RemoveLast()
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestRemoveLastInList(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.RemoveLast()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if v != 4 {
		t.Error("Value isnt correct Expected 4 got ", v)
	}
	if l.head.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.head.value)
	}
	if l.tail.value != 3 {
		t.Error("Value isnt correct Expected 3 got ", l.tail.value)
	}
	if l.size != 4 {
		t.Error("Size wasnt updated Expected 4 got ", l.size)
	}
}

func TestRemoveInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	_, err := l.Remove(0)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestRemoveInList(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.Remove(2)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if v != 2 {
		t.Error("Value isnt correct Expected 2 got ", v)
	}
	if l.head.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.head.value)
	}
	if l.tail.value != 4 {
		t.Error("Value isnt correct Expected 4 got ", l.tail.value)
	}
	if l.size != 4 {
		t.Error("Size wasnt updated Expected 4 got ", l.size)
	}
}

func TestRemoveInListOutOfBounds(t *testing.T) {
	l := listOfFiveInts()
	_, err := l.Remove(5)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestRemoveInListOutOfBoundsNegative(t *testing.T) {
	l := listOfFiveInts()
	_, err := l.Remove(-1)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestRemoveInListAtStart(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.Remove(0)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if v != 0 {
		t.Error("Value isnt correct Expected 0 got ", v)
	}
	if l.head.value != 1 {
		t.Error("Value isnt correct Expected 1 got ", l.head.value)
	}
	if l.tail.value != 4 {
		t.Error("Value isnt correct Expected 4 got ", l.tail.value)
	}
	if l.size != 4 {
		t.Error("Size wasnt updated Expected 4 got ", l.size)
	}
}

func TestRemoveInListAtEnd(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.Remove(4)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if v != 4 {
		t.Error("Value isnt correct Expected 4 got ", v)
	}
	if l.head.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.head.value)
	}
	if l.tail.value != 3 {
		t.Error("Value isnt correct Expected 3 got ", l.tail.value)
	}
	if l.size != 4 {
		t.Error("Size wasnt updated Expected 4 got ", l.size)
	}
}

func TestRemoveFirstInListToEmpty(t *testing.T) {
	l := listOfFiveInts()
	v1, _ := l.RemoveFirst()
	v2, _ := l.RemoveFirst()
	v3, _ := l.RemoveFirst()
	v4, _ := l.RemoveFirst()
	v5, _ := l.RemoveFirst()
	if v1 != 0 || v2 != 1 || v3 != 2 || v4 != 3 || v5 != 4 {
		t.Error("Values arent correct Expected 0 1 2 3 4 got ", v1, v2, v3, v4, v5)
	}
	if l.head != nil {
		t.Error("Head wasnt set to nil")
	}
	if l.tail != nil {
		t.Error("Tail wasnt set to nil")
	}
	if l.size != 0 {
		t.Error("Size wasnt updated Expected 0 got ", l.size)
	}
}

func TestRemoveLastInListToEmpty(t *testing.T) {
	l := listOfFiveInts()
	v1, _ := l.RemoveLast()
	v2, _ := l.RemoveLast()
	v3, _ := l.RemoveLast()
	v4, _ := l.RemoveLast()
	v5, _ := l.RemoveLast()
	if v1 != 4 || v2 != 3 || v3 != 2 || v4 != 1 || v5 != 0 {
		t.Error("Values arent correct Expected 4 3 2 1 0 got ", v1, v2, v3, v4, v5)
	}
	if l.head != nil {
		t.Error("Head wasnt set to nil")
	}
	if l.tail != nil {
		t.Error("Tail wasnt set to nil")
	}
	if l.size != 0 {
		t.Error("Size wasnt updated Expected 0 got ", l.size)
	}
}

func TestRemoveInListToEmpty(t *testing.T) {
	l := listOfFiveInts()
	v1, err1 := l.Remove(4)
	v2, err2 := l.Remove(2)
	v3, err3 := l.Remove(2)
	v4, err4 := l.Remove(0)
	v5, err5 := l.Remove(0)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		t.Error("Error was raised", err1, err2, err3, err4, err5)
	}
	if v1 != 4 || v2 != 2 || v3 != 3 || v4 != 0 || v5 != 1 {
		t.Error("Values arent correct Expected 4 2 3 0 1 got ", v1, v2, v3, v4, v5)
	}
	if l.head != nil {
		t.Error("Head wasnt set to nil")
	}
	if l.tail != nil {
		t.Error("Tail wasnt set to nil")
	}

}
