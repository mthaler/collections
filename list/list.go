package list

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

func (list *LinkedList[T]) Append(data T) {
	newNode := &Node[T]{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *LinkedList[T]) Prepend(data T) {
	if list.head == nil {
		newNode := &Node[T]{data: data, next: nil}
		list.head = newNode
		return
	}

	newNode := &Node[T]{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkedList[T]) DeleteFront() {
	if list.head != nil {
		list.head = list.head.next
		fmt.Printf("Head node of linked list has been deleted\n")
		return
	}
}

func (list *LinkedList[T]) DeleteLast() {
	if list.head == nil {
		fmt.Printf("Linked list is empty\n")
	}

	if list.head.next == nil {
		list.head = nil
		fmt.Printf("Last node of linked list has been deleted\n")
		return
	}

	var current *Node[T] = list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil

	fmt.Printf("Last node of linked list has been deleted")
}

func (list *LinkedList[T]) CountNodes() int {
	count := 0
	current := list.head
	for current != nil {
		current = current.next
		count++
	}
	return count
}

func (list *LinkedList[T]) FindNodeAt(index int) *Node[T] {
	var count int = 0
	var current *Node[T] = list.head

	for current != nil {
		count++
		current = current.next
	}

	if index <= 0 || index > count {
		return nil
	}

	current = list.head
	for count = 1; count < index; count++ {
		current = current.next
	}
	return current
}

func (list *LinkedList[T]) ToSlice() []T {
	var result []T = make([]T, 0)
	var current *Node[T] = list.head
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
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
