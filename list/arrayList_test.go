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
