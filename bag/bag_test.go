package bag

import "testing"

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
