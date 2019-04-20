// An ordered set is similar to an unordered set
// in the sense that its members are distinct,
// but an ordered set enforces some predefined
// comparision on each of its members to produce
// a set whose members are ordered appropriately.
//
// The ordered set has its order realised by performing
// an inorder traversal upon its backing ree data structure
// which yields the correct ordered sequence of set members.
package set

type Set struct {
	root *node
}

type node struct {
	Value int
	Left  *node
	Right *node
}

func (s *Set) Contains(v int) bool {
	r := s.root
	for r != nil {
		if r.Value == v {
			return true
		} else if v < r.Value {
			r = r.Left
		} else {
			r = r.Right
		}
	}

	return false
}

func (s *Set) Add(v int) {
	if s.root == nil {
		s.root = &node{Value: v}
	} else {
		s.root.add(v)
	}
}

func (r *node) add(v int) {
	if v < r.Value {
		if r.Left == nil {
			r.Left = &node{Value: v}
		} else {
			r.Left.add(v)
		}
	} else if v > r.Value {
		if r.Right == nil {
			r.Right = &node{Value: v}
		} else {
			r.Right.add(v)
		}
	}
}

func (s *Set) Union(v *Set) *Set {
	u := &Set{}

	for a := range s.Traverse() {
		u.Add(a)
	}

	for a := range v.Traverse() {
		u.Add(a)
	}

	return u
}

func (s *Set) Intersection(v *Set) *Set {
	i := &Set{}

	for a := range s.Traverse() {
		if v.Contains(a) {
			i.Add(a)
		}
	}

	return i
}

func (s *Set) Traverse() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		s.root.inorder(ch)
	}()

	return ch
}

func (r *node) inorder(ch chan<- int) {
	if r != nil {
		ch <- r.Value
		r.Left.inorder(ch)
		r.Right.inorder(ch)
	}
}
