package queue

type Queue[T any] interface {
	IsEmpty() bool
	Append(data T)
	Prepend(data T)
	Insert(n int, data T)
}
