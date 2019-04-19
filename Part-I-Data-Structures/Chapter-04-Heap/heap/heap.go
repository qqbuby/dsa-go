// A heap can be thought of as a simple tree data structure,
// however a heap usually employs one of the follow strategies:
//   1. min heap; or
//   2. max heap
// If you were to choose the min heap strategy when each parent
// node would have a value that is <= than its children. The
// opposite is true for the max heap strategy.
//
// Unlike other tree data structures, a heap is generally implemented
// as an array rather than a series of nodes which each have references
// to other nodes.
//
// Because we are using an array we ned some way to calculate the index
// of a parent node, and the children of a node. The requried expressions
// for this are defined as follows for a node at index:
//   1. (index - 1) / 2 (parent index)
//   2. 2 * (index + 1) (left child)
//   3. 2 * (index + 2) (right child)
package heap

type Heap struct {
	heap []int
}

func (h *Heap) Count() int {
	return len(h.heap)
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return 2 * (index + 1)
}

func right(index int) int {
	return 2 * (index + 2)
}

// 4.1 Insertion
//
// algorithm Add(value)
//   Pre: value is the value to add to the heap
//        Count is the number of items in the heap
//   Post: the value has been added to the heap
//   head[Count] = value
//   Count = Count + 1
//   MinHeapfiy()
//
// algorithm MinHeapify()
//   Pre: Count is the number of items in the heap
//        heap is the array used to store the heap items
//   Post: the heap has preserved min heap ordering
//   i = Count - 1
//   while i > 0 and heap[i] < heap[(i - 1) / 2]
//     Swap(heap[i], heap[(i -1 ) / 2]
//     i = (i - 1) / 2
func (h *Heap) Add(value int) {
	h.heap = append(h.heap, value)

	// MinHeapify
	i := h.Count() - 1
	p := parent(i)
	for i > 0 && h.heap[i] < h.heap[p] {
		h.heap[i], h.heap[p] = h.heap[p], h.heap[i]
		i = parent(i)
		p = parent(i)
	}
}

// 4.2 Deletion
//
// The algorithm for deletion has tree steps:
//   1. find the index of the value to delete
//   2. put the last value in the heap at the index location of the item to delete
//   3. verify heap ordering for each subtree which used to include the value
//
// algorithm Remove(value)
//   Pre: value is the value to remove fromt he heap
//        left, and right are updated alias' for 2 * index + 1, and 2 * index + 2 repsectively
//        Count is the number items in the heap
//        heap is the array used to store the heap items
//   Post: value is located in the heap and removed, true; otherwise false
//   // Step 1
//   index = FindIndex(heap, value)
//   if index < 0
//     return false
//   Count = Count - 1
//   // Step 2
//   heap[index] = heap[Count]
//   // Step 3
//   while left < Count and heap[index] > heap[left] or heap[index] > heap[right]
//     // promote smallest key from subtree
//     if heap[left] < heap[right]
//       Swap(heap[index], heap[left]
//       index = left
//     else
//       Swap(heap[index], heap[right]
//       index = right
func (h *Heap) Remove(value int) bool {
	index := h.FindIndex(value)
	if index < 0 {
		return false
	}

	h.heap[index] = h.heap[h.Count()-1]
	h.heap = h.heap[:len(h.heap)-1] // Count = Count - 1

	left := left(index)
	right := right(index)

	for index < h.Count() &&
		(left < h.Count() && h.heap[left] < h.heap[index]) ||
		(right < h.Count() && h.heap[right] < h.heap[index]) {
		if left < h.Count() && right < h.Count() {
			if h.heap[left] < h.heap[right] {
				h.heap[left], h.heap[index] = h.heap[index], h.heap[left]
				index = left
			} else {
				h.heap[right], h.heap[index] = h.heap[index], h.heap[right]
				index = right
			}
		} else if left < h.Count() {
			h.heap[left], h.heap[index] = h.heap[index], h.heap[left]
			index = left
		} else {
			h.heap[right], h.heap[index] = h.heap[index], h.heap[right]
			index = right
		}
	}

	return true
}

func (h *Heap) FindIndex(value int) int {
	for i := 0; i < h.Count(); i++ {
		if h.heap[i] == value {
			return i
		}
	}

	return -1
}

// 4.3 Searching
//
// algorithm Contains(value)
//   Pre: value is the value to search the heap for
//        Count is the number of items in the heap
//        heap is the array used to store the heap items
//   Post: value is located in the heap, in which cae true; otherwise false
//   i = 0
//   while i < Count and heap[i] != value
//     i = i + 1
//   if i < Count
//     return true
//   else
//     return false
func (h *Heap) Contains(value int) bool {
	for i := 0; i < h.Count(); i++ {
		if h.heap[i] == value {
			return true
		}
	}

	return false
}

// 4.4 Traversal
//
func (h *Heap) Traverse() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < h.Count(); i++ {
			ch <- h.heap[i]
		}
	}()

	return ch
}
