package collections

import (
	mylist "collections/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasNext(t *testing.T) {
	l := mylist.NewList[int]()
	it := l.CreateIterator()
	assert.False(t, it.HasNext(), "Iterator.HasNext should be false")
	l = mylist.NewList(1, 2, 3)
	it = l.CreateIterator()
	assert.True(t, it.HasNext(), "Iterator.HasNext should be true")
}
