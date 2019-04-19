package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	a := &Set{}
	for _, v := range []int{1, 2, 3} {
		a.Add(v)
	}

	b := &Set{}
	for _, v := range []int{6, 2, 9} {
		b.Add(v)
	}

	u := a.Union(b)
	var tests = []struct {
		input int
		want  bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{6, true},
		{9, true},
	}
	for _, test := range tests {
		if got := u.Contains(test.input); got != test.want {
			t.Errorf("u.Contains(%d) = %t", test.input, got)
		}
	}

	i := a.Intersection(b)
	tests = []struct {
		input int
		want  bool
	}{
		{1, false},
		{2, true},
		{3, false},
		{6, false},
		{9, false},
	}
	for _, test := range tests {
		if got := i.Contains(test.input); got != test.want {
			t.Errorf("u.Contains(%d) = %t", test.input, got)
		}
	}
}
