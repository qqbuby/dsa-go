// Unlike a standard queue where items are ordered
// in terms of who arrived first, a priority queue
// determines the order of its items by using a form
// of custom comparer to see which item has the
// highest priority.
//
// A sensible implementation of a priority queue is to
// use a heap data structure.
package queue

type PriorityQueue struct {
	s []int
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (q *PriorityQueue) IsEmpty() bool {
	return len(q.s) == 0
}

func (q *PriorityQueue) Enqueue(v int) {
	q.s = append(q.s, v)

	// min heapify
	i := len(q.s) - 1
	p := parent(i)
	for i > 0 {
		if q.s[i] < q.s[p] {
			q.s[i], q.s[p] = q.s[p], q.s[i]
		}

		i = p
		p = parent(i)
	}
}

func (q *PriorityQueue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	v := q.s[0]

	q.s[0] = q.s[len(q.s)-1]
	q.s = q.s[:len(q.s)-1]

	// min heapify
	p := 0
	l := left(p)
	r := right(p)
	m := l
	for l < len(q.s) || r < len(q.s) {
		if l < len(q.s) && r < len(q.s) {
			if q.s[l] < q.s[r] {
				m = l
			} else {
				m = r
			}
		} else if l < len(q.s) {
			m = l
		} else {
			m = r
		}

		if q.s[p] > q.s[m] {
			q.s[p], q.s[m] = q.s[m], q.s[p]
		} else {
			break
		}

		p = m
		l = left(p)
		r = right(p)
	}

	return v, true
}

func (q *PriorityQueue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	return q.s[0], true
}
