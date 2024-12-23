package mymap

import (
	"errors"

	"github.com/mitchellh/hashstructure/v2"
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
	temp := NewWithChains[K, V](chains)
	for i := 0; i < st.m; i++ {
		for x := st.st[i]; x != nil; x = x.next {
			temp.Put(x.Key.(K), x.Value.(V))
		}
	}

	st.m = temp.m
	st.n = temp.n
	st.st = temp.st
}

func (st *HashST[K, V]) hash(key K) int {
	hash, err := hashstructure.Hash(key, hashstructure.FormatV2, nil)
	if err != nil {
		panic(err)
	}
	m := uint64(st.m)
	h := int((hash & 0x7fffffff) % m)
	return h
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
		if key == x.Key {
			return x.Value.(V), nil

		}
	}
	var v V
	return v, errors.New("Key not found")
}

func (st *HashST[K, V]) Put(key K, value V) {
	if st.n >= 10*st.m {
		st.resize(2 * st.m)
	}

	i := st.hash(key)
	NewWithChains
	for x := st.st[i]; x != nil; x = x.next {
		if key == x.Key {
			x.Value = value
			return
		}
	}
	st.n++
	st.st[i] = &Node{Key: key, Value: value, next: st.st[i]}
}

func (st *HashST[K, V]) Remove(key K) {
	i := st.hash(key)
	st.st[i] = st.remove(*st.st[i], key)

	// halve table size if average length of list <= 2
	if st.m > INIT_CAPACITY && st.n <= 2*st.m {
		st.resize(st.m / 2)
	}
}

func (st *HashST[K, V]) remove(x Node, key K) *Node {
	if key == x.Key {
		st.n--
		return x.next
	}
	x.next = st.remove(*x.next, key)
	return &x
}

func New[K comparable, V any]() HashST[K, V] {
	result := HashST[K, V]{m: INIT_CAPACITY}
	return result
}

func NewWithChains[K comparable, V any](chains int) HashST[K, V] {
	result := HashST[K, V]{m: chains}
	return result
}
