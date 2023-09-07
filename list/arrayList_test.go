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
}

func TestIsEmptyInEmptyArrayList(t *testing.T) {
}

func TestAddFirstInEmptyArrayList(t *testing.T) {
}

func TestAddLastInEmptyArrayList(t *testing.T) {
}

func TestIsEmptyInArrayList(t *testing.T) {
}

func TestSizeInEmptyArrayList(t *testing.T) {
}

func TestSizeInArrayList(t *testing.T) {
}

func TestAddFirstInLinkedArrayList(t *testing.T) {
}

func TestAddLastInArrayList(t *testing.T) {
}

func TestLinkedArrayListToSlice(t *testing.T) {
}

func TestAddInArrayListEmptyArrayListAtIndexCero(t *testing.T) {
}

func TestAddInArrayListEmptyArrayListAtIndexOut(t *testing.T) {
}

func TestAddInArrayListAtIndexCero(t *testing.T) {
}

func TestAddInArrayListAtLastIndex(t *testing.T) {

}

func TestAddInArrayListAtIndex(t *testing.T) {
}

func TestGetFirstInEmptyArrayList(t *testing.T) {
}

func TestGetFirstInArrayList(t *testing.T) {
}

func TestGetLastInEmptyArrayList(t *testing.T) {
}

func TestGetLastInArrayList(t *testing.T) {
}

func TestGetAtInEmptyArrayList(t *testing.T) {
}

func TestGetAtInArrayListFirstHalf(t *testing.T) {
}

func TestGetAtInArrayListSecondHalf(t *testing.T) {
}

func TestGetAtInArrayListLastIndex(t *testing.T) {

}

func TestGetAtInArrayListOutOfBounds(t *testing.T) {

}

func TestGetAtInArrayListOutOfBoundsNegative(t *testing.T) {

}

func TestRemoveFirstInEmptyArrayList(t *testing.T) {

}

func TestRemoveFirstInArrayList(t *testing.T) {
}

func TestRemoveLastInEmptyArrayList(t *testing.T) {

}

func TestRemoveLastInArrayList(t *testing.T) {

}

func TestRemoveInEmptyArrayList(t *testing.T) {

}

func TestRemoveInArrayList(t *testing.T) {

}

func TestRemoveInArrayListOutOfBounds(t *testing.T) {
	arr := []int{0, 1, 2}

}

func TestRemoveInArrayListOutOfBoundsNegative(t *testing.T) {
	arr := []int{0, 1, 2}

}

func TestRemoveInArrayListAtStart(t *testing.T) {
	arr := []int{0, 1, 2}

}

func TestRemoveInArrayListAtEnd(t *testing.T) {
	arr := []int{0, 1, 2}

}

func TestRemoveFirstInArrayListToEmpty(t *testing.T) {

}

func TestRemoveLastInArrayListToEmpty(t *testing.T) {

}

func TestRemoveInArrayListToEmpty(t *testing.T) {

}

func TestClearEmptyArrayList(t *testing.T) {

}

func TestClearArrayList(t *testing.T) {

}

func TestContainsInEmptyArrayList(t *testing.T) {

}

func TestContainsInArrayList(t *testing.T) {

}

func TestContainsInArrayListNotInArrayList(t *testing.T) {

}

func TestArrayListContainsWithoutComparator(t *testing.T) {

}

func TestIndexOfInEmptyArrayList(t *testing.T) {

}

func TestIndexOfInArrayList(t *testing.T) {

}

func TestIndexOfInArrayListNotInArrayList(t *testing.T) {

}

func TestArrayListIndexOfWithoutComparator(t *testing.T) {

}

func TestToArrayInEmptyArrayList(t *testing.T) {

}

func TestForEachInEmptyArrayList(t *testing.T) {

}

func TestForEachInArrayList(t *testing.T) {

}

func TestForEachInArrayListNilFunction(t *testing.T) {

}

func TestArrayListFromArray(t *testing.T) {

}

func TestArrayListFromArrayEmptyArray(t *testing.T) {

}

func TestMapInEmptyArrayList(t *testing.T) {

}

func TestMapInArrayList(t *testing.T) {

}

func TestMapInArrayListNilFunction(t *testing.T) {

}

func TestFilterInEmptyArrayList(t *testing.T) {

}

func TestFilterInArrayList(t *testing.T) {

}

func TestFilterInArrayListNilFunction(t *testing.T) {

}

func TestSetEmptyArrayListReturnError(t *testing.T) {

}

func TestSetArrayList(t *testing.T) {

}

func TestSetArrayListOutOfBounds(t *testing.T) {

}

func TestSetArrayListOutOfBoundsNegative(t *testing.T) {

}

func TestSetArrayListAtStart(t *testing.T) {

}

func TestSetArrayListAtEnd(t *testing.T) {

}

func TestAddArrayInEmptyArrayList(t *testing.T) {

}

func TestAddArrayInArrayList(t *testing.T) {

}

func TestArrayListAddArrayWithEmptyArray(t *testing.T) {

}

func TestAddAllAtStartInEmptyArrayList(t *testing.T) {

}

func TestAddAllAtStartInArrayList(t *testing.T) {

}

func TestArrayListAddAllAtWithEmptyArray(t *testing.T) {

}

func TestArrayListAddAllAtWithNilArray(t *testing.T) {

}

func TestAddAllAtEndInArrayList(t *testing.T) {

}

func TestAddAllAtMiddleInArrayList(t *testing.T) {

}
