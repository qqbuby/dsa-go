package avl

import (
	"testing"
)

func TestNodeHeight(t *testing.T) {
	var tests = []struct {
		input *Node
		want  int
	}{
		{nil, -1},
		{&Node{Value: 0}, 0},
	}

	for _, test := range tests {
		h := test.input.Height()
		if h != test.want {
			t.Errorf("%v.Height(): %d, want: %d", test.input, h, test.want)
		}
	}
}
