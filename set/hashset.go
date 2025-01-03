package set

import mymap "collections/map"

type HashSet[K comparable] struct {
	m mymap.HashST[K, any]
}

func (s *HashSet[K]) Size() int {
	return s.m.Size()
}

func (s *HashSet[K]) IsEmpty() bool {
	return s.m.IsEmpty()
}

func (s *HashSet[K]) Contains(key K) bool {
	return s.m.Contains(key)
}

func (s *HashSet[K]) Put(key K) {
	s.m.Put(key, struct{}{})
}

func (s *HashSet[K]) Remove(key K) {
	s.m.Remove(key)
}

func New[K comparable]() HashSet[K] {
	st := mymap.New[K, any]()
	return HashSet[K]{m: st}
}
