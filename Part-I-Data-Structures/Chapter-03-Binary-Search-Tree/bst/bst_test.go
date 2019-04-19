package bst

import (
	"fmt"
)

func ExampleBST() {
	values := []int{23, 14, 31, 7, 9}

	t := &Tree{}

	for _, v := range values {
		t.Insert(v)
	}

	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	for v := range t.Postorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	for v := range t.Inorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	for v := range t.BreathFirst() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	// Output:
	// 23 14 7 9 31
	//  9 7 14 31 23
	//  7 9 14 23 31
	//  23 14 31 7 9
}
