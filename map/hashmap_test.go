package mymap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	h := New[int, int]()
	assert.True(t, h.IsEmpty())
	h.Put(1, 1234)
	assert.False(t, h.IsEmpty())
}

func TestSize(t *testing.T) {
	h := New[int, int]()
	h.Put(1, 1)
	h.Put(2, 2)
	h.Put(3, 3)
	assert.Equal(t, 3, h.Size())
}
