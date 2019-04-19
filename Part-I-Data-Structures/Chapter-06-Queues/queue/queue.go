package queue

type Queue struct {
	head *node
	tail *node
}

type node struct {
	Value int
	Next  *node
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue) Enqueue(v int) {
	if q.head == nil {
		q.head = &node{Value: v}
		q.tail = q.head
	} else {
		q.tail.Next = &node{Value: v}
		q.tail = q.tail.Next
	}
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
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

func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.head.Value, true
}
