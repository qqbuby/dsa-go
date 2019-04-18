package singly

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	Value int
	Next  *Node
}

// 2.1.1 Insertion
//
// algorithm Add(value)
//   Pre: value is the value to add to the list
//   Post: value has been placed at the tail of the list
//   n = node(value)
//   if head = nil
//     head = n
// 	   tail = n
//   else
//     tail.Next = n
// 	tail = n
func (list *List) Add(value int) {
	n := &Node{Value: value}

	if list.head == nil {
		list.head = n
		list.tail = n
	} else {
		list.tail.Next = n
		list.tail = n
	}
}

// 2.1.2 Searching
//
// algorithm Contains(head, value)
//   Pre: head is the node in the list
//        value if the value to search for
//   Post: the item is either in the linked list, true; otherwise false
//   n = head
//   while n != nil && n.Value != value
//     n = n.Next
//  if n = nil
//    return false
//  else
//    return true
func (list *List) Contains(value int) bool {
	n := list.head
	for n != nil && n.Value != value {
		n = n.Next
	}

	return n != nil
}

// 2.1.3 Deletion
//
// 1. the lis is empty; or
// 2. the node to remove is the only node in the linked list; or
// 3. we are removing the head node; or
// 4. we are removing the tail node; or
// 5. the node to remove is somewhere in between the head and the tail; or
// 6. the item to remove doesn't exist in the linked list
//
// algorithm Remove(head, value)
//   Pre: head is the head node in the list
//   	  value is the value to remove from the list
//   Post: value is removed from the list, true; otherwise false
//   if head = nil
//     return false
//   n = head
//   if n.Value = value
//     if head = tail
//       head = nil
//       tail = nil
//     else
//       head = head.Next
//     return true
//   while n.Next != nil && n.Next.Value != value
//     n = n.Next
//   if n.Next != nil
//     if n.Next = tail
//       n.Next = nil
//       tail = n
//     else
//      n.Next = n.Next.Next
//     return true
//   return false
func (list *List) Remove(value int) bool {
	if list.head == nil {
		return false
	}

	n := list.head
	if n.Value == value {
		if list.head == list.tail {
			list.head = nil
			list.tail = nil
		} else {
			list.head = list.head.Next
		}

		return true
	}

	for n.Next != nil && n.Next.Value != value {
		n = n.Next
	}

	if n.Next != nil {
		if n.Next == list.tail {
			n.Next = nil
			list.tail = n
		} else {
			n.Next = n.Next.Next
		}

		return true
	}

	return false
}

// 2.1.4 Traversing the list
//
// algorithm Traverse(head)
//   Pre: head is the head node in the list
//   Post: the items in the list have been traversed
//   n = head
//   while n != nil
//     yield n.Value
func (list *List) Traverse() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		n := list.head
		for n != nil {
			ch <- n.Value
			n = n.Next
		}
	}()

	return ch
}

// 2.1.5 Traversing the list in reverse order
//
// algorithm ReverseTraverse(head, tail)
//   Pre: head and tail belong to the same list
//   Post: the items in the list have been traversed in reverse order
//   if tail != nil
//     curr = tail
//     while curr != head
//       prev = head
//       while pre.Next != curr
//         prev = prev.Next
//       yield curr.Value
//       curr = prev
//     yield curr.Value
func (list *List) ReverseTraverse() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		if list.tail != nil {
			curr := list.tail
			for curr != list.head {
				prev := list.head
				for prev.Next != curr {
					prev = prev.Next
				}
				ch <- curr.Value
				curr = prev
			}

			ch <- curr.Value
		}
	}()

	return ch
}
