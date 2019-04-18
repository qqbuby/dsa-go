package singly

import (
	"fmt"
)

func ExampleSinglyLinkedList() {
	list := List{}
	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	for v := range list.Traverse() {
		fmt.Printf("%d", v)
	}
	fmt.Println()

	for v := range list.ReverseTraverse() {
		fmt.Printf("%d", v)
	}
	fmt.Println()

	for i := 0; i < 10; i = i + 2 {
		list.Remove(i)
	}
	for v := range list.Traverse() {
		fmt.Printf("%d", v)
	}
	fmt.Println()

	fmt.Println(list.Contains(1))
	fmt.Println(list.Contains(2))

	// Output:
	// 0123456789
	// 9876543210
	// 13579
	// true
	// false
}
