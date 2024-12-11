package stack

import (
	"reflect"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	var s Stack[int]
	if !s.IsEmpy() {
		t.Errorf("Stack should be empty")
	}
	s.Push(1)
	if s.IsEmpy() {
		t.Errorf("Stack should be empty")
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
		t.Errorf("s = %v, expected %v", s, expected)
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
