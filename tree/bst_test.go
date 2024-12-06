package tree

import (
	"reflect"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	var tt Tree[int] = Tree[int]{}
	tt.Insert(3)
	tt.Insert(2)
	tt.Insert(1)
	var result = tt.TraversePostOrder()
	if !reflect.DeepEqual(result, []int{1, 2, 3}) {
		t.Errorf("Result = %v, expected %v", result, []int{1, 2, 3})
	}
}
