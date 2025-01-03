package mylist

type List[T any] interface {
	IsEmpty() bool
	Append(data T)
	Prepend(data T)
	Insert(n int, data T)
}
