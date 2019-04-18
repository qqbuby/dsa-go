package doubly

import (
	"fmt"
)

func ExampleDoublyLinkedList() {
	list := &List{}

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	for v := range list.ReverseTraverse() {
		fmt.Print(v)
	}
	fmt.Println()

	for i := 0; i < 10; i += 2 {
		list.Remove(i)
	}
	for v := range list.ReverseTraverse() {
		fmt.Print(v)
	}
	fmt.Println()

	// Output:
	// 9876543210
	// 97531
}
