package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	h := New[int]()
	assert.True(t, h.IsEmpty(), "Map should be empty")
	h.m.Put(1, 1234)
	assert.False(t, h.IsEmpty(), "Map should not be empty")
}

func TestSize(t *testing.T) {
	h := New[int]()
	h.m.Put(1, 1)
	h.m.Put(2, 2)
	h.m.Put(3, 3)
	assert.Equal(t, 3, h.Size())
}
