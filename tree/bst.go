package tree

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Tree[T numbers] struct {
	root *Node[T]
}

type Node[T numbers] struct {
	key   T
	left  *Node[T]
	right *Node[T]
}

func (t *Tree[T]) Insert(data T) {
	if t.root == nil {
		t.root = &Node[T]{key: data}
	} else {
		t.root.insert(data)
	}
}

func (n *Node[T]) insert(data T) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node[T]{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node[T]{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *Tree[T]) traversePreOrder() []T {
	return t.root.traversePreOrder()
}

func (n *Node[T]) traversePreOrder() []T {
	if n == nil {
		return make([]T, 0)
	} else {
		result := make([]T, 0)
		result = append(result, n.key)
		result = append(result, n.left.traversePreOrder()...)
		result = append(result, n.right.traversePreOrder()...)
		return result
	}
}

func (n *Node[T]) traversePostOrder() []T {
	if n == nil {
		return make([]T, 0)
	} else {
		result := make([]T, 0)
		result = append(result, n.left.traversePostOrder()...)
		result = append(result, n.right.traversePostOrder()...)
		result = append(result, n.key)
		return result
	}
}
