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
}
