package set

import mymap "collections/map"

type HashSet[K comparable, V any] struct {
	m mymap.HashST[K, V]
}

func (s *HashSet[K, V]) Size() int {
	return s.Size()
}

func (s *HashSet[K, V]) IsEmpty() bool {
	return s.IsEmpty()
}
