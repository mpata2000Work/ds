package linkedList

type Node struct {
	Next  *Node
	Value any
}

type LinkedList struct {
	First *Node
	Last  *Node
	Size  int
}

func NewList() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func (l *LinkedList) Add(value any) {
	var new_node = new(Node)
	new_node.Value = value

	if l.First == nil {
		l.First = new_node
	} else {
		l.Last.Next = new_node
	}
	l.Last = new_node
	l.Size++
}
