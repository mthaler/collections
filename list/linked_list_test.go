package mylist

import (
	"collections/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	var l LinkedList[int]
	assert.True(t, l.IsEmpty())
	l.Append(1)
	assert.False(t, l.IsEmpty())
}

func TestAppend(t *testing.T) {
	var l LinkedList[int]
	l.Append(1)
	l.Append(2)
	l.Append(3)
	assert.Equal(t, []int{1, 2, 3}, util.ToSclice[int](l))
}

func TestPrepend(t *testing.T) {
	var l LinkedList[int]
	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)
	assert.Equal(t, []int{3, 2, 1}, util.ToSclice[int](l))
}

func TestInsert(t *testing.T) {
	var l LinkedList[int]
	l.Append(1)
	l.Insert(1, 2)
	l.Append(3)
	assert.Equal(t, []int{1, 2, 3}, util.ToSclice[int](l))
}

func TestDeleteFront(t *testing.T) {
	var l LinkedList[int]
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.DeleteFront()
	assert.Equal(t, []int{2, 3}, util.ToSclice[int](l))
}

func TestDeleteBack(t *testing.T) {
	var l LinkedList[int]
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.DeleteBack()
	assert.Equal(t, []int{1, 2}, util.ToSclice[int](l))
}

func TestDelete(t *testing.T) {
	var l LinkedList[int]
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Delete(1)
	assert.Equal(t, []int{1, 3}, util.ToSclice[int](l))
}
