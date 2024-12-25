package util

import (
	mylist "collections/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSlice(t *testing.T) {
	l := mylist.NewList[int]()
	assert.True(t, h.IsEmpty(), "Map should be empty")
	h.Put(1, 1234)
	assert.False(t, h.IsEmpty(), "Map should not be empty")
}
