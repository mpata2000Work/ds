package list

import (
	"testing"
)

func comp(v1 int, v2 int) int {
	return v1 - v2
}
func TestCreateEmptyListWithoutComparator(t *testing.T) {
	l := NewLinkedList[int](nil)
	if l.head != nil || l.tail != nil || l.size != 0 || l.Comparator != nil {
		t.Error("List wasnt correctly initialize ", l)
	}
}

func TestCreateEmptyListWithComparator(t *testing.T) {
	l := NewLinkedList[int](comp)
	if l.head != nil || l.tail != nil || l.size != 0 || l.Comparator == nil {
		t.Error("List wasnt correctly initialize ", l)
	}
}

func TestCanSetComparator(t *testing.T) {
	l := NewLinkedList[int](nil)
	l.Comparator = comp
	if l.Comparator == nil {
		t.Error("Comparator wasnt set")
	}
}

func TestIsEmptyInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	if !l.IsEmpty() {
		t.Error("List isnt empty")
	}
}

func TestAddFirstInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	l.AddFirst(0)
	if l.head == nil || l.tail == nil {
		t.Error("Head or tail want initialize correctly")
	}
	if l.size != 1 {
		t.Error("Size wasnt updated Expected 1 got ", l.size)
	}
	if l.head.value != 0 || l.tail.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.head.value, l.tail.value)
	}
}

func TestAddLastInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	l.AddLast(0)
	if l.head == nil || l.tail == nil {
		t.Error("Head or tail want initialize correctly")
	}
	if l.size != 1 {
		t.Error("Size wasnt updated Expected 1 got ", l.size)
	}
	if l.head.value != 0 {
		t.Error("Value isnt correct Expected 0 got ", l.head.value)
	}
}

func TestIsEmptyInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	if l.IsEmpty() {
		t.Error("List is empty")
	}
}

func TestSizeInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	if l.Size() != 0 {
		t.Error("Size isnt correct Expected 0 got ", l.Size())
	}
}

func TestSizeInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expected := 3
	if l.Size() != expected {
		t.Error("Size isnt correct Expected 3 got ", l.Size())
	}
}

func TestAddFirstInLinkedList(t *testing.T) {
	l := NewLinkedList[int](comp)
	v1 := 0
	v2 := 1
	l.AddFirst(v1)
	l.AddFirst(v2)
	if l.size != 2 {
		t.Error("Size wasnt updated Expected 2 got ", l.size)
	}
	if l.head.next != l.tail {
		t.Error("Head next isnt tail in list of 2 elements")
	}
	if l.tail.previous != l.head {
		t.Error("Tail previous isnt head in list of 2 elements")
	}
	if l.head.value != v2 || l.tail.value != v1 {
		t.Errorf("Values arent correct Expected %d %d got %d %d", v2, v1, l.head.value, l.tail.value)
	}
}

func TestAddLastInList(t *testing.T) {
	l := NewLinkedList[int](comp)
	v1 := 0
	v2 := 1
	l.AddLast(v1)
	l.AddLast(v2)
	if l.head.next != l.tail {
		t.Error("Head next isnt tail in list of 2 elements")
	}
	if l.tail.previous != l.head {
		t.Error("Tail previous isnt head in list of 2 elements")
	}
	if l.size != 2 {
		t.Error("Size wasnt updated Expected 2 got ", l.size)
	}
	if l.head.value != v1 || l.tail.value != v2 {
		t.Errorf("Values arent correct Expected %d %d got %d %d", v1, v2, l.head.value, l.tail.value)
	}
}

func TestLinkedListToSlice(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	arrOut := l.ToSlice()
	if len(arrOut) != 3 {
		t.Error("Size isnt correct Expected 3 got ", len(arr))
	}
	if arrOut[0] != arr[0] || arrOut[1] != arr[1] || arrOut[2] != arr[2] {
		t.Errorf("Values arent correct Expected %v got %v", arr, arrOut)
	}
}

func TestAddInListEmptyListAtIndexCero(t *testing.T) {
	l := NewLinkedList[int](comp)
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
	l := NewLinkedList[int](comp)
	err := l.Add(1, 0)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestAddInListAtIndexCero(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expectedSize := 4
	index := 0
	value := 5
	arrExpected := []int{5, 0, 1, 2}
	l.Add(index, value)
	arrOut := l.ToSlice()
	if l.head == nil || l.tail == nil {
		t.Error("Head or tail reference was reset")
	}
	if l.size != expectedSize {
		t.Errorf("Size wasnt updated expected %d and got %d", expectedSize, l.size)
	}
	if l.head.value != value {
		t.Errorf("Head value isnt correct expected %d and got %d", value, l.head.value)
	}
	for i := 0; i < len(arrExpected); i++ {
		if arrExpected[i] != arrOut[i] {
			t.Errorf("Array isnt correct expected %v and got %v at pos %v", arrExpected, arrOut, i)
		}
	}

}

func TestAddInListAtLastIndex(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expectedSize := 4
	index := 3
	value := 5
	arrExpected := []int{0, 1, 2, 5}
	err := l.Add(index, value)
	if err != nil {
		t.Error("Error was raised")
	}
	arrOut := l.ToSlice()
	if l.head == nil || l.tail == nil {
		t.Error("Head or tail reference was reset")
	}
	if l.size != expectedSize {
		t.Errorf("Size wasnt updated expected %d and got %d", expectedSize, l.size)
	}
	if l.tail.value != value {
		t.Errorf("Tail value isnt correct expected %d and got %d", value, l.tail.value)
	}
	for i := 0; i < len(arrExpected); i++ {
		if arrExpected[i] != arrOut[i] {
			t.Errorf("Array isnt correct expected %v and got %v at pos %v", arrExpected, arrOut, i)
		}
	}
}

func TestAddInListAtIndex(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expectedSize := 4
	index := 1
	value := 5
	arrExpected := []int{0, 5, 1, 2}
	expcPrevNode := l.head
	expcNextNode := l.head.next
	err := l.Add(index, value)
	if err != nil {
		t.Error("Error was raised")
	}
	arrOut := l.ToSlice()
	if l.head == nil || l.tail == nil {
		t.Error("Head or tail reference was reset")
	}
	if l.size != expectedSize {
		t.Errorf("Size wasnt updated expected %d and got %d", expectedSize, l.size)
	}
	if expcPrevNode.next != expcNextNode.previous {
		t.Error("Previous and next nodes arent linked")
	}
	if l.head.next.value != value {
		t.Errorf("Head next value isnt correct expected %d and got %d", value, l.head.next.value)
	}
	for i := 0; i < len(arrExpected); i++ {
		if arrExpected[i] != arrOut[i] {
			t.Errorf("Array isnt correct expected %v and got %v at pos %v", arrExpected, arrOut, i)
		}
	}
}

func TestGetFirstInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	_, err := l.GetFirst()
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestGetFirstInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	v, err := l.GetFirst()
	if err != nil {
		t.Error("Error was raised")
	}
	if v != 0 {
		t.Error("Value isnt correct expected 0 got ", v)
	}
}

func TestGetLastInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	_, err := l.GetLast()
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestGetLastInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expected := 2
	v, err := l.GetLast()
	if err != nil {
		t.Error("Error was raised")
	}
	if v != expected {
		t.Error("Value isnt correct Expected 4 got ", v)
	}
}

func TestGetAtInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	_, err := l.Get(0)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestGetAtInListFirstHalf(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	l.AddLast(3) // 0 1 2 3
	expected := 1
	index := 1
	v, err := l.Get(index)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != expected {
		t.Errorf("Value isnt correct expected %d got %d", expected, v)
	}
	if v != l.head.next.value {
		t.Errorf("Value isnt correct when traversing foward. Expected %d got %d", expected, v)
	}
}

func TestGetAtInListSecondHalf(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	l.AddLast(3) // 0 1 2 3
	expected := 2
	index := 2
	v, err := l.Get(index)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != expected {
		t.Errorf("Value isnt correct expected %d got %d", expected, v)
	}
	if v != l.tail.previous.value {
		t.Errorf("Value isnt correct when traversing backward. Expected %d got %d", expected, v)
	}
}

func TestGetAtInListLastIndex(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	l.AddLast(3) // 0 1 2 3
	expected := 3
	index := 3
	v, err := l.Get(index)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != expected {
		t.Errorf("Value isnt correct expected %d got %d", expected, v)
	}
	if v != l.tail.value {
		t.Errorf("Value isnt correct when traversing backward. Expected %d got %d", expected, v)
	}
}

func TestGetAtInListOutOfBounds(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	_, err := l.Get(l.size) // Any index >= size is out of bounds
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestGetAtInListOutOfBoundsNegative(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	_, err := l.Get(-1)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestRemoveFirstInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	_, err := l.RemoveFirst()
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestRemoveFirstInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expectedValue := 0
	expectedSize := 2
	expectedArr := []int{1, 2}
	expectedHead := l.head.next
	v, err := l.RemoveFirst()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if v != expectedValue {
		t.Errorf("Value isnt correct, expected %d got %d", expectedValue, v)
	}
	if l.head.value != expectedArr[0] {
		t.Errorf("Head value isnt correct, expected %d got %d", expectedArr[0], l.head.value)
	}
	if l.tail.value != expectedArr[1] {
		t.Errorf("Tail value isnt correct, expected %d got %d", expectedArr[1], l.tail.value)
	}
	if l.size != expectedSize {
		t.Errorf("Size isnt correct, expected %d got %d", expectedSize, l.size)
	}
	if l.head != expectedHead {
		t.Error("Head reference wasnt updated to correct node")
	}
}

func TestRemoveLastInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	_, err := l.RemoveLast()
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestRemoveLastInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expectedValue := 2
	expectedSize := 2
	expectedArr := []int{0, 1}
	expectedTail := l.tail.previous
	v, err := l.RemoveLast()
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if v != expectedValue {
		t.Errorf("Value isnt correct, expected %d got %d", expectedValue, v)
	}
	if l.head.value != expectedArr[0] {
		t.Errorf("Head value isnt correct, expected %d got %d", expectedArr[0], l.head.value)
	}
	if l.tail.value != expectedArr[1] {
		t.Errorf("Tail value isnt correct, expected %d got %d", expectedArr[1], l.tail.value)
	}
	if l.size != expectedSize {
		t.Errorf("Size isnt correct, expected %d got %d", expectedSize, l.size)
	}
	if l.tail != expectedTail {
		t.Error("Tail reference wasnt updated to correct node")
	}
}

func TestRemoveInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	_, err := l.Remove(0)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestRemoveInList(t *testing.T) {
	index := 1
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expectedValue := 1
	expectedSize := 2
	expectedArr := []int{0, 2}
	expectedPrevNode := l.head
	expectedNextNode := l.tail
	v, err := l.Remove(index)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if v != expectedValue {
		t.Errorf("Value isnt correct, expected %d got %d", expectedValue, v)
	}
	if l.head.value != expectedArr[0] {
		t.Errorf("Head value isnt correct, expected %d got %d", expectedArr[0], l.head.value)
	}
	if l.tail.value != expectedArr[1] {
		t.Errorf("Tail value isnt correct, expected %d got %d", expectedArr[1], l.tail.value)
	}
	if l.size != expectedSize {
		t.Errorf("Size isnt correct, expected %d got %d", expectedSize, l.size)
	}
	if expectedPrevNode != expectedNextNode.previous || expectedNextNode != expectedPrevNode.next {
		t.Error("Previous and next nodes arent linked")
	}
}

func TestRemoveInListOutOfBounds(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	_, err := l.Remove(l.size)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestRemoveInListOutOfBoundsNegative(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	_, err := l.Remove(-1)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestRemoveInListAtStart(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expectedValue := l.head.value
	expectedSize := 2
	expectedArr := []int{1, 2}
	expectedHead := l.head.next
	v, err := l.Remove(0)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if v != expectedValue {
		t.Errorf("Value isnt correct, expected %d got %d", expectedValue, v)
	}
	if l.head.value != expectedArr[0] {
		t.Errorf("Head value isnt correct, expected %d got %d", expectedArr[0], l.head.value)
	}
	if l.tail.value != expectedArr[1] {
		t.Errorf("Tail value isnt correct, expected %d got %d", expectedArr[1], l.tail.value)
	}
	if l.size != expectedSize {
		t.Errorf("Size isnt correct, expected %d got %d", expectedSize, l.size)
	}
	if l.head != expectedHead {
		t.Error("Head reference wasnt updated to correct node")
	}
}

func TestRemoveInListAtEnd(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expectedValue := l.tail.value
	expectedSize := 2
	expectedArr := []int{0, 1}
	expectedTail := l.tail.previous
	v, err := l.Remove(l.size - 1)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if v != expectedValue {
		t.Errorf("Value isnt correct, expected %d got %d", expectedValue, v)
	}
	if l.head.value != expectedArr[0] {
		t.Errorf("Head value isnt correct, expected %d got %d", expectedArr[0], l.head.value)
	}
	if l.tail.value != expectedArr[1] {
		t.Errorf("Tail value isnt correct, expected %d got %d", expectedArr[1], l.tail.value)
	}
	if l.size != expectedSize {
		t.Errorf("Size isnt correct, expected %d got %d", expectedSize, l.size)
	}
	if l.tail != expectedTail {
		t.Error("Tail reference wasnt updated to correct node")
	}
}

func TestRemoveFirstInListToEmpty(t *testing.T) {
	l := LinkedListFromArray[int](nil, 0, 1)
	expectedV1 := 0
	expectedV2 := 1
	expectedSize := 0
	v1, _ := l.RemoveFirst()
	v2, _ := l.RemoveFirst()
	if v1 != expectedV1 || v2 != expectedV2 {
		t.Errorf("Values arent correct Expected %d %d got %d %d", expectedV1, expectedV2, v1, v2)
	}
	if l.head != nil || l.tail != nil {
		t.Error("Head or tail reference wasnt set to nil")
	}
	if l.size != expectedSize {
		t.Errorf("Size isnt correct Expected %d got %d", expectedSize, l.size)
	}
	if l.IsEmpty() != true {
		t.Error("List isnt empty")
	}
}

func TestRemoveLastInListToEmpty(t *testing.T) {
	l := LinkedListFromArray[int](nil, 0, 1)
	expectedV1 := 1
	expectedV2 := 0
	expectedSize := 0
	v1, _ := l.RemoveLast()
	v2, _ := l.RemoveLast()
	if v1 != expectedV1 || v2 != expectedV2 {
		t.Errorf("Values arent correct Expected %d %d got %d %d", expectedV1, expectedV2, v1, v2)
	}
	if l.head != nil || l.tail != nil {
		t.Error("Head or tail reference wasnt set to nil")
	}
	if l.size != expectedSize {
		t.Errorf("Size isnt correct Expected %d got %d", expectedSize, l.size)
	}
	if l.IsEmpty() != true {
		t.Error("List isnt empty")
	}
}

func TestRemoveInListToEmpty(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expectedV1 := 1
	expectedV2 := 2
	expectedV3 := 0
	expectedSize := 0
	v1, _ := l.Remove(1)
	v2, _ := l.Remove(1)
	v3, _ := l.Remove(0)
	if v1 != expectedV1 || v2 != expectedV2 || v3 != expectedV3 {
		t.Errorf("Values arent correct Expected %d %d %d got %d %d %d", expectedV1, expectedV2, expectedV3, v1, v2, v3)
	}
	if l.head != nil || l.tail != nil {
		t.Error("Head or tail reference wasnt set to nil")
	}
	if l.size != expectedSize {
		t.Errorf("Size isnt correct Expected %d got %d", expectedSize, l.size)
	}
	if l.IsEmpty() != true {
		t.Error("List isnt empty")
	}

}

func TestClearEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	l.Clear()
	if l.head != nil || l.tail != nil || l.size != 0 {
		t.Error("List wasnt cleared correctly")
	}
}

func TestClearList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	l.Clear()
	if l.head != nil || l.tail != nil || l.size != 0 {
		t.Error("List wasnt cleared correctly")
	}
}

func TestContainsInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	v, err := l.Contains(0)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != false {
		t.Error("Value isnt correct Expected false got ", v)
	}
}

func TestContainsInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	v, err := l.Contains(2)
	if err != nil {
		t.Error("Error was raised")
	}
	if !v {
		t.Error("Value wasn't found when in list")
	}
}

func TestContainsInListNotInList(t *testing.T) {
	l := LinkedListFromArray[int](comp, 0, 1)
	v, err := l.Contains(5)
	if err != nil {
		t.Error("Error was raised")
	}
	if v {
		t.Error("Value was found when not in list")
	}
}

func TestContainsWithoutComparator(t *testing.T) {
	l := NewLinkedList[int](nil)
	l.AddLast(0)
	_, err := l.Contains(0)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestIndexOfInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	v, err := l.IndexOf(0)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != -1 {
		t.Error("Value isnt correct Expected -1 got ", v)
	}
}

func TestIndexOfInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expected := 2
	v, err := l.IndexOf(2)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != expected {
		t.Error("Value isnt correct Expected 2 got ", v)
	}
}

func TestIndexOfInListNotInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	v, err := l.IndexOf(5)
	if err != nil {
		t.Error("Error was raised")
	}
	if v != -1 {
		t.Error("Value isnt correct Expected -1 got ", v)
	}
}

func TestIndexOfWithoutComparator(t *testing.T) {
	l := NewLinkedList[int](nil)
	l.AddLast(0)
	_, err := l.IndexOf(0)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestToArrayInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	arr := l.ToSlice()
	if len(arr) != 0 {
		t.Error("Array isnt correct Expected [] got ", arr)
	}
}

func double(v int) int {
	return v * 2
}

func TestForEachInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	l.ForEach(double)
	//No error
	if !l.IsEmpty() {
		t.Error("List isnt empty")
	}
}

func TestForEachInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expected := []int{0, 2, 4}
	l.ForEach(double)
	arrOut := l.ToSlice()
	if len(arrOut) != 3 {
		t.Error("Array isnt correct Expected [0 2 4] got ", arrOut)
	}
	for i := 0; i < len(arrOut); i++ {
		if arrOut[i] != expected[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", expected, arrOut, i)
		}
	}
}

func TestForEachInListNilFunction(t *testing.T) {
	expected := []int{0, 1, 2}
	l := LinkedListFromArray(comp, expected...)
	l.ForEach(nil)
	arr := l.ToSlice()
	if len(arr) != 3 {
		t.Error("Array isnt correct Expected [0 1 2] got ", arr)
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", expected, arr, i)
		}
	}
}

func TestListFromArray(t *testing.T) {
	expectedArr := []int{0, 1, 2, 3, 4}
	l := LinkedListFromArray(comp, expectedArr...)
	arr := l.ToSlice()
	if l.head == nil || l.tail == nil {
		t.Error("Head or tail reference wasnt set")
	}
	if l.head.value != expectedArr[0] || l.tail.value != expectedArr[len(expectedArr)-1] {
		t.Error("Head or tail value isnt correct")
	}
	if len(arr) != len(expectedArr) {
		t.Error("Array isnt correct Expected [0 1 2 3 4] got ", arr)
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != expectedArr[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", expectedArr, arr, i)
		}
	}
}

func TestListFromArrayEmptyArray(t *testing.T) {
	arr := []int{}
	l := LinkedListFromArray(comp, arr...)
	if l.head != nil || l.tail != nil || l.size != 0 {
		t.Error("List isnt correct Expected [] got ", l)
	}
}

func TestMapInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	m := l.Map(double)
	if !m.IsEmpty() {
		t.Error("List isnt empty")
	}
}

func TestMapInList(t *testing.T) {
	expectedList := []int{0, 1, 2}
	l := LinkedListFromArray(comp, expectedList...)
	expectedMap := []int{0, 2, 4}
	m := l.Map(double)
	arrList := l.ToSlice()
	arrMap := m.ToSlice()
	if len(arrList) != 3 {
		t.Error("Array isnt correct Expected [0 1 2] got ", arrList)
	}
	for i := 0; i < len(arrList); i++ {
		if arrList[i] != expectedList[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", expectedList, arrList, i)
		}
		if arrMap[i] != expectedMap[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", expectedMap, arrMap, i)
		}
	}
}

func TestMapInListNilFunction(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	m := l.Map(nil)
	if !m.IsEmpty() {
		t.Error("List isnt empty")
	}
}

func TestFilterInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	m := l.Filter(func(v int) bool { return v%2 == 0 })
	if !m.IsEmpty() {
		t.Error("List isnt empty")
	}
}

func TestFilterInList(t *testing.T) {
	expectedList := []int{0, 1, 2}
	l := LinkedListFromArray(comp, expectedList...)
	expectedFilt := []int{0, 2}
	m := l.Filter(func(v int) bool { return v%2 == 0 })
	arrList := l.ToSlice()
	arrFilt := m.ToSlice()
	if len(arrList) != 3 {
		t.Error("Array isnt correct Expected [0 1 2] got ", arrList)
	}
	for i := 0; i < len(arrList); i++ {
		if arrList[i] != expectedList[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", expectedList, arrList, i)
		}
	}
	if len(arrFilt) != 2 {
		t.Error("Array isnt correct Expected [0 2] got ", arrFilt)
	}
	for i := 0; i < len(arrFilt); i++ {
		if arrFilt[i] != expectedFilt[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", expectedFilt, arrFilt, i)
		}
	}
}

func TestFilterInListNilFunction(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	m := l.Filter(nil)
	if !m.IsEmpty() {
		t.Error("List isnt empty")
	}
}

func TestSetEmptyListReturnError(t *testing.T) {
	l := NewLinkedList[int](comp)
	index := 0
	value := 0
	err := l.Set(index, value)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestSetList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	index := 1
	value := 5
	err := l.Set(index, value)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if l.head.value != 0 || l.tail.value != 2 || l.size != 3 {
		t.Errorf("Other values were changed. Expected for head 0 got %d, for tail 2 got %d and for size 3 got %d", l.head.value, l.tail.value, l.size)
	}
	if l.head.next.value != value {
		t.Errorf("Value isnt correct Expected %d got %d", value, l.head.next.value)
	}
}

func TestSetListOutOfBounds(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	err := l.Set(l.size, 5)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestSetListOutOfBoundsNegative(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	err := l.Set(-1, 5)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestSetListAtStart(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expected := 5
	err := l.Set(0, expected)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if l.head.value != expected {
		t.Errorf("Value isnt correct Expected %d got %d", expected, l.head.value)
	}
}

func TestSetListAtEnd(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	expected := 5
	err := l.Set(l.size-1, expected)
	if err != nil {
		t.Error("Error was raised ", err)
	}
	if l.tail.value != expected {
		t.Errorf("Value isnt correct Expected %d got %d", expected, l.tail.value)
	}
}

func TestAddAllInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	arr := []int{0, 1, 2}
	expectedSize := 3
	l.AddAll(arr...)
	arrOut := l.ToSlice()
	if l.head.value != arr[0] || l.tail.value != arr[len(arr)-1] || l.size != expectedSize {
		t.Error("List isnt correct Expected [0 1 2] got ", l)
	}
	for i := 0; i < len(arrOut); i++ {
		if arrOut[i] != arr[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", arr, arrOut, i)
		}
	}
}

func TestAddAllInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	arrIn := []int{3, 4}
	expectedArr := []int{0, 1, 2, 3, 4}
	expectedSize := 5
	tailNode := l.tail
	l.AddAll(arrIn...)
	arrOut := l.ToSlice()
	if l.head.value != expectedArr[0] || l.tail.value != expectedArr[len(expectedArr)-1] {
		t.Error("List ends arent correct")
	}
	if l.size != expectedSize {
		t.Errorf("Size isnt correct Expected %d got %d", expectedSize, l.size)
	}
	if l.tail.previous.previous != tailNode || tailNode.next != l.tail.previous || tailNode.next.next != l.tail {
		t.Error("Tail previous reference wasnt updated")
	}
	if tailNode.next.value != arrIn[0] || tailNode.next.next.value != arrIn[1] {
		t.Error("Tail next references arent correct")
	}
	for i := 0; i < len(arrOut); i++ {
		if arrOut[i] != expectedArr[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", expectedArr, arrOut, i)
		}
	}
}

func TestAddAllWithEmptyArray(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	l.AddAll([]int{}...)
	arrOut := l.ToSlice()
	if len(arrOut) != len(arr) {
		t.Error("Array isnt correct Expected [0 1 2] got ", arrOut)
	}
	for i := 0; i < len(arrOut); i++ {
		if arrOut[i] != arr[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", arr, arrOut, i)
		}
	}
}

func TestAddAllAtStartInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	arrIn := []int{0, 1, 2}
	expectedSize := 3
	expectedArr := []int{0, 1, 2}
	l.AddAllAt(0, arrIn...)
	arrOut := l.ToSlice()
	if l.head.value != expectedArr[0] || l.tail.value != expectedArr[len(expectedArr)-1] {
		t.Error("List ends arent correct")
	}
	if l.size != expectedSize {
		t.Errorf("Size isnt correct Expected %d got %d", expectedSize, l.size)
	}
	for i := 0; i < len(arrOut); i++ {
		if arrOut[i] != expectedArr[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", expectedArr, arrOut, i)
		}
	}
}

func TestAddAllAtStartInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	arrIn := []int{3, 4, 5}
	expectedArr := []int{3, 4, 5, 0, 1, 2}
	expectedSize := 6
	prevHead := l.head

	l.AddAllAt(0, arrIn...)
	arrOut := l.ToSlice()
	if l.head.value != expectedArr[0] || l.tail.value != expectedArr[len(expectedArr)-1] {
		t.Error("List ends arent correct")
	}
	if l.size != expectedSize {
		t.Errorf("Size isnt correct Expected %d got %d", expectedSize, l.size)
	}
	if l.head.next.next.next != prevHead {
		t.Error("Head next references arent correct")
	}
	if prevHead.previous.previous.previous != l.head {
		t.Error("Head previous references arent correct")
	}
	for i := 0; i < len(arrOut); i++ {
		if arrOut[i] != expectedArr[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", expectedArr, arrOut, i)
		}
	}
}

func TestAddAllAtWithEmptyArray(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	err := l.AddAllAt(0, []int{}...)
	if err == nil {
		t.Error("Error wasnt raised")
	}

}

func TestAddAllAtWithNilArray(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	err := l.AddAllAt(0, nil...)
	if err == nil {
		t.Error("Error wasnt raised")
	}
}

func TestAddAllAtEndInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	arrIn := []int{3, 4}
	expectedArr := []int{0, 1, 2, 3, 4}
	expectedSize := 5
	prevTail := l.tail

	l.AddAllAt(l.size, arrIn...)
	arrOut := l.ToSlice()
	if l.head.value != expectedArr[0] || l.tail.value != expectedArr[len(expectedArr)-1] {
		t.Error("List ends arent correct")
	}
	if l.size != expectedSize {
		t.Errorf("Size isnt correct Expected %d got %d", expectedSize, l.size)
	}
	if l.tail.previous.previous != prevTail {
		t.Error("Tail previous references arent correct")
	}
	if prevTail.next.next != l.tail {
		t.Error("Tail next references arent correct")
	}
	for i := 0; i < len(arrOut); i++ {
		if arrOut[i] != expectedArr[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", expectedArr, arrOut, i)
		}
	}
}

func TestAddAllAtMiddleInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	arrIn := []int{3, 4}
	index := 1
	expectedArr := []int{0, 3, 4, 1, 2}
	expectedSize := 5
	prevNode := l.head
	nextNode := l.head.next

	l.AddAllAt(index, arrIn...)
	arrOut := l.ToSlice()
	if l.head.value != expectedArr[0] || l.tail.value != expectedArr[len(expectedArr)-1] {
		t.Error("List ends arent correct")
	}
	if l.size != expectedSize {
		t.Errorf("Size isnt correct Expected %d got %d", expectedSize, l.size)
	}
	if nextNode.previous.previous.previous != prevNode {
		t.Error("Node previous references arent correct")
	}
	if prevNode.next.next.next != nextNode {
		t.Error("Node next references arent correct")
	}
	for i := 0; i < len(arrOut); i++ {
		if arrOut[i] != expectedArr[i] {
			t.Errorf("Array isnt correct Expected %v got %v at pos %v", expectedArr, arrOut, i)
		}
	}
}

func TestAddAllAtOutOfBounds(t *testing.T) {
	l := LinkedListFromArray(comp, []int{0, 1, 2}...)
	err1 := l.AddAllAt(l.size+1, 3, 4)
	err2 := l.AddAllAt(-1, 3, 4)
	if err1 == nil || err2 == nil {
		t.Error("Error wasnt raised")
	}
}

func TestRemoveElementInEmptyList(t *testing.T) {
	l := NewLinkedList[int](comp)
	v := l.RemoveElement(0)
	if v {
		t.Error("Value was found when not in list")
	}
}

func TestRemoveElementInList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	v := l.RemoveElement(1)
	if !v {
		t.Error("Value wasnt removed when in list")
	}
	if l.head.value != 0 || l.tail.value != 2 || l.size != 2 {
		t.Errorf("Other values were changed. Expected for head 0 got %d, for tail 2 got %d and for size 2 got %d", l.head.value, l.tail.value, l.size)
	}
	if l.head.next.value != 2 {
		t.Errorf("Value isnt correct Expected 2 got %d", l.head.next.value)
	}
}

func TestRemoveElementNotInLinkedList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := LinkedListFromArray(comp, arr...)
	v := l.RemoveElement(5)
	if v {
		t.Error("Value was found when not in list")
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
