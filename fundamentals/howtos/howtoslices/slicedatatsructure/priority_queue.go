package slicedatatsructure

import "fmt"

// MinHeap using slice
type MinHeap []int

func (h *MinHeap) Push(val int) {
	*h = append(*h, val)
	h.up(len(*h) - 1)
}

func (h *MinHeap) Pop() (int, bool) {
	if len(*h) == 0 {
		return 0, false
	}
	val := (*h)[0]
	last := len(*h) - 1
	(*h)[0] = (*h)[last]
	*h = (*h)[:last]
	h.down(0)
	return val, true
}

// heapify up
func (h *MinHeap) up(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2
		if (*h)[idx] >= (*h)[parent] {
			break
		}
		(*h)[idx], (*h)[parent] = (*h)[parent], (*h)[idx]
		idx = parent
	}
}

// heapify down
func (h *MinHeap) down(idx int) {
	n := len(*h)
	for {
		left := 2*idx + 1
		right := 2*idx + 2
		smallest := idx

		if left < n && (*h)[left] < (*h)[smallest] {
			smallest = left
		}
		if right < n && (*h)[right] < (*h)[smallest] {
			smallest = right
		}
		if smallest == idx {
			break
		}
		(*h)[idx], (*h)[smallest] = (*h)[smallest], (*h)[idx]
		idx = smallest
	}
}

// Length helper
func (h *MinHeap) Len() int {
	return len(*h)
}

func Priority_Queue_Slice_DSA() {
	fmt.Println("\n==> Priority Queue DSA")

	h := &MinHeap{}
	h.Push(5)
	h.Push(3)
	h.Push(8)
	h.Push(1)

	for h.Len() > 0 {
		val, _ := h.Pop()
		fmt.Println(val)
	}
}
