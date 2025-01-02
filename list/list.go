package mylist

import (
	"collections"
	"fmt"
)

type Node[T any] struct {
	data T
	next *Node[T]
}

type Iterator[T any] struct {
	next *Node[T]
}

func (it *Iterator[T]) HasNext() bool {
	return it.next != nil
}

func (it *Iterator[T]) Next() T {
	result := it.next.data
	it.next = it.next.next
	return result
}

type LinkedList[T any] struct {
	head *Node[T] // beginning of list
	n    int      // number of elements in list
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

func (l *LinkedList[T]) Append(data T) {
	newNode := &Node[T]{data: data, next: nil}

	if l.head == nil {
		l.head = newNode
		l.n++
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	l.n++
}

func (l *LinkedList[T]) Prepend(data T) {
	if l.head == nil {
		newNode := &Node[T]{data: data, next: nil}
		l.head = newNode
		l.n++
		return
	}

	newNode := &Node[T]{data: data, next: l.head}
	l.head = newNode
	l.n++
}

func (l *LinkedList[T]) Insert(n int, data T) error {
	if n < 0 || n > l.n {
		return fmt.Errorf("ndex %d out of bounds", n)
	}
	if n == 0 {
		l.Prepend(data)
		return nil
	} else if n == l.n-1 {
		l.Append(data)
		return nil
	} else {
		n := l.head
		for ; n.next != nil; n = n.next {
			n.next = &Node[T]{data: data, next: nil}
		}
		l.n++
		return nil
	}
}

func (l *LinkedList[T]) DeleteFront() {
	if l.head != nil {
		l.head = l.head.next
		fmt.Printf("Head node of linked list has been deleted\n")
		return
	}
}

func (l *LinkedList[T]) DeleteBack() {
	if l.head == nil {
		fmt.Printf("Linked list is empty\n")
	}

	if l.head.next == nil {
		l.head = nil
		fmt.Printf("Last node of linked list has been deleted\n")
		return
	}

	var current *Node[T] = l.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil

	fmt.Printf("Last node of linked list has been deleted")
}

func (l *LinkedList[T]) FindNodeAt(index int) *Node[T] {
	var count int = 0
	var current *Node[T] = l.head

	for current != nil {
		count++
		current = current.next
	}

	if index <= 0 || index > count {
		return nil
	}

	current = l.head
	for count = 1; count < index; count++ {
		current = current.next
	}
	return current
}

func (l LinkedList[T]) CreateIterator() collections.Iterator[T] {
	if l.head != nil {
		it := Iterator[T]{next: l.head}
		return &it
	} else {
		it := Iterator[T]{}
		return &it
	}
}

func NewList[T any](values ...T) LinkedList[T] {
	if len(values) == 0 {
		return LinkedList[T]{}
	}
	current := &Node[T]{data: values[0]}
	result := LinkedList[T]{head: current}
	for i, v := range values {
		if i == 0 {
			// do nothing
		} else {
			n := &Node[T]{data: v}
			current.next = n
			current = n
		}
	}
	return result
}
