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

func TestContains(t *testing.T) {
	h := New[int, int]()
	h.Put(1, 1)
	h.Put(2, 2)
	h.Put(3, 3)
	assert.True(t, h.Contains(3))
	assert.False(t, h.Contains(4))
}

func TestGet(t *testing.T) {
	h := New[int, int]()
	h.Put(1, 1234)
	v, _ := h.Get(1)
	assert.Equal(t, 1234, v)
}

func TestPut(t *testing.T) {
	h := New[int, int]()
	h.Put(1, 1234)
	v, _ := h.Get(1)
	assert.Equal(t, 1234, v)
}

/*func TestRemove(t *testing.T) {
	h := New[int, int]()
	h.Put(1, 1)
	h.Put(2, 2)
	h.Put(3, 3)
	h.Remove(1)
	assert.Equal(t, []int{1, 3}, util.ToSclice[h]())
}*/
