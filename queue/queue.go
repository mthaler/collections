package queue

type Queue[T any] interface {
	IsEmpty() bool
	Enqueue(item T)
	Dequeue() T
}
