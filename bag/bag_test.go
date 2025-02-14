package bag

import (
	"collections/util"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	var b Bag[int]
	assert.True(t, b.IsEmpty())
	b.Add(1)
	assert.False(t, b.IsEmpty())
}

func TestSize(t *testing.T) {
	var b Bag[int]
	b.Add(1)
	b.Add(2)
	b.Add(3)
	assert.Equal(t, 3, b.Size())
}

func TestAdd(t *testing.T) {
	var b Bag[int]
	b.Add(1)
	b.Add(2)
	b.Add(3)
	result := util.ToSclice[int](b)
	sort.Ints(result)
	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestToSlice(t *testing.T) {
	var b Bag[int]
	b.Add(1)
	b.Add(2)

	b.Add(3)
	expected := []int{1, 2, 3}
	result := util.ToSclice[int](b)
	sort.Ints(result)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("s = %v, expected %v", result, expected)
	}
}
