package numeric

import (
	"fmt"
	"strings"
	"testing"
)

// 9.2 Base conversions
//
// algorithm ToBinary(n)
//   Pre: n >= 0
//   Post: n has been converted into its base 2 representation
//   while n > 0
//     list.Add(n % 2)
//     n = n / 2
//   return reverse(list)
func ToBinary(n int) string {
	s := []int{}

	for n > 0 {
		s = append(s, n%2)
		n = n / 2
	}

	b := []int{}
	for i := len(s) - 1; i >= 0; i-- {
		b = append(b, s[i])
	}

	bb := strings.Trim(strings.Join(strings.Split(fmt.Sprint(b), " "), ""), "[]")

	return bb
}

func TestToBinary(t *testing.T) {
	n := 742

	b := ToBinary(n)

	w := "1011100110"

	if b != w {
		t.Errorf("ToBinary(%d) want: %v, got: %v", n, w, b)
	}
}
