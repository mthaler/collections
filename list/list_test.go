package list

import (
	"reflect"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	var l LinkedList[int]
	if !l.IsEmpty() {
		t.Errorf("List should be empty")
	}
	l.Prepend(1)
	if l.IsEmpty() {
		t.Errorf("Bag should not be empty")
	}
}

func TestIPrepend(t *testing.T) {
	var l LinkedList[int]
	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)
	expected := []int{3, 2, 1}
	result := l.ToSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("l = %v, expected %v", result, expected)
	}
}
