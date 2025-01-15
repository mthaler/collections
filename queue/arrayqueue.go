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
	items []*T
	i     int
}

func (it *Iterator[T]) HasNext() bool {
	return it.items != nil
}

func (it *Iterator[T]) Next() T {
	result := it.items[it.i]
	it.i++
	return *result
}

func (q *ArrayQueue[T]) CreateIterator() collections.Iterator[T] {
	if len(q.items) == 0 {
		it := Iterator[T]{}
		return &it
	} else {
		var it Iterator[T]
		for i, v := range q.items {
			it.items[i] = &v
		}
		return &it
	}
}

func New[T any]() Queue[T] {
	var q ArrayQueue[T]
	return &q
}
