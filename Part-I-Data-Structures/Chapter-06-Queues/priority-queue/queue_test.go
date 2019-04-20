package queue

import (
	"fmt"
)

func ExamplePriorityQueue() {
	q := &PriorityQueue{}

	q.Enqueue(8)
	q.Enqueue(19)
	q.Enqueue(23)
	q.Enqueue(78)
	q.Enqueue(12)
	q.Enqueue(11)
	q.Enqueue(15)

	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		fmt.Printf(" %d", v)
	}

	// Output:
	// 8 11 12 15 19 23 78
}
