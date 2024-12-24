package collections

type Collection interface {
	createIterator() Iterator
}
