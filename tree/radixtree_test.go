package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	r := createTestTree()
	assert.Equal(t, []string{"romane", "romanus", "romulus", "rubens", "ruber", "rubicon", "rubicundus"}, r.Keys())
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
