package mymap

type Map[K comparable, V any] interface {
	Size() int
	IsEmpty() bool
	Contains(key K)
	Get(key K) (V, error)
	Put(key K, value V)
	Remove(key K)
}
