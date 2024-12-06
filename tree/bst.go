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
