package bag

import (
	"reflect"
	"sort"
	"testing"
)

func TestQueue(t *testing.T) {
	var b Bag[int]
	if !b.IsEmpty() {
		t.Errorf("Bag should be empty")
	}
	b.Add(1)
	if b.IsEmpty() {
		t.Errorf("Bag should not be empty")
	}
}

func TestSize(t *testing.T) {
	var b Bag[int]
	b.Add(1)
	b.Add(2)
	b.Add(3)
	if b.Size() != 3 {
		t.Errorf("s.Size() = %d, expected %d", b.Size(), 3)
	}
}

func TestToSlice(t *testing.T) {
	var b Bag[int]
	b.Add(1)
	b.Add(2)
	b.Add(3)
	expected := []int{1, 2, 3}
	result := b.ToSlice()
	sort.Ints(result)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("s = %v, expected %v", result, expected)
	}
}
