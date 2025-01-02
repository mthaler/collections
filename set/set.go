package set

type Set[K comparable, V any] interface {
	Size() int
	IsEmpty() bool
	Contains(key K)
	Put(key K, value V)
	Remove(key K)
}
