package slicedatatsructure

import "fmt"

// Queue[T] defines a generic queue structures
type Queue[T any] struct {
	items []T // underlying slice to hold elements
}

// Enqueue adds an element to the "back" of the queue
func (q *Queue[T]) Enqueue(value T) {
	// appends puts the value at the end
	q.items = append(q.items, value)
}

// Dequeue removes and returns the element at the "front" of the queu
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	value := q.items[0]

	q.items = q.items[1:]

	return value, true
}

// Peek returns the front element without removing it
func (q *Queue[T]) Peek() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	return q.items[0], true
}

// IsEmpty checks if queue has no elements
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// Print shows the current queue contents
func (q *Queue[T]) Print() {
	fmt.Println("Queue:", q.items)
}

func Queue_Slice_DSA() {
	fmt.Println("\n===> Queue Slice DSA")

	// Create a queue of strings
	var q Queue[string]

	// Enqueue items
	q.Enqueue("A")
	q.Enqueue("B")
	q.Enqueue("C")
	q.Print() // Queue: [A B C]

	// Peek front
	front, _ := q.Peek()
	fmt.Println("Peek:", front) // Peek: A

	// Dequeue items
	val, _ := q.Dequeue()
	fmt.Println("Dequeued:", val) // Dequeued: A

	q.Print() // Queue: [B C]

	// Size
	fmt.Println("Size:", q.Size()) // Size: 2

	// Check empty
	fmt.Println("IsEmpty:", q.IsEmpty()) // IsEmpty: false
}
