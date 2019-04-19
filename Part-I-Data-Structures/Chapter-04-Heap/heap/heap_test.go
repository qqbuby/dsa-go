package heap

import (
	"fmt"
)

func ExampleHeap() {
	h := &Heap{}
	for _, v := range []int{8, 19, 12, 23, 78} {
		h.Add(v)
	}

	for v := range h.Traverse() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	h.Remove(8)
	for v := range h.Traverse() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	h.Add(8)
	for v := range h.Traverse() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	h.Add(15)
	h.Add(11)
	for v := range h.Traverse() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	// Output:
	//  8 19 12 23 78
	//  12 19 78 23
	//  8 12 78 23 19
	//  8 12 11 23 19 78 15
}
