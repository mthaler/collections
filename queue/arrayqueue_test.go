package queue

import (
	"collections/util"
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

func TestToSlice(t *testing.T) {
	var q ArrayQueue[int]
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	result := util.ToSclice[int](q)
	assert.Equal(t, []int{1, 2, 3}, result)
}
