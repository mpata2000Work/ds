package list

import (
	"testing"
)

func TestCreateEmptyArrayList(t *testing.T) {
	list := ArrayList[int]{}
	if list.Size() != 0 {
		t.Error("Expected size to be 0")
	}
	if !list.IsEmpty() {
		t.Error("Expected list to be empty")
	}
}

func TestCreateEmptyArrayListWithComparator(t *testing.T) {
	l := NewArrayList[int](func(a, b int) int {
		return a - b
	})
	if l.Size() != 0 {
		t.Error("Expected size to be 0")
	}
	if !l.IsEmpty() {
		t.Error("Expected list to be empty")
	}
}

func TestArrayListCanSetComparator(t *testing.T) {
	l := NewArrayList[int](nil)
	l.Comparator = func(a, b int) int {
		return b - a
	}
	if l.Size() != 0 {
		t.Error("Expected size to be 0")
	}
	if !l.IsEmpty() {
		t.Error("Expected list to be empty")
	}
}

func TestAddFirstInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	l.AddFirst(1)
	if l.Size() != 1 {
		t.Error("Expected size to be 1")
	}
	if l.IsEmpty() {
		t.Error("Expected list not to be empty")
	}
	if l.arr[0] != 1 {
		t.Error("Expected first element to be 1")
	}
}

func TestAddLastInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	l.AddLast(1)
	if l.Size() != 1 {
		t.Error("Expected size to be 1")
	}
	if l.IsEmpty() {
		t.Error("Expected list not to be empty")
	}
	if l.arr[0] != 1 {
		t.Error("Expected first element to be 1")
	}
}

func TestAddFirstInLinkedArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	value := 1
	l.AddFirst(value)
	if l.Size() != 1 {
		t.Error("Expected size to be 1")
	}
	if l.IsEmpty() {
		t.Error("Expected list not to be empty")
	}
	if l.arr[0] != value {
		t.Error("Expected first element to be 1")
	}
}

func TestAddLastInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	value := 3
	expectedArr := []int{0, 1, 2, 3}
	expectedSize := 4
	l.AddLast(value)
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}
}

func TestArrayListToSlice(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	slice := l.ToSlice()
	if len(slice) != len(arr) {
		t.Errorf("Expected slice length to be %d", len(arr))
	}
	for i := 0; i < len(arr); i++ {
		if slice[i] != arr[i] {
			t.Errorf("Expected element at index %d to be %d", i, arr[i])
		}
	}
}

func TestAddInArrayListEmptyArrayListAtIndexCero(t *testing.T) {
	l := NewArrayList[int](nil)
	value := 1
	err := l.Add(0, value)
	if err != nil {
		t.Error("Expected no error")
	}
	if l.Size() != 1 {
		t.Error("Expected size to be 1")
	}
	if l.arr[0] != value {
		t.Error("Expected first element to be 1")
	}
}

func TestAddInArrayListEmptyArrayListAtIndexOut(t *testing.T) {
	l := NewArrayList[int](nil)
	value := 1
	err := l.Add(1, value)
	if err == nil {
		t.Error("Expected error")
	}
	if l.Size() != 0 {
		t.Error("Expected size to be 0")
	}
}

func TestAddInArrayListAtIndexCero(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	value := 3
	expectedArr := []int{3, 0, 1, 2}
	expectedSize := 4
	err := l.Add(0, value)
	if err != nil {
		t.Error("Expected no error")
	}
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}
}

func TestAddInArrayListAtSize(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	value := 3
	expectedArr := []int{0, 1, 2, 3}
	expectedSize := 4
	err := l.Add(l.Size(), value)
	if err != nil {
		t.Error("Expected no error")
	}
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}
}

func TestAddInArrayListAtIndex(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	value := 3
	expectedArr := []int{0, 1, 3, 2}
	expectedSize := 4
	index := 2
	err := l.Add(index, value)
	if err != nil {
		t.Error("Expected no error")
	}
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}
}

func TestGetFirstInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	_, err := l.GetFirst()
	if err == nil {
		t.Error("Expected error")
	}
}

func TestGetFirstInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	value, err := l.GetFirst()
	if err != nil {
		t.Error("Expected no error")
	}
	if value != arr[0] {
		t.Errorf("Expected first element to be %d", arr[0])
	}
}

func TestGetLastInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	_, err := l.GetLast()
	if err == nil {
		t.Error("Expected error")
	}
}

func TestGetLastInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	value, err := l.GetLast()
	if err != nil {
		t.Error("Expected no error")
	}
	if value != arr[len(arr)-1] {
		t.Errorf("Expected last element to be %d", arr[len(arr)-1])
	}
}

func TestGetAtInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	_, err := l.Get(0)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestGetAtInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	value, err := l.Get(1)
	if err != nil {
		t.Error("Expected no error")
	}
	if value != arr[1] {
		t.Errorf("Expected element at index 1 to be %d", arr[1])
	}
}

func TestGetAtInArrayListOutOfBoundsNegative(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	_, err := l.Get(-1)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestRemoveFirstInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	_, err := l.RemoveFirst()
	if err == nil {
		t.Error("Expected error")
	}
}

func TestRemoveFirstInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	expectedArr := []int{1, 2}
	expectedSize := 2
	expectedValue := arr[0]
	value, err := l.RemoveFirst()
	if err != nil {
		t.Error("Expected no error")
	}
	if value != expectedValue {
		t.Errorf("Expected removed element to be %d", expectedValue)
	}
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}
}

func TestRemoveLastInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	_, err := l.RemoveLast()
	if err == nil {
		t.Error("Expected error")
	}
}

func TestRemoveLastInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	expectedArr := []int{0, 1}
	expectedSize := 2
	expectedValue := arr[len(arr)-1]
	value, err := l.RemoveLast()
	if err != nil {
		t.Error("Expected no error")
	}
	if value != expectedValue {
		t.Errorf("Expected removed element to be %d", expectedValue)
	}
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}
}

func TestRemoveInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	_, err := l.Remove(0)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestRemoveInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	expectedArr := []int{0, 2}
	expectedSize := 2
	expectedValue := arr[1]
	value, err := l.Remove(1)
	if err != nil {
		t.Error("Expected no error")
	}
	if value != expectedValue {
		t.Errorf("Expected removed element to be %d", expectedValue)
	}
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}
}

func TestRemoveInArrayListOutOfBounds(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	_, err := l.Remove(3)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestRemoveInArrayListOutOfBoundsNegative(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	_, err := l.Remove(-1)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestRemoveInArrayListAtStart(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	expectedArr := []int{1, 2}
	expectedSize := 2
	expectedValue := arr[0]
	value, err := l.Remove(0)
	if err != nil {
		t.Error("Expected no error")
	}
	if value != expectedValue {
		t.Errorf("Expected removed element to be %d", expectedValue)
	}
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}

}

func TestRemoveInArrayListAtEnd(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	expectedArr := []int{0, 1}
	expectedSize := 2
	expectedValue := arr[len(arr)-1]
	value, err := l.Remove(2)
	if err != nil {
		t.Error("Expected no error")
	}
	if value != expectedValue {
		t.Errorf("Expected removed element to be %d", expectedValue)
	}
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}
}

func TestRemoveFirstInArrayListToEmpty(t *testing.T) {
	l := NewArrayList[int](nil)
	_, err := l.RemoveFirst()
	if err == nil {
		t.Error("Expected error")
	}
}

func TestRemoveLastInArrayListToEmpty(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	v1, err1 := l.RemoveLast()
	v2, err2 := l.RemoveLast()
	v3, err3 := l.RemoveLast()
	if err1 != nil || err2 != nil || err3 != nil {
		t.Error("Expected no error")
	}
	if v1 != 2 || v2 != 1 || v3 != 0 {
		t.Error("Expected values to be 2, 1 and 0")
	}
}

func TestRemoveInArrayListToEmpty(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	v1, err1 := l.Remove(1)
	v2, err2 := l.Remove(0)
	v3, err3 := l.Remove(0)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Error("Expected no error")
	}
	if v1 != 1 || v2 != 0 || v3 != 2 {
		t.Error("Expected values to be 1, 0 and 2")
	}
}

func TestClearEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	l.Clear()
	if l.Size() != 0 {
		t.Error("Expected size to be 0")
	}
	if !l.IsEmpty() {
		t.Error("Expected list to be empty")
	}
}

func TestClearArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	l.Clear()
	if l.Size() != 0 {
		t.Error("Expected size to be 0")
	}
	if !l.IsEmpty() {
		t.Error("Expected list to be empty")
	}
}

func compArrayListTest(a, b int) int {
	return a - b
}

func TestContainsInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](compArrayListTest)
	v, err := l.Contains(0)
	if err != nil {
		t.Error("Expected no error")
	}
	if v {
		t.Error("Expected value to be false")
	}
}

func TestContainsInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(compArrayListTest, arr...)
	v, err := l.Contains(1)
	if err != nil {
		t.Error("Expected no error")
	}
	if !v {
		t.Error("Expected value to be true")
	}
}

func TestContainsInArrayListNotInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(compArrayListTest, arr...)
	v, err := l.Contains(3)
	if err != nil {
		t.Error("Expected no error")
	}
	if v {
		t.Error("Expected value to be false")
	}
}

func TestArrayListContainsWithoutComparator(t *testing.T) {
	l := NewArrayList[int](nil)
	_, err := l.Contains(0)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestIndexOfInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](compArrayListTest)
	_, err := l.IndexOf(0)
	if err != nil {
		t.Error("Expected no error")
	}
}

func TestIndexOfInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(compArrayListTest, arr...)
	index, err := l.IndexOf(1)
	if err != nil {
		t.Error("Expected no error")
	}
	if index != 1 {
		t.Errorf("Expected index to be 1")
	}
}

func TestIndexOfInArrayListNotInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(compArrayListTest, arr...)
	index, err := l.IndexOf(3)
	if err != nil {
		t.Error("Expected no error")
	}
	if index != -1 {
		t.Errorf("Expected index to be -1")
	}
}

func TestArrayListIndexOfWithoutComparator(t *testing.T) {
	l := NewArrayList[int](nil)
	_, err := l.IndexOf(0)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestToSliceInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	arr := l.ToSlice()
	if len(arr) != 0 {
		t.Error("Expected array length to be 0")
	}
}

func TestForEachInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	l.ForEach(func(v int) int {
		return v + 1
	})
	if l.Size() != 0 {
		t.Error("Expected size to be 0")
	}
}

func TestForEachInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	l.ForEach(func(v int) int {
		return v + 1
	})
	if l.Size() != len(arr) {
		t.Errorf("Expected size to be %d", len(arr))
	}
	for i := 0; i < len(arr); i++ {
		if l.arr[i] != arr[i]+1 {
			t.Errorf("Expected element at index %d to be %d", i, arr[i]+1)
		}
	}
}

func TestForEachInArrayListNilFunction(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	l.ForEach(nil)
	if l.Size() != len(arr) {
		t.Errorf("Expected size to be %d", len(arr))
	}
	for i := 0; i < len(arr); i++ {
		if l.arr[i] != arr[i] {
			t.Errorf("Expected element at index %d to be %d", i, arr[i])
		}
	}
}

func TestArrayListFromArrayEmptyArray(t *testing.T) {
	l := ArrayListFromArray(compArrayListTest)
	if l.Size() != 0 {
		t.Error("Expected size to be 0")
	}
}

func TestMapInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	out := l.Map(func(v int) int {
		return v + 1
	})
	if out.Size() != 0 {
		t.Error("Expected size to be 0")
	}
}

func TestMapInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	out := l.Map(func(v int) int {
		return v + 1
	})
	if out.Size() != len(arr) {
		t.Errorf("Expected size to be %d", len(arr))
	}
	outArr := out.ToSlice()
	for i := 0; i < len(arr); i++ {
		if outArr[i] != arr[i]+1 {
			t.Errorf("Expected element at index %d to be %d", i, arr[i]+1)
		}
	}
}

func TestMapInArrayListNilFunction(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	out := l.Map(nil)
	if out.Size() != 0 {
		t.Error("Expected size to be 0")
	}
}

func TestFilterInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	out := l.Filter(func(v int) bool {
		return v%2 == 0
	})
	if out.Size() != 0 {
		t.Error("Expected size to be 0")
	}
}

func TestFilterInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	out := l.Filter(func(v int) bool {
		return v%2 == 0
	})
	if out.Size() != 2 {
		t.Error("Expected size to be 2")
	}
	outArr := out.ToSlice()
	for i := 0; i < len(outArr); i++ {
		if outArr[i] != arr[i*2] {
			t.Errorf("Expected element at index %d to be %d", i, arr[i*2])
		}
	}
}

func TestFilterInArrayListNilFunction(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	out := l.Filter(nil)
	if out.Size() != 0 {
		t.Error("Expected size to be 0")
	}
}

func TestSetEmptyArrayListReturnError(t *testing.T) {
	l := NewArrayList[int](nil)
	err := l.Set(0, 0)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestSetArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	value := 3
	expectedArr := []int{3, 1, 2}
	expectedSize := 3
	index := 0
	err := l.Set(index, value)
	if err != nil {
		t.Error("Expected no error")
	}
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}
}

func TestSetArrayListOutOfBounds(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	value := 3
	index := 3
	err := l.Set(index, value)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestSetArrayListOutOfBoundsNegative(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	value := 3
	index := -1
	err := l.Set(index, value)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestAddAllInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	arr2 := []int{3, 4, 5}
	expectedArr := []int{0, 1, 2, 3, 4, 5}
	expectedSize := 6
	l.AddAll(arr2...)
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}
}

func TestArrayListAddArrayWithEmptyArray(t *testing.T) {
	l := NewArrayList[int](nil)
	l.AddAll()
	if l.Size() != 0 {
		t.Error("Expected size to be 0")
	}
}

func TestAddAllAtStartInEmptyArrayList(t *testing.T) {
	l := NewArrayList[int](nil)
	arr := []int{0, 1, 2}
	expectedArr := []int{0, 1, 2}
	expectedSize := 3
	l.AddAllAt(0, arr...)
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}
}

func TestAddAllAtStartInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	arr2 := []int{3, 4, 5}
	expectedArr := []int{3, 4, 5, 0, 1, 2}
	expectedSize := 6
	l.AddAllAt(0, arr2...)
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}
}

func TestArrayListAddAllAtOutOfBounds(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	arr2 := []int{3, 4, 5}
	index := 4
	err1 := l.AddAllAt(index, arr2...)
	err2 := l.AddAllAt(-1, arr2...)
	if err1 == nil || err2 == nil {
		t.Error("Expected error")
	}
}

func TestArrayListAddAllAtWithNilArray(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	err := l.AddAllAt(0)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestAddAllAtEndInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	arr2 := []int{3, 4, 5}
	expectedArr := []int{0, 1, 2, 3, 4, 5}
	expectedSize := 6
	l.AddAllAt(l.Size(), arr2...)
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}
}

func TestAddAllAtMiddleInArrayList(t *testing.T) {
	arr := []int{0, 1, 2}
	l := ArrayListFromArray(nil, arr...)
	arr2 := []int{3, 4, 5}
	expectedArr := []int{0, 1, 3, 4, 5, 2}
	expectedSize := 6
	l.AddAllAt(2, arr2...)
	if l.Size() != expectedSize {
		t.Errorf("Expected size to be %d", expectedSize)
	}
	for i := 0; i < expectedSize; i++ {
		if l.arr[i] != expectedArr[i] {
			t.Errorf("Expected element at index %d to be %d", i, expectedArr[i])
		}
	}

}
