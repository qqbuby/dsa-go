package doubly

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

// 2.2.1 Insertion
//
// algorithm Add(value)
//   Pre: value is the value to add to the list
//   Post: value had been placed at the tail of the list
//   n = node(value)
//   if head = nil
//     head = n
//     tail = n
//   else
//     n.Previous = tail
//     tail.Next = n
//     tail = n
func (list *List) Add(value int) {
	n := &Node{Value: value}

	if list.head == nil {
		list.head = n
		list.tail = n
	} else {
		n.Prev = list.tail
		list.tail.Next = n
		list.tail = n
	}
}

// 2.2.2 Deletion
//
// algorithm Remove(head, value)
//   Pre: head is the head node of the list
//        value is the value to remove from the list
//   Post: value is removed from the list, true; otherwise false
//   if head = nil
//     return false
//   if head.Value = value
//     if head = tail
//       head = nil
//       tail = nil
//     else
//       head = head.Next
//       head.Next.Previous = nil
//     return true
//   n = head.Next
//   while n != nil && n.Value != value
//     n = n.Next
//   if n != nil
//     if n == tail
//       tail = tail.Previous
//       tail.Next = nil
//       return true
//     else
//       n.Previous.Next = n.Next
//       n.Next.Previous = n.Previous
//       return true
//   return false
func (list *List) Remove(value int) bool {
	if list.head == nil {
		return false
	}

	if list.head.Value == value {
		if list.head == list.tail {
			list.head = nil
			list.tail = nil
		} else {
			list.head = list.head.Next
			list.head.Prev = nil
		}
	}

	n := list.head.Next
	for n != nil && n.Value != value {
		n = n.Next
	}

	if n != nil {
		if n == list.tail {
			list.tail = list.tail.Prev
			list.tail.Next = nil
		} else {
			n.Prev.Next = n.Next
			n.Next.Prev = n.Prev
		}

		return true
	}

	return false
}

// 2.1.3 ReverseTraverse
//
// algorithm ReverseTraverse(tail)
//   Pre: tail is the tail node of the list
//   Post: the list has been traversed in reverse order
//   n = tail
//   while n != nil
//     yield n.Value
//     n = n.Previous
func (list *List) ReverseTraverse() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		n := list.tail
		for n != nil {
			ch <- n.Value
			n = n.Prev
		}
	}()

	return ch
}
