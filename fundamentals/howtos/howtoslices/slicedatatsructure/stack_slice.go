package slicedatatsructure

import "fmt"

// Stack[T] defines a generic stack structure
// T is a type parameter (can abe init, string, custom struct, etc.)
type Stack[T any] struct {
	items []T // underlying slice thta stores the stack's element
}

// Push adds a new element onto the stack
func (s *Stack[T]) Push(value T) {
	// append adds the value to the end of the slice
	// the "end" of the slice is considered the "top" of the stack
	s.items = append(s.items, value)
}

// Pop removes and returns the top element of the stack
// It also returns a boolean to indicate if the opration was successful
func (s *Stack[T]) Pop() (T, bool) {
	// if the stack is empty, return the zero value of T and false
	if (len(s.items)) == 0 {
		var zero T
		return zero, false
	}

	// get the last element (top of the stack)
	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]

	// shrink the slice to remove the last element
	s.items = s.items[:lastIndex]

	return value, true
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// Size returns how many elements are in the stack
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns how many elements are in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Print is a helper function to display the stack
func (s *Stack[T]) Print() {
	fmt.Println("stack: ", s.items)
}

func Stack_Slice_DSA() {
	fmt.Println("\n==> Stack Slice DSA")

	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	intStack.Print()

	// Peek the top
	top, _ := intStack.Peek()
	fmt.Println("Peek:", top) // Peek: 30

	// Pop values
	val, _ := intStack.Pop()
	fmt.Println("Popped:", val) // Popped: 30

	intStack.Print() // Stack: [10 20]

	// Size
	fmt.Println("Size:", intStack.Size()) // Size: 2

	// Check if empty
	fmt.Println("IsEmpty:", intStack.IsEmpty()) // IsEmpty: false
}
