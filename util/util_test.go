package util

import (
	mylist "collections/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSlice(t *testing.T) {
	l := mylist.New[int]()
	assert.True(t, l.IsEmpty())
	l = mylist.New[int](1, 2, 3)
	result := ToSclice[int](l)
	expected := []int{1, 2, 3}
	assert.Equal(t, expected, result)
}
