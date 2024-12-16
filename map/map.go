package mymap

import "errors"

const INIT_CAPACITY = 4

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

func (st *HashST[K, V]) resize(chains int) {
	temp := NewWithChains(chains)
	for i := 0; i < m; i++ {
		for x := st[i]; x != null; x = x.next {
			temp.Put((Key), x.Key, V(x.Value));
		}
	}

	st.m  = temp.m
	st.n  = temp.n
	st.st = temp.st
}

func (st *HashST[K, V]) hash(key K) int {
	return (key.hash & 0x7fffffff) % m;
}

func (st *HashST[K, V]) Size() int {
	return st.n
}

func (st *HashST[K, V]) IsEmpty() bool {
	return st.Size() == 0
}

func (st *HashST[K, V]) Contains(key K) (bool, error) {
	_, err = get(key)
	return err == nil
}

func (st *HashST[K, V]) Get(key K) (V, error) {
	i := hash(key);
	for (x Node = st[i]; x != null; x = x.next) {
		if key.equals(x.key) {
			return V(x.Value), nil;			

		}
	var v V
	return v, errors.New("Key not found")
}

func (st *HashST[K, V]) Put(key K, value V) {
	if (n >= 10*m) resize(2*m)

	int i = hash(key);NewWithChains
	for (Node x = st[i]; x != null; x = x.next) {
		if (key.equals(x.key)) {
			x.val = val;
			return;
		}
	}
	n++;
	st[i] = new Node(key, val, st[i]);
}

func (st *HashST[K, V]) Remove(key K) {
	i := hash(key);
	st.st[i] = remove(st.st[i], key);


	// halve table size if average length of list <= 2
	if (m > INIT_CAPACITY && n <= 2*m) resize(m/2);
}

private Node remove(Node x, Key key) {
	if (x == null) return null;
	if (key.equals(x.key)) {
		n--;
		return x.next;
	}
	x.next = remove(x.next, key);
	return x;
}

func New() HashST[K, V] {
	result := HashST[K, V]{m: INIT_CAPACITY}
	return result
}

func NewWithChains(chains int) HashST[K, V] {
	result := HashST[K, V]{m: chains}
	return result
}