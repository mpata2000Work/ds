package list

import (
	"testing"
)

func comp(v1 int, v2 int) int {
	return v1 - v2
}

func listOfFiveInts() LinkedList[int] {
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
	if l.head != nil || l.tail != nil || l.size != 0 || l.Comparator != nil {
		t.Error("List wasnt correctly initialize ", l)
	}
}

func TestCreateEmptyListWithComparator(t *testing.T) {
	l := NewList[int](comp)
	if l.head != nil || l.tail != nil || l.size != 0 || l.Comparator == nil {
		t.Error("List wasnt correctly initialize ", l)
	}
}

func TestIsEmptyInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	if !l.IsEmpty() {
		t.Error("List isnt empty")
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

func TestIsEmptyInList(t *testing.T) {
	l := listOfFiveInts()
	if l.IsEmpty() {
		t.Error("List is empty")
	}
}

func TestSizeInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	if l.Size() != 0 {
		t.Error("Size isnt correct Expected 0 got ", l.Size())
	}
}

func TestSizeInList(t *testing.T) {
	l := listOfFiveInts()
	if l.Size() != 5 {
		t.Error("Size isnt correct Expected 5 got ", l.Size())
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

func TestGetAtInListFirstHalf(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.GetAt(2)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != 2 {
		t.Error("Value isnt correct Expected 2 got ", v)
	}
}

func TestGetAtInListSecondHalf(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.GetAt(4)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != 4 {
		t.Error("Value isnt correct Expected 4 got ", v)
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

func TestClearEmptyList(t *testing.T) {
	l := NewList[int](comp)
	l.Clear()
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

func TestClearList(t *testing.T) {
	l := listOfFiveInts()
	l.Clear()
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

func TestContainsInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	v, err := l.Contains(0)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != false {
		t.Error("Value isnt correct Expected false got ", v)
	}
}

func TestContainsInList(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.Contains(2)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != true {
		t.Error("Value isnt correct Expected true got ", v)
	}
}

func TestContainsInListNotInList(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.Contains(5)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != false {
		t.Error("Value isnt correct Expected false got ", v)
	}
}

func TestContainsWithoutComparator(t *testing.T) {
	l := NewList[int](nil)
	l.AddLast(0)
	_, err := l.Contains(0)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestIndexOfInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	v, err := l.IndexOf(0)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != -1 {
		t.Error("Value isnt correct Expected -1 got ", v)
	}
}

func TestIndexOfInList(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.IndexOf(2)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != 2 {
		t.Error("Value isnt correct Expected 2 got ", v)
	}
}

func TestIndexOfInListNotInList(t *testing.T) {
	l := listOfFiveInts()
	v, err := l.IndexOf(5)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != -1 {
		t.Error("Value isnt correct Expected -1 got ", v)
	}
}

func TestIndexOfWithoutComparator(t *testing.T) {
	l := NewList[int](nil)
	l.AddLast(0)
	_, err := l.IndexOf(0)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestToArrayInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	arr := l.ToArray()
	if len(arr) != 0 {
		t.Error("Array isnt correct Expected [] got ", arr)
	}
}

func TestToArrayInList(t *testing.T) {
	l := listOfFiveInts()
	arr := l.ToArray()
	if len(arr) != 5 {
		t.Error("Array isnt correct Expected [0 1 2 3 4] got ", arr)
	}
	if arr[0] != 0 || arr[1] != 1 || arr[2] != 2 || arr[3] != 3 || arr[4] != 4 {
		t.Error("Array isnt correct Expected [0 1 2 3 4] got ", arr)
	}
}

func TestToArrayInListAfterRemove(t *testing.T) {
	l := listOfFiveInts()
	l.RemoveFirst()
	l.RemoveLast()
	arr := l.ToArray()
	if len(arr) != 3 {
		t.Error("Array isnt correct Expected [1 2 3] got ", arr)
	}
	if arr[0] != 1 || arr[1] != 2 || arr[2] != 3 {
		t.Error("Array isnt correct Expected [1 2 3] got ", arr)
	}
}

func double(v int) int {
	return v * 2
}

func TestForEachInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	l.ForEach(double)
	//No error
	if !l.IsEmpty() {
		t.Error("List isnt empty")
	}
}

func TestForEachInList(t *testing.T) {
	l := listOfFiveInts()
	l.ForEach(double)
	arr := l.ToArray()
	if len(arr) != 5 {
		t.Error("Array isnt correct Expected [0 2 4 6 8] got ", arr)
	}
	if arr[0] != 0 || arr[1] != 2 || arr[2] != 4 || arr[3] != 6 || arr[4] != 8 {
		t.Error("Array isnt correct Expected [0 2 4 6 8] got ", arr)
	}
}

func TestForEachInListNilFunction(t *testing.T) {
	l := listOfFiveInts()
	l.ForEach(nil)
	arr := l.ToArray()
	if len(arr) != 5 {
		t.Error("Array isnt correct Expected [0 1 2 3 4] got ", arr)
	}
	if arr[0] != 0 || arr[1] != 1 || arr[2] != 2 || arr[3] != 3 || arr[4] != 4 {
		t.Error("Array isnt correct Expected [0 1 2 3 4] got ", arr)
	}
}

func TestListFromArray(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4}
	l := ListFromArray(arr, comp)
	if l.head.value != 0 || l.tail.value != 4 || l.size != 5 {
		t.Error("List isnt correct Expected [0 1 2 3 4] got ", l)
	}
}

func TestListFromArrayEmptyArray(t *testing.T) {
	arr := []int{}
	l := ListFromArray(arr, comp)
	if l.head != nil || l.tail != nil || l.size != 0 {
		t.Error("List isnt correct Expected [] got ", l)
	}
}

func TestListFromArrayNilArray(t *testing.T) {
	l := ListFromArray(nil, comp)
	if l.head != nil || l.tail != nil || l.size != 0 {
		t.Error("List isnt correct Expected [] got ", l)
	}
}

func TestMapInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	m := l.Map(double)
	if !m.IsEmpty() {
		t.Error("List isnt empty")
	}
}

func TestMapInList(t *testing.T) {
	l := listOfFiveInts()
	m := l.Map(double)
	arrList := l.ToArray()
	arrMap := m.ToArray()
	if len(arrList) != 5 {
		t.Error("Array isnt correct Expected [0 1 2 3 4] got ", arrList)
	}
	if arrList[0] != 0 || arrList[1] != 1 || arrList[2] != 2 || arrList[3] != 3 || arrList[4] != 4 {
		t.Error("Array isnt correct Expected [0 1 2 3 4] got ", arrList, " original list shouldnt be modified")
	}
	if len(arrMap) != 5 {
		t.Error("Array isnt correct Expected [0 2 4 6 8] got ", arrMap)
	}
	if arrMap[0] != 0 || arrMap[1] != 2 || arrMap[2] != 4 || arrMap[3] != 6 || arrMap[4] != 8 {
		t.Error("Array isnt correct Expected [0 2 4 6 8] got ", arrMap)
	}
}

func TestMapInListNilFunction(t *testing.T) {
	l := listOfFiveInts()
	m := l.Map(nil)
	if !m.IsEmpty() {
		t.Error("List isnt empty")
	}
}

func TestFilterInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	m := l.Filter(func(v int) bool { return v%2 == 0 })
	if !m.IsEmpty() {
		t.Error("List isnt empty")
	}
}

func TestFilterInList(t *testing.T) {
	l := listOfFiveInts()
	m := l.Filter(func(v int) bool { return v%2 == 0 })
	arrList := l.ToArray()
	arrMap := m.ToArray()
	if len(arrList) != 5 {
		t.Error("Array isnt correct Expected [0 1 2 3 4] got ", arrList)
	}
	if arrList[0] != 0 || arrList[1] != 1 || arrList[2] != 2 || arrList[3] != 3 || arrList[4] != 4 {
		t.Error("Array isnt correct Expected [0 1 2 3 4] got ", arrList, " original list shouldnt be modified")
	}
	if len(arrMap) != 3 {
		t.Error("Array isnt correct Expected [0 2 4] got ", arrMap)
	}
	if arrMap[0] != 0 || arrMap[1] != 2 || arrMap[2] != 4 {
		t.Error("Array isnt correct Expected [0 2 4] got ", arrMap)
	}
}

func TestFilterInListNilFunction(t *testing.T) {
	l := listOfFiveInts()
	m := l.Filter(nil)
	if !m.IsEmpty() {
		t.Error("List isnt empty")
	}
}

func TestSetEmptyList(t *testing.T) {
	l := NewList[int](comp)
	err := l.Set(0, 0)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestSetList(t *testing.T) {
	l := listOfFiveInts()
	err := l.Set(1, 5)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if l.head.next.value != 5 {
		t.Error("Value isnt correct Expected 5 got ", l.head.value)
	}
	if l.tail.value != 4 {
		t.Error("Value isnt correct Expected 4 got ", l.tail.value)
	}
	if l.size != 5 {
		t.Error("Size wasnt updated Expected 5 got ", l.size)
	}
}

func TestSetListOutOfBounds(t *testing.T) {
	l := listOfFiveInts()
	err := l.Set(5, 5)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestSetListOutOfBoundsNegative(t *testing.T) {
	l := listOfFiveInts()
	err := l.Set(-1, 5)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestSetListAtStart(t *testing.T) {
	l := listOfFiveInts()
	err := l.Set(0, 5)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if l.head.value != 5 {
		t.Error("Value isnt correct Expected 5 got ", l.head.value)
	}
	if l.tail.value != 4 {
		t.Error("Value isnt correct Expected 4 got ", l.tail.value)
	}
	if l.size != 5 {
		t.Error("Size wasnt updated Expected 5 got ", l.size)
	}
}

func TestSetListAtEnd(t *testing.T) {
	l := listOfFiveInts()
	err := l.Set(4, 5)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if l.head.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.head.value)
	}
	if l.tail.value != 5 {
		t.Error("Value isnt correct Expected 5 got ", l.tail.value)
	}
	if l.size != 5 {
		t.Error("Size wasnt updated Expected 5 got ", l.size)
	}
}

func TestAddArrayInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	l.AddArray([]int{0, 1, 2, 3, 4})
	if l.head.value != 0 || l.tail.value != 4 || l.size != 5 {
		t.Error("List isnt correct Expected [0 1 2 3 4] got ", l)
	}
}

func TestAddArrayInList(t *testing.T) {
	l := listOfFiveInts()
	l.AddArray([]int{5, 6, 7, 8, 9})
	if l.head.value != 0 || l.tail.value != 9 || l.size != 10 {
		t.Error("List isnt correct Expected [0 1 2 3 4 5 6 7 8 9] got ", l)
	}
	val, err := l.GetAt(7)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 7 {
		t.Error("Value isnt correct Expected 7 got ", val)
	}
}

func TestAddArrayWithEmptyArray(t *testing.T) {
	l := listOfFiveInts()
	l.AddArray([]int{})
	if l.head.value != 0 || l.tail.value != 4 || l.size != 5 {
		t.Error("List isnt correct Expected [0 1 2 3 4] got ", l)
	}
}

func TestAddArrayWithNilArray(t *testing.T) {
	l := listOfFiveInts()
	l.AddArray(nil)
	if l.head.value != 0 || l.tail.value != 4 || l.size != 5 {
		t.Error("List isnt correct Expected [0 1 2 3 4] got ", l)
	}
}

func TestAddArrayAtStartInEmptyList(t *testing.T) {
	l := NewList[int](comp)
	l.AddArrayAt(0, []int{0, 1, 2, 3, 4})
	if l.head.value != 0 || l.tail.value != 4 || l.size != 5 {
		t.Error("List isnt correct Expected [0 1 2 3 4] got ", l)
	}
}

func TestAddArrayAtStartInList(t *testing.T) {
	l := listOfFiveInts()
	l.AddArrayAt(0, []int{5, 6, 7, 8, 9})
	if l.head.value != 5 || l.tail.value != 4 || l.size != 10 {
		arr := l.ToArray()
		t.Error("List isnt correct Expected [5 6 7 8 9 0 1 2 3 4] got ", arr)
	}
	val, err := l.GetAt(7)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 2 {
		t.Error("Value isnt correct Expected 2 got ", val)
	}
}

func TestAddArrayAtWithEmptyArray(t *testing.T) {
	l := listOfFiveInts()
	err := l.AddArrayAt(0, []int{})
	if err == nil {
		t.Error("Error wasnt raised")
	}

}

func TestAddArrayAtWithNilArray(t *testing.T) {
	l := listOfFiveInts()
	err := l.AddArrayAt(0, nil)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestAddArrayAtEndInList(t *testing.T) {
	l := listOfFiveInts()
	l.AddArrayAt(5, []int{5, 6, 7, 8, 9})
	if l.head.value != 0 || l.tail.value != 9 || l.size != 10 {
		t.Error("List isnt correct Expected [0 1 2 3 4 5 6 7 8 9] got ", l)
	}
	val, err := l.GetAt(7)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 7 {
		t.Error("Value isnt correct Expected 7 got ", val)
	}
}

func TestAddArrayAtMiddleInList(t *testing.T) {
	l := listOfFiveInts()
	l.AddArrayAt(2, []int{5, 6, 7, 8, 9})
	arr := l.ToArray()
	expArr := []int{0, 1, 5, 6, 7, 8, 9, 2, 3, 4}

	for i := 0; i < len(expArr); i++ {
		if arr[i] != expArr[i] {
			t.Error("List isnt correct Expected ", expArr, " got ", arr)
			break
		}
	}

	val, err := l.GetAt(2)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 5 {
		t.Error("Value isnt correct Expected 5 got ", val)
	}
	val, err = l.GetAt(4)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 7 {
		t.Error("Value isnt correct Expected 7 got ", val)
	}
	val, err = l.GetAt(6)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 9 {
		t.Error("Value isnt correct Expected 9 got ", val)
	}
}

// Test Queue Interface

func TestEnqueueInEmptyList(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(0)
	if q.Size() != 1 {
		t.Error("Size wasnt updated Expected 1 got ", q.Size())
	}
	val, err := q.Peek()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 0 {
		t.Error("Value isnt correct Expected 0 got ", val)
	}
}

func TestPeekInEmptyList(t *testing.T) {
	q := NewQueue[int]()
	_, err := q.Peek()
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestEnqueueInList(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(0)
	q.Enqueue(1)
	val, err := q.Peek()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 0 {
		t.Error("Value isnt correct Expected 0 got ", val)
	}
}

func TestDequeueInEmptyList(t *testing.T) {
	q := NewQueue[int]()
	_, err := q.Dequeue()
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestDequeueInList(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(0)
	q.Enqueue(1)
	val, err := q.Dequeue()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 0 {
		t.Error("Value isnt correct Expected 0 got ", val)
	}
	val, err = q.Peek()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 1 {
		t.Error("Peek value isnt correct Expected 1 got ", val)
	}
	if q.Size() != 1 {
		t.Error("Size wasnt updated Expected 1 got ", q.Size())
	}
}

func TestDequeueInListToEmpty(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(0)
	q.Enqueue(1)
	val, err := q.Dequeue()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 0 {
		t.Error("Value isnt correct Expected 0 got ", val)
	}
	val, err = q.Dequeue()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 1 {
		t.Error("Value isnt correct Expected 1 got ", val)
	}
	if q.Size() != 0 {
		t.Error("Size wasnt updated Expected 0 got ", q.Size())
	}
}

// Test Stack Interface

func TestPushInEmptyList(t *testing.T) {
	s := NewStack[int]()
	s.Push(0)
	if s.Size() != 1 {
		t.Error("Size wasnt updated Expected 1 got ", s.Size())
	}
	val, err := s.Top()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 0 {
		t.Error("Value isnt correct Expected 0 got ", val)
	}
}

func TestTopInEmptyList(t *testing.T) {
	s := NewStack[int]()
	_, err := s.Top()
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestPushInList(t *testing.T) {
	s := NewStack[int]()
	s.Push(0)
	s.Push(1)
	val, err := s.Top()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 1 {
		t.Error("Value isnt correct Expected 1 got ", val)
	}
}

func TestPopInEmptyList(t *testing.T) {
	s := NewStack[int]()
	_, err := s.Pop()
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestPopInList(t *testing.T) {
	s := NewStack[int]()
	s.Push(0)
	s.Push(1)
	val, err := s.Pop()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 1 {
		t.Error("Value isnt correct Expected 1 got ", val)
	}
	val, err = s.Top()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 0 {
		t.Error("Top value isnt correct Expected 0 got ", val)
	}
	if s.Size() != 1 {
		t.Error("Size wasnt updated Expected 1 got ", s.Size())
	}
}

func TestPopInListToEmpty(t *testing.T) {
	s := NewStack[int]()
	s.Push(0)
	s.Push(1)
	val, err := s.Pop()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 1 {
		t.Error("Value isnt correct Expected 1 got ", val)
	}
	val, err = s.Pop()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if val != 0 {
		t.Error("Value isnt correct Expected 0 got ", val)
	}
	if s.Size() != 0 {
		t.Error("Size wasnt updated Expected 0 got ", s.Size())
	}
}
