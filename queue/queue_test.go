package queue

import "testing"

func TestQueue(t *testing.T) {
	var q Queue[int] = Queue[int]{}
	q.Enqueue(1)
	var result = q.Dequeue()
	if result != 1 {
		t.Errorf("q.Dequeue() = %d, expeczed %d", result, 1)
	}
}
