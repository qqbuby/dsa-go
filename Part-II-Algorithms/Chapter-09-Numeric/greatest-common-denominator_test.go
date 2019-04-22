package numeric

import (
	"fmt"
)

// 9.3 Attaining the greatest common denominator of two numbers
//
// algorithm GreatestCommonDenominator(m, n)
//   Pre: m and n are integers
//   Post: the grestest common denominator of the two integers is calculated
//   if n = 0
//     return m
//   return GreatestCommonDenominator(m, m % n)
func GreatestCommonDenominator(m, n int) int {
	if n == 0 {
		return m
	}

	if m < n {
		m, n = n, m
	}

	return GreatestCommonDenominator(n, m%n)
}

func ExampleGreatestCommonDenominator() {
	g := GreatestCommonDenominator(9, 15)
	fmt.Println(g)

	g = GreatestCommonDenominator(15, 9)
	fmt.Println(g)

	// Output:
	// 3
	// 3
}
