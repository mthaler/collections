package util

import "collections"

func ToSclice[T any](c collections.Collection[T]) []T {
	it := c.CreateIterator()
	result := make([]T, 0)
	if it.HasNext() {
		result = append(result, it.Next())
	}
	return result
}
