package bag

/*
Port of Robert Sedgewicks hash table code to Go
*/
type Bag[T any] struct {
	first *Node[T] // beginning of bag
	n     int      // number of elements in bag
}

type Node[T any] struct {
	item T
	next *Node[T]
}

/**
 * Returns true if this bag is empty.
 *
 * @return {@code true} if this bag is empty;
 *         {@code false} otherwise
 */
func (b *Bag[T]) IsEmpty() bool {
	return b.first == nil
}

/**
 * Returns the number of items in this bag.
 *
 * @return the number of items in this bag
 */
func (b *Bag[T]) Size() int {
	return b.n
}

/**
 * Adds the item to this bag.
 *
 * @param  item the item to add to this bag
 */
func (b *Bag[T]) Add(t T) {
	oldFirst := b.first
	var first Node[T]
	first.item = t
	first.next = oldFirst
	b.first = &first
	b.n++
}

func (list *Bag[T]) ToSlice() []T {
	var result []T = make([]T, 0)
	var current *Node[T] = list.first
	for current != nil {
		result = append(result, current.item)
		current = current.next
	}
	return result
}
