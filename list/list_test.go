package mylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	var l LinkedList[int]
	assert.True(t, l.IsEmpty())
	l.Append(1)
	assert.False(t, l.IsEmpty())
}
