package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	s := New[int]()
	assert.True(t, s.IsEmpty())
	s.m.Put(1, 1234)
	assert.False(t, s.IsEmpty())
}

func TestSize(t *testing.T) {
	s := New[int]()
	s.m.Put(1, 1)
	s.m.Put(2, 2)
	s.m.Put(3, 3)
	assert.Equal(t, 3, s.Size())
}

func TestContains(t *testing.T) {
	s := New[int]()
	s.Put(1)
	s.Put(2)
	s.Put(3)
	assert.True(t, s.Contains(3))
	assert.False(t, s.Contains(4))
}

func TestPut(t *testing.T) {
	s := New[int]()
	assert.False(t, s.Contains(1))
	s.Put(1)
	assert.True(t, s.Contains(1))
}

func TestRemove(t *testing.T) {
	s := New[int]()
	s.Put(1)
	s.Put(2)
	s.Put(3)
	s.Remove(1)
	assert.False(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.True(t, s.Contains(3))
}
