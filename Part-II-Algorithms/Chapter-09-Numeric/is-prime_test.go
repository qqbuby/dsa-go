package numeric

import (
	"math"
	"testing"
)

// 9.1 Primality Test
//
// algorithm IsPrime(n)
//   Post: n is determined to be prime or not
//   for i = 2 to n
//     for j = 1 to sqrt(n)
//       if i*j = n
//         return false
func IsPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i < n; i++ {
		for j := 1; j <= sqrt; j++ {
			if i*j == n {
				return false
			}
		}
	}

	return true
}

func TestIsPrime(t *testing.T) {
	var tests = []struct {
		test int
		want bool
	}{
		{1, true},
		{2, true},
		{4, false},
		{7, true},
		{9, false},
	}

	for _, v := range tests {
		if got := IsPrime(v.test); got != v.want {
			t.Errorf("IsPrime(%d) got: %t, want: %t", v.test, got, v.want)
		}
	}
}
