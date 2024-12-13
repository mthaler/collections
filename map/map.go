package mymap

import "errors"

type Node struct {
	Key   any
	Value any
	next  *Node
}

type HashST[K comparable, V any] struct {
	n  int    // number of key-value pairs
	m  int    // number of chains
	st []Node // array of linked-list symbol tables
}

func (st *HashST[K, V]) Size() int {
	return st.n
}

func (st *HashST[K, V]) IsEmpty() bool {
	return st.Size() == 0
}

func (st *HashST[K, V]) Contains(key K) (bool, error) {
	if key == nil {
		return false, errors.New("argument to contains() is nil")
	}
	return get(key) != nil
}
