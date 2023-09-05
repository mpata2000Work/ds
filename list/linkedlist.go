package list

type Node[T any] struct {
	next, previous *Node[T]
	value          T
}

type LinkedList[T any] struct {
	head, tail *Node[T]
	size       int
}

func NewList[T any]() *LinkedList[T] {
	return &LinkedList[T]{nil, nil, 0}
}

func (l LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) Add(value T) {
	var new_node = Node[T]{nil, nil, value}

	if l.head == nil {
		l.head = &new_node
	} else {
		l.tail.next = &new_node
		new_node.previous = l.tail
	}
	l.tail = &new_node
	l.size++
}

func (l LinkedList[T]) Peek() T {
	return l.head.value
}
