package queue

import (
	"fmt"
)

func ExampleQueue() {
	q := &Queue{}

	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	if v, ok := q.Peek(); ok {
		fmt.Println(v)
	}

	for v, ok := q.Dequeue(); ok; {
		fmt.Printf(" %d", v)
		v, ok = q.Dequeue()
	}
	fmt.Println()

	_, ok := q.Peek()
	fmt.Println(ok)

	// Output:
	// 0
	//  0 1 2 3 4 5 6 7 8 9
	// false
}
