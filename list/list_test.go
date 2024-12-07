package list

import (
	"reflect"
	"testing"
)

func TestRadixTree_Keys(t *testing.T) {
	words := []string{"alligator", "baby", "chicken"}
	tree := makeList(words)
	tree.Prepend("doll")
	s := tree.ToSlice()
	expected := []string{"doll", "alligator", "baby", "chicken"}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("s = %v, expected %v", s, expected)
	}
}

func makeList(values []string) LinkedList[string] {
	if len(values) == 0 {
		return LinkedList[string]{}

	}
	current := &Node[string]{data: values[0]}
	result := LinkedList[string]{head: current}
	for i, v := range values {
		if i == 0 {
			// do nothing
		} else {
			n := &Node[string]{data: v}
			current.next = n
			current = n
		}
	}
	return result
}
