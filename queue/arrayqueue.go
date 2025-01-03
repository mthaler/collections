package queue

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
