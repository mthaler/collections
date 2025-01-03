package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	var q Queue[int]
	assert.True(t, q.IsEmpty())
	q.Enqueue(1)
	assert.False(t, q.IsEmpty())
}

func TestQueue(t *testing.T) {
	var q Queue[int] = Queue[int]{}
	q.Enqueue(1)
	var result = q.Dequeue()
	if result != 1 {
		t.Errorf("q.Dequeue() = %d, expected %d", result, 1)
	}
}
