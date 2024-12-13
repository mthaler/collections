package list

import "testing"

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
