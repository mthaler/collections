package list

import (
	"reflect"
	"testing"
)

func TestRadixTree_Keys(t *testing.T) {
	tree := NewList("alligator", "baby", "chicken")
	tree.Prepend("doll")
	s := tree.ToSlice()
	expected := []string{"doll", "alligator", "baby", "chicken"}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("s = %v, expected %v", s, expected)
	}
}
