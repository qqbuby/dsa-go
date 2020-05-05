package avl

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := &Stack{}

	s.Push(&Node{Value: 11})
	s.Push(&Node{Value: 15})

	if v, _ := s.Pop(); v.Value != 15 {
		t.Errorf("s.Pop() should be %d, got: %d", 15, v.Value)
	}

	if v, _ := s.Peek(); v.Value != 11 {
		t.Errorf("s.Peek() should be %d, got: %d", 11, v.Value)
	}

	if v, _ := s.Pop(); v.Value != 11 {
		t.Errorf("s.Pop() should be %d, got: %d", 11, v.Value)
	}

	if _, ok := s.Pop(); ok {
		t.Errorf("s.Pop() should be %t, got: %t", false, ok)
	}

	s.Push(&Node{Value: 1115})
	if _, ok := s.Pop(); !ok {
		t.Errorf("s.Pop() should be %t, got: %t", true, ok)
	}
}
