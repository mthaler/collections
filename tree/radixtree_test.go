package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	r := createTestTree()
	assert.True(t, r.Contains("romane"))
	assert.False(t, r.Contains("test"))
}

func TestSize(t *testing.T) {
	r := createTestTree()
	assert.Equal(t, 7, r.Size())
}

func TestIsEmpty(t *testing.T) {
	r := TrieST{}
	assert.True(t, r.IsEmpty())
	r2 := createTestTree()
	assert.False(t, r2.IsEmpty())
}

func TestKeys(t *testing.T) {
	r := createTestTree()
	assert.Equal(t, []string{"romane", "romanus", "romulus", "rubens", "ruber", "rubicon", "rubicundus"}, r.Keys())
}

func TestKeysWithPrefix(t *testing.T) {
	r := createTestTree()
	assert.Equal(t, []string{"rubens", "ruber", "rubicon", "rubicundus"}, r.KeysWithPrefix("ru"))
}

func TestKeysThatMatch(t *testing.T) {
	r := createTestTree()
	assert.Equal(t, []string{"romane"}, r.KeysThatMatch("roman."))
}

func createTestTree() TrieST {
	r := TrieST{}
	r.Put("romane", 1)
	r.Put("romanus", 2)
	r.Put("romulus", 3)
	r.Put("rubens", 4)
	r.Put("ruber", 5)
	r.Put("rubicon", 6)
	r.Put("rubicundus", 7)
	return r
}
