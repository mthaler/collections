package collections

type Collection[T any] interface {
	CreateIterator() Iterator[T]
}
