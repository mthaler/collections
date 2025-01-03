package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	var q ArrayQueue[int]
	assert.True(t, q.IsEmpty())
	q.Enqueue(1)
	assert.False(t, q.IsEmpty())
}

func TestQueue(t *testing.T) {
	var q ArrayQueue[int]
	q.Enqueue(1)
	assert.Equal(t, 1, q.Dequeue())
}
