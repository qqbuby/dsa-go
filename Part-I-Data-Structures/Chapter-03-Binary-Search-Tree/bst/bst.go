// Binary search trees (BSTs) are very simple to understand. We start with a root
// node with value x, where the left subtree of x contains nodes with values < x
// and the right subtree contains nodes whose values are >= x. Each node follows
// the same rules with respect to nodes in their left and right subtrees.
//
// BSTs are of interest because they have operations which are favourably fast:
// insertion, look up, and deletion can all be done in O(log n) time. It is important
// to note that the O(log n) times for these operations can only be attained if
// the BST is reasonably balanced.
package bst

type Tree struct {
	Root  *Node
	Count int
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// 3.1 Insertion
//
// algorithm Insert(value)
//   Pre: value has passed custom type checks for type T
//   Post: value has been placed in the correct location in the tree
//   if root = nil
//     root = node(value)
//   else
//     InsertNode(root, value)
//
// algorithm InsertNode(current, value)
//   Pre: current is the node to start from
//   Post: value has been placed in the correct location in the tree
//   if value < current.Value
//     if current.Left != nil
//       InsertNode(current.Left, value)
//     else
//       current.Left = node(value)
//   else
//     if current.Right != nil
//       InsertNode(current.Right, value)
//     else
//       current.Right = node(value)
func (t *Tree) Insert(value int) {
	if t.Root == nil {
		t.Root = &Node{Value: value}
	} else {
		insertNode(t.Root, value)
	}

	t.Count++
}

func insertNode(curr *Node, value int) {
	if value < curr.Value {
		if curr.Left != nil {
			insertNode(curr.Left, value)
		} else {
			curr.Left = &Node{Value: value}
		}
	} else {
		if curr.Right != nil {
			insertNode(curr.Right, value)
		} else {
			curr.Right = &Node{Value: value}
		}
	}
}

// 3.2 Searching
//
// 1. the root = nil in which case value is not in the BST; or
// 2. root.Value = value in which case value is in the BST; or
// 3. value < root.Value, we must inspect the left subtree of root for value; or
// 3. value > root.Value, we must inspect the right subtree of root for value.
//
// alogrithm Contains(root, value)
//   Pre: root is the root node of the tree, value is what we would like to locate
//   Post: value is either located or not
//   if root = nil
//     return false
//   else
//     if root.Value = value
//       return true
//     else if value < root.Value
//       return Contains(root.Left, value)
//     else
//       return Contains(root.Right, value)
func Contains(root *Node, value int) bool {
	if root == nil {
		return false
	}

	if root.Value == value {
		return true
	} else if value < root.Value {
		return Contains(root.Left, value)
	} else {
		return Contains(root.Right, value)
	}
}

func (t *Tree) Contains(value int) bool {
	return Contains(t.Root, value)
}

// 3.3 Deletion
//
// 1. the value to remove is a leaf node; or
// 2. the value to remove has a right subtree, but no left subtree; or
// 3. the value to remove has a left subtree, but no right subtree; or
// 4. the value to remove has both a left and right subtree in which case we
// promote the largest value in the left subtree (or promote the least value in the right subtree ?)
//
// There is also an implicit fifth case whereby the node to be removed is the
// only node in the tree. This case is already covered by the ¯rst, but should be
// noted as a possibility nonetheless.
//
// Of course in a BST a value may occur more than once. In such a case the
// first occurrence of that value in the BST will be removed.
//
// algorithm Remove(value)
//   Pre: value is the value of the node to remove, root is the root node of the BST
//        Count is the number of items of the BST
//   Post: node iwht value is removed if found in which case yields true, otherwise false
//   nodeToRemove = FindNode(value)
//   if nodeToRemove = nil
//     return false
//   parent = FindParent(value)
//   if Count = 1
//     root = nil
//   else if nodeToRemove.Left = nil and nodeToRemove.Right = nil
//     if nodeToRemove.Value < parent.Value
//       parent.Left = nil
//     else
//       parent.Right = nil
//   else if nodetoRemove.Left = nil and nodeToRemove.Right != nil
//     if nodeToRemove.value < parent.Value
//       parent.Left = nodeToRemove.Right
//     else
//       parent.Right = nodeToRemove.Right
//   else if nodetoRemove.Left != nil and nodeToRemove.Right = nil
//     if nodeToRemove.value < parent.Value
//       parent.Left = nodeToRemove.Left
//     else
//       parent.Right = nodeToRemove.Left
//   else
//     largestValue = nodeToRemove.Left
//     while largestValue.Right != nil
//       // find the largest value in the left subtree of nodeToRemove
//       largestValue = largetValue.Right
//     FindParent(largetValue.Value).Right = nil
//     nodeToRemove.Value = largestValue.Value
//   Count = Count - 1
//   return true
func (t *Tree) Remove(value int) bool {
	nodeToRemove := t.FindNode(value)
	if nodeToRemove == nil {
		return false
	}

	if t.Count == 1 {
		t.Root = nil
		return true
	}

	parent := t.FindParent(value)

	if nodeToRemove.Left == nil && nodeToRemove.Right == nil {
		if nodeToRemove.Value < parent.Value {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	} else if nodeToRemove.Left == nil && nodeToRemove.Right != nil {
		if nodeToRemove.Value < parent.Value {
			parent.Left = nodeToRemove.Right
		} else {
			parent.Right = nodeToRemove.Right
		}
	} else if nodeToRemove.Left != nil && nodeToRemove.Right == nil {
		if nodeToRemove.Value < parent.Value {
			parent.Left = nodeToRemove.Left
		} else {
			parent.Right = nodeToRemove.Left
		}
	} else {
		largest := nodeToRemove.Left
		for largest.Right != nil {
			largest = largest.Right
		}
		t.FindParent(largest.Value).Right = nil
		nodeToRemove.Value = largest.Value
	}

	t.Count--

	return true
}

// 3.4 Find the parent of a give node
//
// algorithm FindParent(value, root)
//   Pre: value is the value of the node we want to find the parent of
//		  root is the root node of the BST and is != nil
//   Post: a reference to the parent node of value if found; otherwise nil
//   if value == root.Value
//     return nil
//   if value < root.Value
//     if root.Left = nil
// 	     return nil
//     else if root.Left.Value = value
//       return root
//     else
//       return FindParent(value, root.Left)
//   else
//     if root.Right = nil
//       return nil
//     else if root.Right.Value = value
//       return root
//     else
//       return FindParent(value, root.Right)
func (t *Tree) FindParent(value int) *Node {
	if t.Root == nil || t.Root.Value == value {
		return nil
	} else {
		return findParent(t.Root, value)
	}
}

func findParent(root *Node, value int) *Node {
	if value < root.Value {
		if root.Left != nil && root.Left.Value == value {
			return root
		} else {
			return findParent(root.Left, value)
		}
	} else {
		if root.Right != nil && root.Right.Value == value {
			return root
		} else {
			return findParent(root.Right, value)
		}
	}
}

// 3.5 Attaning a reference to a node
//
// algorithm FindNode(root, value)
//   Pre: value is the value of the node we want to find the parent of
//        root is the root node of the BST
//   Post: a reference to the node of value if found; otherwise nil
//   if root = nil
//     return nil
//   if root.Value = value
//     return root
//   else if value < root.Value
//     FindNode(root.Left, value)
//   else
//     FindNode(root.Right, value)
func (t *Tree) FindNode(value int) *Node {
	return findNode(t.Root, value)
}

func findNode(root *Node, value int) *Node {
	if root == nil {
		return nil
	} else if root.Value == value {
		return root
	} else if value < root.Value {
		return findNode(root.Left, value)
	} else {
		return findNode(root.Right, value)
	}
}

// 3.6 Find the smallest and largest values in the binary search tree
//
// algorithm FindMin(root)
//   Pre: root is the root node of the BST
//        root != nil
// 	 Post: the smallest value in the BST
func (t *Tree) FindMin() *Node {
	if t.Root == nil {
		return nil
	}

	min := t.Root
	for min.Left != nil {
		min = min.Left
	}

	return min
}

// algorithm FindMax(root)
//   Pre: root is the root node of the BST
//        root != nil
//   Post: the largest value in the BST
func (t *Tree) FindMax() *Node {
	if t.Root == nil {
		return nil
	}

	max := t.Root
	for max.Right != nil {
		max = max.Right
	}

	return max
}

// 3.7.1 Preorder Tree Traversal
//
// algorithm Preorder(root)
//   Pre: root is the root node of the BST
//   Post: the nodes in the BST have been visited in preorder
//   if root != nil
//     yield root.Value
//     Preorder(root.Left)
//     Preorder(root.Right)
func (t *Tree) Preorder() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		preorder(t.Root, ch)
	}()

	return ch
}

func preorder(root *Node, ch chan<- int) {
	if root != nil {
		ch <- root.Value
		preorder(root.Left, ch)
		preorder(root.Right, ch)
	}
}

// 3.7.1 Postorder
//
// algorithm Postorder(root)
//   Pre: root is the root node of the BST
//   Post: the nodes in the BST have been visited in postorder
//   if root != nil
//     Postorder(root.Left)
//     Postorder(root.Right)
//     yield root.Value
func (t *Tree) Postorder() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		postorder(t.Root, ch)
	}()

	return ch
}

func postorder(root *Node, ch chan<- int) {
	if root != nil {
		postorder(root.Left, ch)
		postorder(root.Right, ch)
		ch <- root.Value
	}
}

// 3.7.3 Inorder
//
// algorithm Inorder(root)
//   Pre: root is the root node of the BST
//   Post: the nodes in the BST have been visited in iorder
//   if root != nil
//     Inorder(root.Left)
//     yield root.Value
//     Inorder(root.Right)
func (t *Tree) Inorder() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		inorder(t.Root, ch)
	}()

	return ch
}

func inorder(root *Node, ch chan<- int) {
	if root != nil {
		inorder(root.Left, ch)
		ch <- root.Value
		inorder(root.Right, ch)
	}
}

// 3.7.4 Breadth First
// Traversing a tree in breath first order yields the values of all nodes of a particular depth in the tree before any deeper ones.
//
// algorithm BreathFirst(root)
//   Pre: root is the root node of the BST
//   Post: the nodes in the BST have been visited in breath first order
//   q = queue()
//   while root != nil
//     yield root.Value
//     if root.Left != nil
//       q.Enqueue(root.Left)
//     if root.Right != nil
//       q.Enqueue(root.Right)
//     if !q.IsEmpty()
//       root = q.Dequeue()
//     else
//       root = nil
func (t *Tree) BreathFirst() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		q := &Queue{}
		root := t.Root
		for root != nil {
			ch <- root.Value

			if root.Left != nil {
				q.Enqueue(root.Left)
			}

			if root.Right != nil {
				q.Enqueue(root.Right)
			}

			root, _ = q.Dequeue()
		}
	}()

	return ch
}
