package mymap

import (
	"errors"

	"github.com/mitchellh/hashstructure"
)

const INIT_CAPACITY = 4

type Node struct {
	Key   any
	Value any
	next  *Node
}

type HashST[K comparable, V any] struct {
	n  int     // number of key-value pairs
	m  int     // number of chains
	st []*Node // array of linked-list symbol tables
}

func (st *HashST[K, V]) resize(chains int) {
	temp := NewWithChains{K, V}(chains)
	for i := 0; i < st.m; i++ {
		for x := st.st[i]; x != nil; x = x.next {
			temp.Put((Key), x.Key, V(x.Value))
		}
	}

	st.m = temp.m
	st.n = temp.n
	st.st = temp.st
}

func (st *HashST[K, V]) hash(key K) int {
	hash, err := hashstructure.Hash(key, hashstructure.FormatV2, nil)
	return (hash & 0x7fffffff) % st.m
}

func (st *HashST[K, V]) Size() int {
	return st.n
}

func (st *HashST[K, V]) IsEmpty() bool {
	return st.Size() == 0
}

func (st *HashST[K, V]) Contains(key K) bool {
	_, err := st.Get(key)
	return err == nil
}

func (st *HashST[K, V]) Get(key K) (V, error) {
	i := st.hash(key)
	for x := st.st[i]; x != nil; x = x.next {
		if key.equals(x.key) {
			return V(x.Value), nil

		}
	}
	var v V
	return v, errors.New("Key not found")
}

func (st *HashST[K, V]) Put(key K, value V) {
	if st.n >= 10*st.m {
		st.resize(2 * st.m)
	}

	i := hash(key)
	NewWithChains
	for x := st[i]; x != null; x = x.next {
		if key.equals(x.key) {
			x.Value = val
			return
		}
	}
	st.n++
	st[i] = Node(key, val, st[i])
}

func (st *HashST[K, V]) Remove(key K) {
	i := hash(key)
	st.st[i] = remove(st.st[i], key)

	// halve table size if average length of list <= 2
	if m > INIT_CAPACITY && n <= 2*m {
		st.resize(m / 2)
	}
}

func (st *HashST[K, V]) remove(x Node, key K) *Node {
	if key.equals(x.Key) {
		st.n--
		return x.next
		x.next = st.remove(x.next, key)
		return x
	}
}

func New[K comparable, V any]() HashST[K, V] {
	result := HashST[K, V]{m: INIT_CAPACITY}
	return result
}

func NewWithChains[K comparable, V any](chains int) HashST[K, V] {
	result := HashST[K, V]{m: chains}
	return result
}
