package deque

import (
	"reflect"
	"testing"
)

func TestDeque(t *testing.T) {
	q := &Deque{}
	q.EnqueueBack(12)   // 12
	q.EnqueueFront(1)   // 1 12
	q.EnqueueBack(23)   // 1 12 23
	q.EnqueueFront(908) // 908 1 12 23

	a := []int{908, 1, 12, 23}
	b := []int{}
	for v := range q.Traverse() {
		b = append(b, v)
	}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("a %v is not equal to b %v", a, b)
	}

	q.DequeueFront() // 1 12 23
	q.DequeueBack()  // 1 12
	a = []int{1, 12}
	b = []int{}
	for v := range q.Traverse() {
		b = append(b, v)
	}
	if !reflect.DeepEqual(a, b) {
		t.Errorf("a %v is not equal to b %v", a, b)
	}

	if v, _ := q.PeekFront(); v != 1 {
		t.Errorf("q.PeekFront() is not equal to %d, got: %d", 1, v)
	}

	if v, _ := q.PeekBack(); v != 12 {
		t.Errorf("q.PeekFront() is not equal to %d, got: %d", 12, v)
	}
}
