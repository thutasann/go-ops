package slicedatatsructure

import "fmt"

type Set[T comparable] struct {
	data []T
}

// Add element
func (s *Set[T]) Add(val T) {
	if !s.Contains(val) {
		s.data = append(s.data, val)
	}
}

// Check existence
func (s *Set[T]) Contains(val T) bool {
	for _, v := range s.data {
		if v == val {
			return true
		}
	}
	return false
}

// Remove element
func (s *Set[T]) Remove(val T) {
	for i, v := range s.data {
		if v == val {
			s.data = append(s.data[:i], s.data[i+1:]...)
			return
		}
	}
}

func Set_Slice_DSA() {

	fmt.Println("\n===> Set Slice DSA")

	s := &Set[int]{}
	s.Add(1)
	s.Add(2)
	s.Add(1)            // ignored
	fmt.Println(s.data) // [1 2]

	s.Remove(1)
	fmt.Println(s.data) // [2]
}
