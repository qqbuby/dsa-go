package bst

type Queue struct {
	head *node
	tail *node
}

type node struct {
	Value *Node // Value is a type `Node` of the tree node of BST
	Next  *node
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue) Enqueue(v *Node) {
	if q.head == nil {
		q.head = &node{Value: v}
		q.tail = q.head
	} else {
		q.tail.Next = &node{Value: v}
		q.tail = q.tail.Next
	}
}

func (q *Queue) Dequeue() (*Node, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	ret := q.head.Value
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.Next
	}

	return ret, true
}

func (q *Queue) Peek() (*Node, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	return q.head.Value, true
}
