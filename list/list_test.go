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

func TestPrepend(t *testing.T) {
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

func TestAppend(t *testing.T) {
	var l LinkedList[int]
	l.Append(1)
	l.Append(2)
	l.Append(3)
	expected := []int{1, 2, 3}
	result := l.ToSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("l = %v, expected %v", result, expected)
	}
}

func TestDeleteFront(t *testing.T) {
	var l LinkedList[int]
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.DeleteFront()
	expected := []int{2, 3}
	result := l.ToSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("l = %v, expected %v", result, expected)
	}
}

func TestDeleteBack(t *testing.T) {
	var l LinkedList[int]
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.DeleteBack()
	expected := []int{1, 2}
	result := l.ToSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("l = %v, expected %v", result, expected)
	}
}
