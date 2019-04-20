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

type Tree struct {
	Root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
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
func (t *Tree) LeftRotation(n *Node) {
	if n == nil || n.Right == nil {
		return
	}

	r := n.Right
	n.Right = r.Left
	r.Left = n

	if n == t.Root {
		t.Root = r
	}
}

func (t *Tree) RightRotation(n *Node) {
	if n == nil || n.Left == nil {
		return
	}

	l := n.Left
	n.Left = l.Right
	l.Right = n

	if n == t.Root {
		t.Root = l
	}
}

func (t *Tree) LeftRightRotation(node *Node) {
	node1 := node.Left.Right

	node.Left.Right = node1.Left
	node1.Left = node.Left

	node.Left = node1.Right
	node1.Right = node

	if node == t.Root {
		t.Root = node1
	}
}

func (t *Tree) RightLeftRotation(node *Node) {
	node1 := node.Right.Left

	node.Right.Left = node1.Right
	node1.Right = node.Right

	node.Right = node1.Left
	node1.Left = node

	if node == t.Root {
		t.Root = node1
	}
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
	if t.Root == nil {
		t.Root = &Node{Value: v}
	} else {
		t.Root.Insert(v)
	}
}

func (r *Node) Insert(v int) {
	if v < r.Value {
		if r.Left == nil {
			r.Left = &Node{Value: v}
		} else {
			r.Left.Insert(v)
		}
	} else {
		if r.Right == nil {
			r.Right = &Node{Value: v}
		} else {
			r.Right.Insert(v)
		}
	}
}

func (t *Tree) Preorder() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		t.Root.Preorder(ch)
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
