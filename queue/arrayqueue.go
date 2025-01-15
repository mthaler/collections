package queue

import "collections"

type ArrayQueue[T any] struct {
	items []T
}

func (q *ArrayQueue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *ArrayQueue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *ArrayQueue[T]) Dequeue() T {
	if len(q.items) < 1 {
		var result T
		return result // returns -1 if the queue is empty
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

type Iterator[T any] struct {
	q ArrayQueue[T]
}

func (it *Iterator[T]) HasNext() bool {
	return !it.q.IsEmpty()
}

func (it *Iterator[T]) Next() T {
	return it.q.Dequeue()
}

func (q ArrayQueue[T]) CreateIterator() collections.Iterator[T] {
	if q.IsEmpty() {
		it := Iterator[T]{}
		return &it
	} else {
		var it Iterator[T]
		it.q = q
		return &it
	}
}

func New[T any]() ArrayQueue[T] {
	return ArrayQueue[T]{}
}
