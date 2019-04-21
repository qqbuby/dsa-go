package avl

import (
	"fmt"
)

//func ExampleLeftRotationAndRightRotationAVLTree() {
//	t := &Tree{}
//	for _, v := range []int{14, 8, 24, 2, 11, 1} {
//		t.Add(v)
//	}
//
//	for v := range t.Preorder() {
//		fmt.Printf(" %d", v)
//	}
//	fmt.Println()
//
//	t.rightRotation(t.root.Left)
//	for v := range t.Preorder() {
//		fmt.Printf(" %d", v)
//	}
//	fmt.Println()
//
//	t.leftRotation(t.root.Left)
//	for v := range t.Preorder() {
//		fmt.Printf(" %d", v)
//	}
//	fmt.Println()
//
//	// Output:
//	// 14 8 2 1 11 24
//	//  14 2 1 8 11 24
//	//  14 8 2 1 11 24
//}
//
//func ExampleLeftRightRotationAVLTree() {
//	t := &Tree{}
//	for _, v := range []int{5, 3, 4} {
//		t.Add(v)
//	}
//
//	for v := range t.Preorder() {
//		fmt.Printf(" %d", v)
//	}
//	fmt.Println()
//
//	t.leftRightRotation(t.root)
//	for v := range t.Preorder() {
//		fmt.Printf(" %d", v)
//	}
//	fmt.Println()
//
//	// Output:
//	// 5 3 4
//	//  4 3 5
//}
//
//func ExampleRightLeftRotationAVLTree() {
//	t := &Tree{}
//	for _, v := range []int{1, 5, 4} {
//		t.Add(v)
//	}
//
//	for v := range t.Preorder() {
//		fmt.Printf(" %d", v)
//	}
//	fmt.Println()
//
//	t.rightLeftRotation(t.root)
//	for v := range t.Preorder() {
//		fmt.Printf(" %d", v)
//	}
//	fmt.Println()
//
//	// Output:
//	// 1 5 4
//	//  4 1 5
//}

func ExampleLeftCheckBalance() {
	t := &Tree{}
	for _, v := range []int{1, 2, 3} {
		t.Add(v)
	}

	t.CheckBalance(t.root)

	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	// Output:
	// 2 1 3
}

func ExampleRightCheckBalance() {
	t := &Tree{}
	for _, v := range []int{3, 2, 1} {
		t.Add(v)
	}

	t.CheckBalance(t.root)

	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	// Output:
	// 2 1 3
}

func ExampleLeftRightCheckBalance() {
	t := &Tree{}
	for _, v := range []int{3, 1, 2} {
		t.Add(v)
	}

	t.CheckBalance(t.root)

	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	// Output:
	// 2 1 3
}

//func ExampleRightLeftCheckBalance() {
//	t := &Tree{}
//	for _, v := range []int{1, 3, 2} {
//		t.Add(v)
//	}
//
//	t.CheckBalance(t.root)
//
//	for v := range t.Preorder() {
//		fmt.Printf(" %d", v)
//	}
//	fmt.Println()
//
//	// Output:
//	// 2 1 3
//}

func ExampleAVLTree() {
	t := &Tree{}

	for _, v := range []int{2, 8, 11, 15, 14, 24} {
		t.Add(v)
	}

	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	for v := range t.Inorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	t.Remove(24)
	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	t.Remove(8)
	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	t.Remove(14)
	t.Remove(2)
	for v := range t.Preorder() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	// Output:
	// 14 8 2 11 15 24
	//  2 8 11 14 15 24
	//  14 8 2 11 15
	//  14 2 11 15
	//  11 15
}
