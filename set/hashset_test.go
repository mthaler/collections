package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	h := New[int, int]()
	assert.True(t, h.IsEmpty(), "Map should be empty")
	h.m.Put(1, 1234)
	assert.False(t, h.IsEmpty(), "Map should not be empty")
}
