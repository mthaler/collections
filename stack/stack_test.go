package stack

import (
	"reflect"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	var s Stack[int]
	if !s.IsEmpty() {
		t.Errorf("Stack should be empty")
	}
	s.Push(1)
	if s.IsEmpty() {
		t.Errorf("Stack should be empty")
	}
}

func TestSize(t *testing.T) {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Size() != 3 {
		t.Errorf("s.Size() = %d, expected %d", s.Size(), 3)
	}
}

func TestPush(t *testing.T) {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)
	expected := []int{1, 2, 3}
	result := s.ToSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("s.ToSlice() = %v, expected %v", result, expected)
	}
}

func TestTop(t *testing.T) {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)
	var top, _ = s.Top()
	if top != 3 {
		t.Errorf("s.Top() = %d, expected %d", top, 3)
	}
}

func TestToSlice(t *testing.T) {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)
	expected := []int{1, 2, 3}
	result := s.ToSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("s = %v, expected %v", s, expected)
	}
}
