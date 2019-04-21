package avl

import (
	"fmt"
)

func ExampleLeftRotationAndRightRotationAVLTree() {
	t := &Tree{}
	for _, v := range []int{14, 8, 24, 2, 11, 1} {
		t.Insert(v)
	}

	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	t.rightRotation(t.root.Left)
	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	t.leftRotation(t.root.Left)
	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	// Output:
	// 14 8 2 1 11 24
	//  14 2 1 8 11 24
	//  14 8 2 1 11 24
}

func ExampleLeftRightRotationAVLTree() {
	t := &Tree{}
	for _, v := range []int{5, 3, 4} {
		t.Insert(v)
	}

	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	t.leftRightRotation(t.root)
	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	// Output:
	// 5 3 4
	//  4 3 5
}

func ExampleRightLeftRotationAVLTree() {
	t := &Tree{}
	for _, v := range []int{1, 5, 4} {
		t.Insert(v)
	}

	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	t.rightLeftRotation(t.root)
	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	// Output:
	// 1 5 4
	//  4 1 5
}
