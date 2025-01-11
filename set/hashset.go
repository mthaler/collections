package set

import (
	"collections"
	mymap "collections/map"
)

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

type Iterator[K comparable] struct {
	it mymap.Iterator[*mymap.Node]
}

func (it *Iterator[K]) HasNext() bool {
	return it.it.HasNext()
}

func (it *Iterator[K]) Next() K {
	result := it.it.Next()
	return result.Key.(K)
}

func (s *HashSet[K]) CreateIterator() collections.Iterator[K] {
	if s.IsEmpty() {
		return &Iterator[K]{}
	} else {
		return &Iterator[K]{it: s.m.CreateIterator[*mymap.Node]()}
	}
}

func New[K comparable]() HashSet[K] {
	st := mymap.New[K, any]()
	return HashSet[K]{m: st}
}
