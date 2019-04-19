// A set contains a number of vlaues, in no particular order.
// The values within the set are distinct from one another.
package set

type Set struct {
	m map[int]bool
}

func (s *Set) Count() int {
	return len(s.m)
}

func (s *Set) Add(v int) {
	if s.m == nil {
		s.m = make(map[int]bool)
	}

	s.m[v] = true
}

func (s *Set) Remove(v int) {
	delete(s.m, v)
}

func (s *Set) Contains(v int) bool {
	return s.m[v]
}

func (s *Set) Traverse() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for k, _ := range s.m {
			ch <- k
		}
	}()

	return ch
}

// algorithm Union(set1, set2)
//   Pre: set1, and set2 != nil
//        union is a set
//   Post: A union of set1, and set2 has been created
//   foreach item in set1
//     union.Add(item)
//   foreach item in set2
//    union.Add(item)
//   return union
func (s1 *Set) Union(s *Set) *Set {
	union := &Set{}

	if s != nil {
		for k, _ := range s.m {
			union.Add(k)
		}
	}

	if s1 != nil {
		for k, _ := range s1.m {
			union.Add(k)
		}
	}

	return union
}

// algorithm Intersection(set1, set2)
//   Pre: set1, and set2 != nil
//        intersection, and smallerSet are sets
//   Post: An intersection of set1, and set2 has been created
//   if set1.Count < set2.Count
//     smallerSet = set1
//   else
//     smallerSet = set2
//   foreach item in smallerSet
//     if set1.Contains(item) && set2.Contains(item)
//       intersection.Add(item)
func (s1 *Set) Intersection(s *Set) *Set {
	intersection := &Set{}
	smaller := s

	if s1.Count() < s.Count() {
		smaller = s1
	}

	for k, _ := range smaller.m {
		if s1.Contains(k) && s.Contains(k) {
			intersection.Add(k)
		}
	}

	return intersection
}
