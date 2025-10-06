package slicedatatsructure

import "fmt"

type RingBuffer[T any] struct {
	data  []T
	size  int
	start int
	end   int
	full  bool
}

// Initialize a new Ring buffer
func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{data: make([]T, size), size: size}
}

// Add item
func (r *RingBuffer[T]) Add(val T) {
	r.data[r.end] = val
	r.end = (r.end + 1) % r.size
	if r.full {
		r.start = (r.start + 1) % r.size
	}
	if r.end == r.start {
		r.full = true
	}
}

// Get all items in order
func (r *RingBuffer[T]) Items() []T {
	if !r.full && r.end == 0 {
		return []T{}
	}
	var result []T
	i := r.start
	for {
		if !r.full && i == r.end {
			break
		}
		result = append(result, r.data[i])
		i = (i + 1) % r.size
		if r.full && i == r.end {
			break
		}
	}
	return result
}

func Ring_Buffer_Slice_DSA() {
	fmt.Println("\n==> Ring Buffer Slice DSA")

	buf := NewRingBuffer[int](4)

	buf.Add(1)
	buf.Add(2)
	buf.Add(3)
	buf.Add(4)
	buf.Add(5) // overwrites oldest

	fmt.Println(buf.Items()) // [2 3 4 5]
}
