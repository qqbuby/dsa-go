package avl

type Stack struct {
	head *node
}

type node struct {
	Value *Node
	Next  *node
}

func (s *Stack) Push(v *Node) {
	n := &node{Value: v}

	if s.head == nil {
		s.head = n
	} else {
		n.Next = s.head
		s.head = n
	}
}

func (s *Stack) Pop() (*Node, bool) {
	if s.head == nil {
		return nil, false
	}

	v := s.head.Value
	s.head = s.head.Next

	return v, true
}

func (s *Stack) Peek() (*Node, bool) {
	if s.head == nil {
		return nil, false
	}

	return s.head.Value, true
}
