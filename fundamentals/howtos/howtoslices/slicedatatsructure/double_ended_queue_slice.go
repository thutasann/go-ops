package slicedatatsructure

import "fmt"

// Dequeue structure using slice
type Deque[T any] struct {
	items []T
}

// Add to back
func (d *Deque[T]) PushBack(value T) {
	d.items = append(d.items, value)
}

// Remove from back
func (d *Deque[T]) PopBack() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}
	val := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return val, true
}

// Add to front
func (d *Deque[T]) PushFront(val T) {
	d.items = append(d.items, val)
}

// Remove from front
func (d *Deque[T]) PopFront() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}
	val := d.items[0]
	d.items = d.items[1:]
	return val, true
}

// Length of Deque
func (d *Deque[T]) Len() int {
	return len(d.items)
}

func Double_Ended_Queue_slice_DSA() {
	fmt.Println("\n==> DEQUEU DSA")

	deque := &Deque[int]{}
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushFront(0)

	fmt.Println(deque.PopFront()) // 0, true
	fmt.Println(deque.PopBack())  // 2, true
	fmt.Println(deque.Len())      // 1
}
