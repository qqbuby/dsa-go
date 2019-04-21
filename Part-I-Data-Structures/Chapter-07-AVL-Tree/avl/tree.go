// In the early 60's G.M. Adelson-Velsky and E.M. Landis
// invented the first self-balancing binary search tree
// data structure, calling it AVL Tree.
//
// An AVL tree is a binary search tree between the height
// of the left and right subtrees cannot be no more than one.
//
// The AVL balance condition, known also as the node balance
// factor represents and additional piece of information stored
// for each node. The balance factor is defined as an integer in
// range of [-1, 1], where -1 will be an extra node on the right
// subtree, and 1 is an extral node on the left subtree, accounting
// for an odd number of nodes.
package avl

import (
	_ "fmt"
)

type Tree struct {
	root *Node
}

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

// 7.1 Tree Rotations
//
// algorithm LeftRotation(node)
//   Pre: node.Right != nil
//   Post: node.Right is the new root of the subtree,
//         node has become node.Right's left child and,
//         BST propeties are preserved
//   RightNode = node.Right
//   node.Right = RightNode.Left
//   RightNode.Left = node
//
// algorithm RightRotation(node)
//   Pre: node.Left != nil
//   Post: node.Left is the new root of the subtree,
//         node has become node.Left's right child and,
//         BST properties preserved
//   LeftNode = node.Left
//   node.Left = LeftNode.Right
//   LeftNode.Right = node

//  A                B
//   \              / \
//    B     =>     A   C
//   / \            \
//  X   C            X
func (t *Tree) leftRotation(node *Node) {
	if node == nil || node.Right == nil {
		return
	}

	pivot := node.Right
	node.Right = pivot.Left

	if pivot.Left != nil {
		pivot.Left.Parent = node
	}

	pivot.Left = node

	if node.Parent != nil {
		if node.Parent.Right == node {
			node.Parent.Right = pivot
		} else {
			node.Parent.Left = pivot
		}
	} else {
		t.root = pivot
	}

	pivot.Parent = node.Parent
	node.Parent = pivot
}

//      C            B
//     /            / \
//    B      =>    A   C
//   / \              /
//  A   X            X
func (t *Tree) rightRotation(node *Node) {
	if node == nil || node.Left == nil {
		return
	}

	pivot := node.Left

	node.Left = pivot.Right

	if pivot.Right != nil {
		pivot.Right.Parent = node
	}

	pivot.Right = node

	if node.Parent != nil {
		if node.Parent.Left == node {
			node.Parent.Left = pivot
		} else {
			node.Parent.Right = pivot
		}
	} else {
		t.root = pivot
	}

	pivot.Parent = node.Parent
	node.Parent = pivot
}

//    C               C               B
//   /               /               / \
//  A       =>      B        =>     A   C
//   \             / \               \ /
//    B           A   X               X
//   / \           \
//  X   X           X
func (t *Tree) leftRightRotation(node *Node) {
	if node == nil ||
		node.Left == nil ||
		node.Left.Right == nil {
		return
	}

	pivot := node.Left.Right // B

	// left rotation
	node.Left.Right = pivot.Left

	if pivot.Left != nil {
		pivot.Left.Parent = node.Left
	}

	pivot.Left = node.Left
	node.Left.Parent = pivot

	// right rotation
	node.Left = pivot.Right

	if pivot.Right != nil {
		pivot.Right.Parent = node
	}

	pivot.Right = node

	if node.Parent != nil {
		if node.Parent.Right == node {
			node.Parent.Right = pivot
		} else {
			node.Parent.Left = pivot
		}
	} else {
		t.root = pivot
	}

	pivot.Parent = node.Parent
	node.Parent = pivot
}

//   A               A                 B
//    \               \               / \
//     C      =>       B      =>     A   C
//    /               / \             \ /
//   B               X   C             X
//  / \                 /
// X   X               X
func (t *Tree) rightLeftRotation(node *Node) {
	if node == nil ||
		node.Right == nil ||
		node.Right.Left == nil {
		return
	}

	pivot := node.Right.Left

	// right rotation
	node.Right.Left = pivot.Right

	if pivot.Right != nil {
		pivot.Right.Parent = node.Right
	}

	pivot.Right = node.Right
	node.Right.Parent = pivot

	// left rotation
	node.Right = pivot.Left

	if pivot.Left != nil {
		pivot.Left.Parent = node
	}

	pivot.Left = node

	if node.Parent != nil {
		if node.Parent.Left == node {
			node.Parent.Left = pivot
		} else {
			node.Parent.Right = pivot
		}
	} else {
		t.root = pivot
	}

	pivot.Parent = node.Parent
	node.Parent = pivot
}

func (n *Node) Height() int {
	if n == nil {
		return -1
	}

	return max(n.Left.Height(), n.Right.Height()) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// 7.2 Tree Rebalancing
//
// algorithm CheckBalance(current)
//   Pre: current is the node to start from balancing
//   Post: current height has been updated while tree balance is if needed
//         restored through rotations
//   if current.Left = nil and current.Right = nil
//     current.Height = -1
//   else
//     current.Height = Max(Height(current.Left), Height(current.Right)) + 1
//   if Height(current.Left) - Height(current.Right) > 1
//     if Height(current.Left.Left) - Height(current.Left.Right) > 0
//       RightRotation(current)
//     else
//       LeftAndRightRotation(current)
//   else if Height(current.Left) - Height(current.Right) < -1
//     if Height(current.Right.Left) - Height(current.Right.Right) < 0
//       LeftRotation(current)
//     else
//       RightAndLeftRotation(current)
func (t *Tree) CheckBalance(current *Node) {
}

// 7.3 Insertion
//
// algorithm Insert(value)
//   Pre: value has passed custom type checks for type T
//   Post: vlaue has been placed in the correct location in the tree
//   if root = nil
//     root = node(value)
//   else
//     InsertNode(root, value)
//
// algorithm InsertNode(current, value)
//   Pre: current is the node to start from
//   Post: value has been placed in the correct location in the tree while
//         preserving tree balance
//   if value < current.Value
//     if current.Left = nil
//       current.Left = node(value)
//     else
//       InsertNode(current.Left, value)
//   else
//     if current.Right = nil
//       current.Right = node(value)
//     else
//       InsertNode(current.Right, value)
//   CheckBalance(current)
func (t *Tree) Insert(v int) {
	if t.root == nil {
		t.root = &Node{Value: v}
	} else {
		t.root.Insert(v)
	}
}

func (r *Node) Insert(v int) {
	if v < r.Value {
		if r.Left == nil {
			r.Left = &Node{Value: v}
			r.Left.Parent = r
		} else {
			r.Left.Insert(v)
		}
	} else {
		if r.Right == nil {
			r.Right = &Node{Value: v}
		} else {
			r.Right.Insert(v)
			r.Right.Parent = r
		}
	}
}

func (t *Tree) Preorder() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		t.root.Preorder(ch)
	}()

	return ch
}

func (r *Node) Preorder(ch chan<- int) {
	if r != nil {
		ch <- r.Value
		r.Left.Preorder(ch)
		r.Right.Preorder(ch)
	}
}
