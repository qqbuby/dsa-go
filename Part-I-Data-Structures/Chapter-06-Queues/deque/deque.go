// A double ended is commonly known as a deque that
// allows you to access the items at both the front
// and back of the queue.
package deque

type Deque struct {
	front *node
	back  *node
}

type node struct {
	Value int
	Next  *node
	Prev  *node
}

func (q *Deque) IsEmpty() bool {
	return q.front == nil
}

func (q *Deque) EnqueueFront(v int) {
	n := &node{Value: v}

	if q.front == nil {
		q.front = n
		q.back = q.front
	} else {
		n.Next = q.front
		q.front.Prev = n
		q.front = n
	}
}

func (q *Deque) EnqueueBack(v int) {
	n := &node{Value: v}

	if q.front == nil {
		q.front = n
		q.back = q.front
	} else {
		n.Prev = q.back
		q.back.Next = n
		q.back = n
	}
}

func (q *Deque) DequeueFront() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	v := q.front.Value

	if q.front == q.back {
		q.front = nil
		q.back = nil
	} else {
		q.front = q.front.Next
		q.front.Prev = nil
	}

	return v, true
}

func (q *Deque) DequeueBack() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	v := q.back.Value

	if q.front == q.back {
		q.front = nil
		q.back = nil
	} else {
		q.back = q.back.Prev
		q.back.Next = nil
	}

	return v, true
}

func (q *Deque) PeekFront() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	v := q.front.Value

	return v, true
}

func (q *Deque) PeekBack() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	v := q.back.Value

	return v, true
}

func (q *Deque) Traverse() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		p := q.front
		for p != nil {
			ch <- p.Value
			p = p.Next
		}
	}()

	return ch
}
