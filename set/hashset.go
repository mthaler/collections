package set

import mymap "collections/map"

type HashSet[K comparable, V any] struct {
	m mymap.HashST[K, V]
}

func (s *HashSet[K, V]) Size() int {
	return s.m.Size()
}

func (s *HashSet[K, V]) IsEmpty() bool {
	return s.m.IsEmpty()
}

func (s *HashSet[K, V]) Contains(key K) bool {
	return s.m.Contains(key)
}

func (s *HashSet[K, V]) Put(key K) {
	var value V
	s.m.Put(key, value)
}

func New[K comparable, V any]() HashSet[K, V] {
	st := mymap.New[K, V]()
	return HashSet[K, V]{m: st}
}
