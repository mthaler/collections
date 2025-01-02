package queue

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.items) < 1 {
		var result T
		return result // returns -1 if the queue is empty
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
