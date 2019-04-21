package avl

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func (n *Node) Height() int {
	if n == nil {
		return -1
	}

	return max(n.Left.Height(), n.Right.Height()) + 1
}

func (n *Node) Add(v int) {
	if v < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: v}
			n.Left.Parent = n
		} else {
			n.Left.Add(v)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: v}
		} else {
			n.Right.Add(v)
			n.Right.Parent = n
		}
	}
}

func (n *Node) Preorder(ch chan<- int) {
	if n != nil {
		ch <- n.Value
		n.Left.Preorder(ch)
		n.Right.Preorder(ch)
	}
}
