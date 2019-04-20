package set

import (
	"fmt"
)

func ExampleOrderedSet() {
	a := &Set{}
	for _, v := range []int{1, 2, 3} {
		a.Add(v)
	}

	b := &Set{}
	for _, v := range []int{6, 2, 9} {
		b.Add(v)
	}

	u := a.Union(b)
	for v := range u.Traverse() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	i := a.Intersection(b)
	for v := range i.Traverse() {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	// Output:
	// 1 2 3 6 9
	//  2
}
