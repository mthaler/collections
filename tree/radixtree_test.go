package tree

import (
	"collections/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	r := createTestTree()
	assert.Equal(t, 1, r.Get("romane"))
}

func TestContains(t *testing.T) {
	r := createTestTree()
	assert.True(t, r.Contains("romane"))
	assert.False(t, r.Contains("test"))
}

func TestPut(t *testing.T) {
	r := createTestTree()
	r.Put("test", 42)
	assert.Equal(t, 42, r.Get("test"))
}

func TestSize(t *testing.T) {
	r := createTestTree()
	assert.Equal(t, 7, r.Size())
}

func TestIsEmpty(t *testing.T) {
	r := TrieST[int]{}
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

func TestLongestPrefixOf(t *testing.T) {
	r := TrieST[int]{}
	r.Put("romane", 1)
	r.Put("romanus", 2)
	r.Put("romulus", 3)
	assert.Equal(t, "rom", r.LongestPrefixOf(""))
}

func TestDelete(t *testing.T) {
	r := createTestTree()
	r.Delete("romane")
	assert.Equal(t, 6, r.Size())
}

func TestToSlice(t *testing.T) {
	r := createTestTree()
	keys := make([]string, 0)
	values := make([]int, 0)
	result := util.ToSclice[KV[int]](r)
	for _, kv := range result {
		keys = append(keys, kv.Key)
		values = append(values, kv.Value)
	}
	assert.Equal(t, []string{"romane", "romanus", "romulus", "rubens", "ruber", "rubicon", "rubicundus"}, keys)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, values)
}

func createTestTree() TrieST[int] {
	r := TrieST[int]{}
	r.Put("romane", 1)
	r.Put("romanus", 2)
	r.Put("romulus", 3)
	r.Put("rubens", 4)
	r.Put("ruber", 5)
	r.Put("rubicon", 6)
	r.Put("rubicundus", 7)
	return r
}
