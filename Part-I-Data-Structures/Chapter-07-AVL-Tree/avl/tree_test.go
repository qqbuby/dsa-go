package avl

import (
	"fmt"
)

func ExampleLeftRotationAndRightRotationAVLTree() {
	t := &Tree{}
	for _, v := range []int{14, 8, 24, 2, 11} {
		t.Insert(v)
	}

	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	t.RightRotation(t.Root)
	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	t.LeftRotation(t.Root)
	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	// Output:
	// 14 8 2 11 24
	//  8 2 14 11 24
	//  14 8 2 11 24
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

	t.LeftRightRotation(t.Root)
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

	t.RightLeftRotation(t.Root)
	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	// Output:
	// 1 5 4
	//  4 1 5
}
