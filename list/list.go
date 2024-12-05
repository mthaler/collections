package list

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func (list *LinkedList[T]) insertAtBack(data T) {
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
